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

## Rewriting working software is proof the original design was already good

`rupa/z` is just good. The matching, the aging, the flags, none of it needed fixing, so I didn't touch any of it. I just moved the same design into Rust and kept my hands off the parts that were already right. rupa wrote a man page instead of a friendly README, kept the whole thing to one shell script, and never bloated it with features nobody asked for. That's taste. It's rarer than it should be.

Most of the OSS I run into these days wants to be everything. Plugins, config files, a dashboard, a roadmap. `rupa/z` never went there, and it's still exactly as useful seventeen years later as it was on day one. We need more of that: people who build one small thing, get the shape of it right, and then leave it alone.

## Why bother

A rewrite for its own sake is a waste of time. 
At some point it was a matter of writing Rust for the fun and seeing how much handholding I need with an LLM to assist me.
Turned out well, since I use it for a few years now.

## The end result

`zjyo` is under 800 lines of Rust across five files. `database.rs` handles persistence and matching. `entry.rs` is a plain struct with a frecency calculation. `cli.rs` wires up `clap` and dispatches. That's the whole surface area, and most of what maintaining a small tool actually looks like: not new features, but noticing the gap between what you assumed was true and what was actually happening.

## The bug that mattered most

The tmux bug I hit while using tmux-continuum is worth mentioning. First, using tmux and tmux-continuum is something that I don't plan to change in the next 10 years too. It just works. The issue is when you make an assumption that `cd` is going to always work and wondering hey why does `zjyo` do it differently, some things start to make sense, because it's a class of bug that's easy to dismiss as "user error" and hard to find by reading code. The database *was* updating, for every shell session I'd opened since editing `.zshrc`. But `tmux-continuum` restores sessions across machine restarts, and panes that survive a restore don't re-source your rc files. Config and the `cd` wrapper are two different things, and only one of them updates when you edit a file.

The fix, using `precmd_functions` instead of overriding `cd`, happens to also route around the entire class of bug, since it's additive rather than a redefinition. It doesn't fix stale shells. Nothing can fix a process that's already running with old code loaded. But it means the next time something like this happens, at least new panes won't diverge from what you think your shell does without you noticing.

### Walking through it

The original integration looked like this:

```bash
cd() {
    builtin cd "$@" && zjyo --add
}
```

Redefining `cd` as a shell function is the obvious approach, anyone who has skimmed a `z` integration script has seen this pattern. It works, right up until something else in your shell startup also defines `cd`. Whichever definition loads last wins, with no error and no warning. The earlier one just stops existing.

That wasn't actually my bug though, no other plugin was fighting for `cd` in my case. The real problem was simpler and harder to see: some of my `zjyo`-tracked directories just weren't showing up in `z -l`, despite me having `cd`'d into them dozens of times that day. The database file (`~/.z`) was clearly being written to. `stat ~/.z` showed a recent mtime, other directories were tracked fine, and `zjyo --add` worked perfectly when I ran it by hand in the same pane. So the binary was fine, the wrapper function was fine when invoked, and yet specific panes weren't invoking it.

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

If `cd` reports as a builtin instead of a function, that pane never picked up the wrapper. No amount of retrying `z --add` in that shell fixes it, because the function it would need to call doesn't exist there.

`precmd_functions` gets around this. Every zsh prompt draw runs every function registered in that array:

```bash
_zjyo_precmd() {
    (zjyo --add &)
}
precmd_functions+=(_zjyo_precmd)
```

`+=` means appending, not overwriting, so it can't replace someone else's hook the way redefining `cd` can. And because it hooks the prompt rather than the `cd` builtin specifically, it also picks up `pushd`, `popd`, and any script that changes directory without going through the user's interactive `cd` at all. It doesn't retroactively fix the stale panes sitting there with the old function loaded, those still need a `source ~/.zshrc` or a fresh pane. But new panes stop quietly diverging from what your config says.

## Why keep doing this

There's no ecosystem reason to prefer `zjyo` over `zoxide`. `zoxide` has more users, more contributors. The reason to maintain your own small thing isn't that it's objectively better. It's that you understand every line of it, you can fix what actually bothers you instead of filing an issue and waiting, and the maintenance itself is quite lean when there's not much functionality and need for it, to begin with.

800 lines and a decade-old database format don't need a roadmap. They need someone willing to keep them that small.
