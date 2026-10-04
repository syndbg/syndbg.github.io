---
title: "On writing small useful things like zjyo"
date: 2026-09-19T01:48:25+03:00
draft: false
tags: ["rust", "cli", "tooling", "developer-tools", "z", "frecency"]
categories: ["Programming", "Developer Tools"]
description: "Why I rewrote rupa/z in Rust instead of switching to zoxide or jump, and why the smallest tool I've written is my most useful one to me."
---

I've used `z` for over a decade. Not `zoxide`, not `jump`, not `fasd`, the original [rupa/z](https://github.com/rupa/z), a few hundred lines of POSIX shell that tracks the directories you `cd` into and lets you jump back with a fuzzy pattern instead of a full path.

At some point I tried the newer alternatives. Each one changed something I didn't ask it to change: a different ranking algorithm, a different database format, a new set of flags to relearn. None of them were bad tools. They just weren't `z`. I wanted the exact behavior I already had muscle memory for, just without a shell interpreter parsing a database file on every invocation.

So I wrote [zjyo](https://github.com/syndbg/zjyo) a while back, a 1:1 Rust port. The only thing that changed is what's running under the hood.

## Rewriting a small tool showed me what was already right

`rupa/z` already had the behavior I wanted: matching, aging, and a small set of flags. I kept those choices and moved the implementation to Rust. Its single shell script and man page are part of the appeal. The tool does one job and has stayed useful without growing a dashboard or a plugin system. That is the kind of scope I wanted to preserve.

## Why bother

I started the rewrite to practice Rust and see how much help I would need from an LLM. Keeping `z`'s behavior gave me a concrete target. It turned out well enough that I still use `zjyo` years later. I review the code myself; the model helped with Rust, not with deciding what the tool should do.

## The end result

At the time I wrote this, `zjyo` was under 800 lines of Rust across five files. `database.rs` handles persistence and matching. `entry.rs` stores a path and its frecency score. `cli.rs` wires up `clap` and dispatches commands. Most maintenance has been checking assumptions against what the shell and database actually do.

## The bug that mattered most

At the time, I was using tmux with tmux-continuum. I have since moved my dotfiles to Zellij, but the bug was in the restored shell process. Continuum brought panes back after a restart without re-reading `.zshrc`. Fresh panes used the new `cd` wrapper and updated the database. Restored panes still had the old function definitions.

The fix, using `precmd_functions` instead of overriding `cd`, happens to also route around the entire class of bug, since it's additive rather than a redefinition. It doesn't fix stale shells. Nothing can fix a process that's already running with old code loaded. But it means the next time something like this happens, at least new panes won't diverge from what you think your shell does without you noticing.

### Walking through it

The original integration looked like this:

```bash
cd() {
    builtin cd "$@" && zjyo --add
}
```

Redefining `cd` as a shell function is the obvious approach, anyone who has skimmed a `z` integration script has seen this pattern. It works, right up until something else in your shell startup also defines `cd`. Whichever definition loads last wins, with no error and no warning. The earlier one just stops existing.

No other plugin was replacing `cd`. The database file (`~/.z`) had a recent modification time, other directories appeared in `z -l`, and `zjyo --add` worked when run by hand in the suspect pane. The binary and database were fine. That shell simply had not loaded the wrapper.

The panes in question all had one thing in common: they were tmux panes that survived a `tmux-continuum` restore, meaning they were spawned before I'd last edited `.zshrc` to add the `zjyo` integration. A shell process reads its rc file once, at startup. Editing `~/.zshrc` after the fact does nothing to a shell that's already running. The function definitions it loaded at spawn time are the only ones it has. `tmux-continuum` is built to survive machine restarts by serializing pane state and restoring it later, so a pane you're typing into today can be running a shell process that's been alive, uninterrupted, since before a config change you made weeks ago.

Confirming this took nothing more than:

```bash
# in a suspect pane
type cd
# cd is a shell builtin
```

versus a freshly spawned pane:

```bash
type cd
# cd is a shell function from /Users/syndbg/.zshrc
```

If `cd` reports as a builtin instead of a function, that shell never loaded the wrapper. Running `zjyo --add` by hand can update the database, but it does not install the missing function. Source `~/.zshrc` or open a fresh pane.

`precmd_functions` gets around this. Every zsh prompt draw runs every function registered in that array:

```bash
_zjyo_precmd() {
    (zjyo --add &)
}
precmd_functions+=(_zjyo_precmd)
```

`+=` means appending, not overwriting, so it can't replace someone else's hook the way redefining `cd` can. And because it hooks the prompt rather than the `cd` builtin specifically, it also picks up `pushd`, `popd`, and any script that changes directory without going through the user's interactive `cd` at all. It doesn't retroactively fix the stale panes sitting there with the old function loaded, those still need a `source ~/.zshrc` or a fresh pane. But new panes stop quietly diverging from what your config says.

## Why keep doing this

There is no ecosystem reason to choose `zjyo` over `zoxide`; `zoxide` has more users and contributors. I keep `zjyo` because I know its code and can change the behavior I rely on. A small tool stays small when I resist adding features that do not solve a problem I have.
