---
ShowToc: false
TocOpen: false
date: 2022-12-17 21:42:21+01:00
draft: false
summary: Good practices during code review.
tags:
- workflow
title: "\U0001F50E Code review"
updated: 2022-12-17 20:40:27+01:00
---

## Assert using previous production code

All tests added in PR, must fail:

```bash
# 1. Restore the production code from master
git restore -s master <path to production file>

# 2. Re-run tests
# 3. Verify all new tests fail
# 4. Restore files back to current PR branch's state
git restore <path to production file>
```