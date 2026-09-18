---
title: "Projects"
date: 2025-01-01T00:00:00Z
draft: false
---

A few things I've built and some I kept maintaining too.

**[zjyo](https://github.com/syndbg/zjyo)** ![GitHub stars](https://img.shields.io/github/stars/syndbg/zjyo?style=flat&label=%E2%98%85) — a Rust port of [rupa/z](https://github.com/rupa/z), same frecency algorithm and database format. The one I actually use every day to jump between directories.

**[goenv](https://github.com/syndbg/goenv)** ![GitHub stars](https://img.shields.io/github/stars/syndbg/goenv?style=flat&label=%E2%98%85) — a Go version manager, like pyenv or rbenv but for Go. My most-used project by other people.

**[gocat](https://github.com/sumup-oss/gocat)** ![GitHub stars](https://img.shields.io/github/stars/sumup-oss/gocat?style=flat&label=%E2%98%85) — like `socat`, but in Go. A multipurpose relay for data transfer and monitoring. It was useful at my time in SumUp when we wanted to proxy SSH traffic for Docker auth and `socat` was randomly hanging. As part of a build system with nearly thousand apps to build, that was why this made sense.

**[vaulted](https://github.com/sumup-oss/vaulted)** ![GitHub stars](https://img.shields.io/github/stars/sumup-oss/vaulted?style=flat&label=%E2%98%85) — encryption/decryption tool using AES256-GCM, for teams that need auditable secrets workflows. Enabled SumUp to shift-left on secret management.

**[terraform-provider-vaulted](https://github.com/sumup-oss/terraform-provider-vaulted)** ![GitHub stars](https://img.shields.io/github/stars/sumup-oss/terraform-provider-vaulted?style=flat&label=%E2%98%85) — a Terraform provider for storing vaulted secrets safely in source control. When it made sense to also enable provisioning your HashiCorp Vault secrets.

**[terraform-provider-vaulted-null](https://github.com/syndbg/terraform-provider-vaulted-null)** ![GitHub stars](https://img.shields.io/github/stars/syndbg/terraform-provider-vaulted-null?style=flat&label=%E2%98%85) — secure secrets for every SCM and every Terraform resource, without needing a running HashiCorp Vault backend.

**[webpack-gcs-plugin](https://github.com/syndbg/webpack-google-cloud-storage-plugin)** ![GitHub stars](https://img.shields.io/github/stars/syndbg/webpack-google-cloud-storage-plugin?style=flat&label=%E2%98%85) — back when uploading Webpack assets to Google Cloud Storage wasn't supported out of the box. Whether that's still true I genuinely don't know, haven't touched it in years, but it picked up traction on its own.

---

More on my [GitHub profile](https://github.com/syndbg).
