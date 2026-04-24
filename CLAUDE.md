# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Serve locally with drafts
hugo server -D

# Build site
hugo

# New post (then edit content/posts/<slug>.md)
hugo new posts/<slug>.md

# Rename posts to date-prefixed filenames (reads date from frontmatter)
go run rename-posts.go
```

## Architecture

Hugo static site using the [hugo-blog-awesome](https://github.com/hugo-sid/hugo-blog-awesome) theme (imported as a Go module). No custom theme code — all theme config lives in `hugo.toml`.

### Content structure

- `content/posts/` — blog posts, filenames prefixed `YYYY-MM-DD-<slug>.md`
- `content/about.md` — about page
- `content/projects/` — projects page
- `static/images/` — static assets (avatar, etc.)
- `archetypes/default.md` — template for `hugo new` (TOML frontmatter, draft=true)

### Post frontmatter

Posts use YAML frontmatter with these fields:

```yaml
title: "..."
date: 2025-01-13T03:00:00Z
draft: false
tags: ["go", "kubernetes"]
categories: ["Programming"]
description: "..."
```

### Theme config (`hugo.toml`)

Taxonomy: `categories`, `tags`, `series`. Menu: Home / Posts / About. Syntax highlighting via `markup.highlight` (github style, fenced code blocks). Google Analytics ID set under `[services.googleAnalytics]`.
