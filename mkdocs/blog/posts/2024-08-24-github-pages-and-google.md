---
date: 2024-08-24
draft: false
authors:
  - fredrikaverpil
comments: true
tags:
  - github
---

# GitHub Pages and Google

Turns out that if you publish a blog (like I do) on
[GitHub pages](https://pages.github.com/), and you want the site to be indexed
by Google, it's not so easy.

I'm not sure entirely what the root cause is, but you have to manually add your
site in the [Google Search Console](https://search.google.com/search-console)
and then manually add each URL individually for indexing.
[This discussion](https://github.com/orgs/community/discussions/50379) outlines
the problem perfectly.

Perhaps the reason for this is there are a LOT of pages under the `github.io`
domain, so Google doesn't automatically index it all, or times out. Or perhaps
it's because GitHub has some restriction in place which avoids attracting too
much traffic to their free hosting solution. 🤷
