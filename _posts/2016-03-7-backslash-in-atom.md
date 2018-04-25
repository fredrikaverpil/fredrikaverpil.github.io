---
layout: post
title: Backslash in Atom
tags: [windows]
---

On Windows 10, Swedish keyboard. When I hit `AltGr` (right `Alt` button) and `+` (which is supposed to give me `\`), I get nothing in [Atom](https://atom.io).

<!--more-->

After having installed the [keyboard-localization](https://atom.io/packages/keyboard-localization) package, I now get `~` when I hit this key combo, which is not correct. If I hold down *both* `Alt` keys and hit `+` I get `\` with this package installed. However this is not acceptable behavior.

The workaround which worked for me was to uninstall the keyboard-localization package and follow the instructions posted by @csvn [here]( https://github.com/atom/atom/issues/8820#issuecomment-146959203) which explains to add the following to the `keymap.cson`:

```
'atom-workspace atom-pane':
  'ctrl-alt-=': 'unset!'
```

To quickly open up `keymap.cson` from inside of Atom, hit `Ctrl+,` to enter settings, then choose `Keybindings`. Click the link in the text to open up `keymap.cson`.
