---
title: 🔎 Code review
tags: [workflow]
draft: false
summary: "Good practices during code review."

# PaperMod
ShowToc: true
TocOpen: true

created: 2022-11-14T20:42:48+01:00

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

