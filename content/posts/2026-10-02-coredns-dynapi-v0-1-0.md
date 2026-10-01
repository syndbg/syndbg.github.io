---
title: "coredns/dynapi v0.1.0: an HTTP API for CoreDNS records, 8 years later"
date: 2026-10-02T00:00:00+03:00
draft: false
tags: ["go", "coredns", "dns", "kubernetes", "open-source", "api"]
categories: ["Programming", "Infrastructure"]
description: "coredns/dynapi v0.1.0 is out. It lets applications read, replace and delete A and AAAA records through an authenticated HTTP API. The idea started as coredns/coredns PR 1822 in May 2018."
---

![dynapi demo](https://github.com/coredns/dynapi/raw/v0.1.0/docs/preview.gif)

tl;dr

- [coredns/dynapi v0.1.0](https://github.com/coredns/dynapi/releases/tag/v0.1.0) is out, published on 2026-10-01.
- It is an external CoreDNS plugin. Applications read, replace and delete A and AAAA record sets through a JSON HTTP API.
- The idea started as [coredns/coredns#1822](https://github.com/coredns/coredns/pull/1822). I opened that PR on 2018-05-20. It was closed on 2018-09-29, when the `coredns/dynapi` repository was created.
- That repository stayed empty for eight years. The plugin in v0.1.0 is a new design, not a port of the old code.

---

## Why I needed it

At SumUp we had what felt like an amazing idea: build our own take on [ArgoCD](https://argo-cd.readthedocs.io/). It was a Kubernetes-based deployment system for on-demand dev environments, and it ran on bare metal instances. It worked.

Every new environment needed DNS names, and it needed them fast. So we needed a DNS server that was reliable, simple to run and extensible. [CoreDNS](https://coredns.io/) fit those three needs. One thing was missing: a way to register addresses dynamically. The obvious answer was an HTTP API, because any language and any tool can call one.

## The 2018 story

[PR #1822](https://github.com/coredns/coredns/pull/1822) was titled "[WIP] Dynamic updates API with listen server". It was an experimental REST API. A POST request with an A record added it to a zone managed by the [`file` plugin](https://coredns.io/plugins/file/).

My use case was narrow:

- CoreDNS is the main DNS server.
- I add records through an API.
- There is no database as a second source of truth. In-memory storage was good enough for me.

We ran it in production from a little before 20 May 2018. A small daemon read DHCP leases and sent them to CoreDNS to register A records. It worked well, with one catch. Records created without a TTL got the default of 3600 seconds. That is far too long for short-lived environments, so the create and update requests got a `TTL` field.

The maintainers had a fair question. [Miek Gieben](https://github.com/miekg) asked why CoreDNS should provision records at all. Access control, users and permissions grow out of every API like this. The advice was to build it as an external plugin and put it at `coredns/dynapi`. [John Belamaric](https://github.com/johnbelamaric) suggested a different split: a Go "Writable" interface that plugins can implement, and separate plugins for the write protocols (DDNS, HTTPS, gRPC).

The same request was already open in [#522](https://github.com/coredns/coredns/issues/522) ("provide an API to manage host records", 2017). In 2020 John Belamaric listed what a proper solution needs: a "writable" interface for backends, a plugin that exposes the REST API, and authentication that works across API plugins. [#2073](https://github.com/coredns/coredns/issues/2073) asked for the same thing from the Ansible side, and [#2517](https://github.com/coredns/coredns/issues/2517) was closed as unsupported, although a maintainer later said DNS UPDATE is "hard to do right, but I'm not against it".

The repository was created in September 2018 and the PR was closed. Then my part stalled. COVID happened, I am bad at GitHub notifications, and I decided nobody needed the project. I never moved the code.

I was wrong about the last part. For years, people kept asking in the closed PR. Some found the empty repository. Some asked why CoreDNS has no REST API. Some discussed which record types to support and whether Envoy's [xDS](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol) is a better model.

On 2026-10-01, out of nowhere, I checked again and saw the idea was still relevant. And here we are.

## What v0.1.0 does

dynapi runs inside CoreDNS and opens its own HTTP listener. The default address is `127.0.0.1:8080`, and only loopback addresses are allowed.

There is one URL shape:

```text
/v1/zones/{zone}/records/{name}/{type}
```

The type is A or AAAA. Names must be full names inside the configured zone.

- GET returns the exact stored set, for example `{"ttl":60,"addresses":["192.0.2.10"]}`.
- PUT replaces the whole set in one atomic step.
- DELETE removes the set and returns 204, also when the set does not exist.

PUT is strict. It needs `Content-Type: application/json`, an explicit TTL from 0 to 2147483647, and 1 to 256 addresses of the right family. Duplicate addresses are removed. Bodies are limited to 64 KiB. Unknown fields and duplicate keys are rejected. The [OpenAPI specification](https://github.com/coredns/dynapi/blob/v0.1.0/openapi.yaml) is generated from the Go models, and `make verify` checks that it is current.

```sh
curl -X PUT http://127.0.0.1:8080/v1/zones/example.org/records/host.example.org/A \
  -H "Authorization: Bearer $DYNAPI_TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"ttl":60,"addresses":["192.0.2.10","192.0.2.11"]}'
```

## dynupdate: the piece that showed up while I was away

In 2018 CoreDNS had no way to change records at runtime, and I wrote my own. In 2026 it has one. The [`dynupdate`](https://github.com/coredns/coredns/tree/master/plugin/dynupdate) plugin landed in CoreDNS itself in [PR #8520](https://github.com/coredns/coredns/pull/8520). Contributor houyuwushang wrote it and Yong Tang merged it in September 2026. It speaks the standard protocol for this job, RFC 2136 DNS UPDATE, so `nsupdate` and DHCP servers such as Kea can already talk to it. The work started from [#6254](https://github.com/coredns/coredns/issues/6254) ("Support DDNS", open since 2023), and two core changes came first: [#8469](https://github.com/coredns/coredns/pull/8469) lets the server accept UPDATE messages and [#8471](https://github.com/coredns/coredns/pull/8471) exposes the validated TSIG identity to plugins.

How it works:

- **One writable zone per server block.** You give it a seed zone file. The plugin never modifies that file.
- **Persistence is optional.** With `database`, every update goes to an embedded [bbolt](https://github.com/etcd-io/bbolt) file before the new snapshot is visible and before the client sees success. Without it, updates live in memory and a restart loses them. The first access creates the database from the seed. After that, the database is the source of truth.
- **Authentication is TSIG.** The `tsig` plugin validates the signature. dynupdate never sees the secret.
- **Authorization is explicit.** Every change must match an `allow KEY NAME TYPE` rule. Nothing is allowed by default.
- **The protocol is complete enough.** It supports RFC 2136 prerequisites, add and delete operations, CNAME and apex SOA and NS rules, and automatic SOA serial updates. AXFR reads the current snapshot, and a successful change sends a best-effort NOTIFY.
- **Caching is handled.** The `cache` plugin bypasses dynamic zones, so authoritative answers are always current.
- **Limits are stated.** No DNSSEC records, no IXFR, no multi-primary replication. The README calls the plugin experimental and says it is meant for small zones that change rarely.

That last point fits my own use case well. It also answers Miek's old worry. CoreDNS has a place for write permissions and persistence now, and the HTTP layer can stay thin.

## How it works

dynapi has no record store of its own. It sends signed DNS requests to the same CoreDNS server over pooled loopback TCP connections.

- GET is a signed AXFR (zone transfer).
- PUT and DELETE are DNS UPDATE transactions.
- [`dynupdate`](https://github.com/coredns/coredns/tree/master/plugin/dynupdate) owns the zone database. It handles persistence, atomic updates and write permissions.
- [`tsig`](https://coredns.io/plugins/tsig/) authenticates the DNS requests.
- [`transfer`](https://coredns.io/plugins/transfer/) allows the AXFR.

There are two separate credentials. HTTP clients use a bearer token of at least 32 characters. It permits reads across the whole zone. The DNS side uses a TSIG key (HMAC-SHA256), and `dynupdate allow` rules limit what that key can write.

In the code, a PUT becomes one DNS UPDATE message:

1. A prerequisite asserts that no CNAME exists at the name. RFC 2136 would silently ignore an add next to a CNAME, so this check keeps the HTTP response honest.
2. The first update deletes the whole record set of that type.
3. The next updates add the new addresses with the requested TTL.

The server applies all of it as one transaction, so readers never see a half-replaced set. DELETE is the same message without the prerequisite and the adds. GET sends a signed AXFR and scans the transfer for the one name and type. The pool reuses TCP connections, and a connection that failed or was cancelled is discarded.

dynupdate answers are mapped to HTTP status codes. A failed prerequisite is 409, a refused update is 403, and any other rejection is 502.

This is a complete Corefile:

```corefile
example.org:1053 {
    bind 127.0.0.1

    dynapi 127.0.0.1:8080 {
        token_env DYNAPI_TOKEN
        upstream 127.0.0.1:1053
        identity update-key.example.org.
        secret_env DYNAPI_TSIG_SECRET
    }

    tsig {
        secret update-key.example.org. {$DYNAPI_TSIG_SECRET}
        require_opcode UPDATE
        require AXFR
    }

    transfer {
        to 127.0.0.1
    }

    dynupdate {
        file example.org.zone
        database example.org.db
        allow update-key.example.org. host.example.org. A AAAA
    }
}
```

## Why signed DNS and not a record store

This design answers Miek's 2018 concern. dynapi does not own users, permissions or storage. `dynupdate` and `tsig` already do that work, and dynapi only translates HTTP into signed DNS.

[ADR 0001](https://github.com/coredns/dynapi/blob/v0.1.0/docs/adr/0001-use-signed-dns-for-record-access.md) records the decision and what it rejects:

- **Its own record store.** It would duplicate persistence, permissions and authoritative DNS behavior.
- **Direct calls into dynupdate.** That needs an upstream interface with defined transactions, permissions and lifecycle rules.
- **Normal DNS queries for reads.** They can expand wildcards and follow CNAMEs, so they cannot return the exact stored set. AXFR can.
- **A web framework.** Three HTTP operations do not need one, so dynapi uses `net/http`.

PUT also checks for a conflicting CNAME in the same transaction.

The design has costs. GET scans the whole zone. Connections come from a pool limited by `max_requests`, which defaults to 32. Writes are never retried, because a change may commit before the response is lost.

## What is not done

- Only A and AAAA records work. The 2018 thread asked for more types, and Miek wanted storage that does not care about record types. That is future work.
- It needs Go 1.27 or newer. It also needs `dynupdate` and `tsig` from a pinned CoreDNS revision, so the development build is a CoreDNS binary built from the dynapi repository.
- Corefile reloads are rejected. Restart CoreDNS to change the configuration.
- Conditional writes (record revisions) are still open.

## Next: 0.2.0

The loopback DNS bridge works, but it is a detour. An HTTP request turns into a signed DNS message, goes over TCP to the same process, gets parsed and verified, and only then reaches the zone. A GET transfers the whole zone to read one record set.

[ADR 0002](https://github.com/coredns/dynapi/blob/v0.1.0/docs/adr/0002-add-a-coredns-record-management-interface.md) proposes the fix for 0.2.0. dynapi would call the zone provider directly. That removes the TCP hop, the full-zone reads and the TSIG key from dynapi. It is also the "Writable" interface John Belamaric described in 2018.

I am not the first to ask. [#7259](https://github.com/coredns/coredns/issues/7259) is an open proposal for a REST API plugin with a backend-agnostic `api.Backend` interface. Its author argues that the interface should live in the main CoreDNS project and that backends such as etcd should implement it. [#7858](https://github.com/coredns/coredns/issues/7858) asks for a plugin that manages records through API or gRPC. Both point at the same gap.

The change in CoreDNS is needed first, because today plugins have no shared way to be written to:

- **A Go interface** to read, replace and delete stored record sets, with explicit errors. It starts with exact A and AAAA sets.
- **Provider lookup by zone.** CoreDNS would find the writable provider for a zone and manage its startup and shutdown.
- **A provider that implements it.** The interface alone stores nothing. dynupdate can be the first one.
- **Clear rules** for permissions, for when a write counts as durable, and for how cached answers are invalidated.
- **Reload behavior.** A Corefile reload must replace providers without sending requests to a retired zone. This also fixes the "reloads are rejected" limit above.
- **Record revisions**, so conditional writes can stop two clients from overwriting each other.

The HTTP API stays the same during the move. The signed DNS adapter from v0.1.0 stays until the interface exists. Whether the writable provider ships with CoreDNS or as an optional plugin is still an open question. The first step is to agree on the interface upstream in the CoreDNS repository, before dynapi changes.

## Try it

`make` builds CoreDNS with dynapi, dynupdate and tsig. Run `./coredns -plugins` to check that all three are in the binary.

To add dynapi to another CoreDNS build, put this line in `plugin.cfg` after [`acl`](https://coredns.io/plugins/acl/):

```text
dynapi:github.com/coredns/dynapi/plugins/dynapi
```

Then run:

```sh
go get github.com/coredns/dynapi/plugins/dynapi@v0.1.0
go generate coredns.go
go build -o coredns .
```

The [`examples/`](https://github.com/coredns/dynapi/tree/v0.1.0/examples) directory has a starter Corefile and an annotated Go client.

Sorry to all the patient people waiting. :D Now we'll follow normal OSS cadence. Any feedback, issues and PRs are welcome in the [GitHub repo](https://github.com/coredns/dynapi).
