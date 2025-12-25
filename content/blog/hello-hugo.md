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

Here is a static Python code snippet:

```python
def hello_world():
    print("Hello, Hugo!")

if __name__ == "__main__":
    hello_world()
```

### Interactive Code Block

Now here's an **interactive** Python code snippet using Codapi. Click "Run" to execute it in your browser!

{{< codapi sandbox="python" >}}def greet(name):
    print(f"Hello, {name}!")

greet("Hugo")
greet("World")
{{< /codapi >}}

Try modifying the code above and running it again. Here's another example with a simple calculation:

{{< codapi sandbox="python" >}}# Calculate factorial
def factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)

for i in range(1, 6):
    print(f"{i}! = {factorial(i)}")
{{< /codapi >}}

### Interactive JavaScript

Here's an interactive JavaScript example that runs directly in the browser (no WASI needed):

{{< codapi sandbox="javascript" >}}const numbers = [1, 2, 3, 4, 5];
const squared = numbers.map(n => n * n);
console.log("Original:", numbers);
console.log("Squared:", squared);

// Try adding your own code!
const sum = squared.reduce((a, b) => a + b, 0);
console.log("Sum of squares:", sum);
{{< /codapi >}}

### Interactive SQLite

And here's a SQLite example using WASI (runs entirely in your browser):

{{< codapi sandbox="sqlite" >}}SELECT 'Hello, SQLite!' as message, 42 as answer
UNION ALL
SELECT 'It works!', 100;
{{< /codapi >}}

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

### Collapsible Callouts

You can make callouts collapsible by adding `+` (open by default) or `-` (minimized by default) after the type.

> [!NOTE-] This is a **minimized** note.
>
> You have to click to see this content! Hidden by default.

> [!TIP+] This is an **expanded** tip.
>
> It starts open, but you can click the header to collapse it.

> [!EXAMPLE-] Interactive code in a minimized callout
>
> {{< codapi sandbox="python" >}}print("Peek-a-boo!"){{< /codapi >}}

### Nested Code Block in Callout

> [!TIP] You can even put code inside a callout!
>
> ```bash
> echo "This is inside a tip"
> ```

### Interactive Code in Callout

You can also use interactive Codapi snippets inside callouts, and they follow the compounding contrast rules:

> [!EXAMPLE] Try this interactive Python example inside a callout
>
> {{< codapi sandbox="python" >}}# Calculate Fibonacci sequence
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

for i in range(10):
    print(f"F({i}) = {fibonacci(i)}")
{{< /codapi >}}
>
> Text underneath it!

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

