---
ShowToc: false
TocOpen: false
date: 2022-12-17 21:41:38+01:00
draft: false
summary: Nice things I've picked up along the way.
tags:
- workflow
- testing
title: "\U0001F9EA Testing"
---

## Test-driven development

### Rules

1.  You may not write any production code when you do _**not**_ have a failing test
2.  You may not write any test code when you have a failing test
3.  You may only write the least possible amount of produiction code to make the test pass

## Red-green-refactor

1.  Write failing test
2.  Write production code (making test pass)
3.  Refactor
4. (Rinse and repeat)

## Behavior driven development

Steps:
1. Given
2. When
3. Then

There's a Python package called [behave](https://github.com/behave/behave).