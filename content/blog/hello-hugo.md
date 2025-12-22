---
title: "Hello Hugo"
date: 2025-12-22T10:00:00+01:00
draft: true
tags: ["hugo", "blog", "testing"]
categories: ["meta"]
toc: true
---

# Hello Hugo

This is a **short ingress** to test the new Hugo setup. We are migrating from
MkDocs to Hugo to have more control and a minimalist "terminal" aesthetic.

## The Body

Here is some regular text body. We are testing various markdown features to
ensure the theme handles them correctly.

### Code Block

Here is a Python code snippet:

```python
def hello_world():
    print("Hello, Hugo!")

if __name__ == "__main__":
    hello_world()
```

### Callouts

We are implementing GitHub-style callouts using a custom render hook.

> [!NOTE] This is a note callout! It should have a specific style.

> [!TIP] This is a tip. Useful for hints and tricks.
>
> The trick!

> [!IMPORTANT] This is important. Don't miss this information!

> [!WARNING] This is a warning. Be careful!

> [!CAUTION] This is a caution. High alert!

> [!EXAMPLE] This is an example callout.

> [!QUOTE] Quoting [ConradIrwin](https://github.com/ConradIrwin)
>
> This is a quote callout, useful for highlighting text or citations. It supports the special "inline title" styling we added.

> [!INFO] This is an info callout (alias for Note).

### Nested Code Block in Callout

> [!TIP] You can even put code inside a callout!
>
> ```bash
> echo "This is inside a tip"
> ```

### GoAT Diagram

Here is an ASCII diagram using GoAT (which we might need to enable or render
somehow, but for now just as a code block or pre-formatted text):

```goat
.---.       .---.
| A |---->| B |
'---'       '---'
  ^           |
  |           v
  |         .---.
  '---------| C |
            '---'
```

