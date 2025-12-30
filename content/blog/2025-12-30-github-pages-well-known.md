---
title: "Deploying hidden dotfiles to GitHub Pages"
date: 2025-12-30
draft: false
tags: ["github-actions", "github-pages"]
categories: []
---

If you're using GitHub Pages with the official
[`actions/upload-pages-artifact`](https://github.com/actions/upload-pages-artifact)
action and need to serve files from a file/folder prefixed with a period (`.`)
sign, like the `.well-known` directory (e.g., for a custom Bluesky handle),
you'll hit a problem: **v4 of the action excludes all dotfiles by default**.

This is a
[breaking change introduced in v4.0.0](https://github.com/actions/upload-pages-artifact/releases/tag/v4.0.0)
for security reasons.

## The workaround

Create the tar archive manually from the `./public` folder and upload it with
`actions/upload-artifact`:

```yaml
- name: Create artifact
  run: |
    tar \
      --dereference --hard-dereference \
      -cvf "$RUNNER_TEMP/artifact.tar" \
      -C ./public .

- uses: actions/upload-artifact@v4
  with:
    name: github-pages
    path: ${{ runner.temp }}/artifact.tar

- uses: actions/deploy-pages@v4
```

This creates a tar without the `--exclude=".[^/]*"` pattern that
`upload-pages-artifact` uses (see
[source](https://github.com/actions/upload-pages-artifact/blob/main/action.yml)
here), allowing `.well-known` and other dotfiles through.

## Example: Custom Bluesky handle

To use a custom domain as your Bluesky handle, you need to serve your DID at
`/.well-known/atproto-did`. For a Hugo site, create
`static/.well-known/atproto-did` containing your DID:

```
did:plc:x42y69
```

With the workaround above, this file will be included in your GitHub Pages
deployment.
