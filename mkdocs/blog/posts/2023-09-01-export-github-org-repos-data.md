---
date: 2023-09-01
draft: false
tags:
  - github
---

# Export repo data from GitHub organization

This is a quickie, using the [GitHub CLI](https://cli.github.com/) and some Python.

Replace `YOUR_ORG` with your organization and export all your organization's repo data to `repos.json`:

```bash
gh api \
        -H "Accept: application/vnd.github+json" \
        -H "X-GitHub-Api-Version: 2022-11-28" \
        /orgs/YOUR_ORG/repos \
        --paginate > repos.json
```

You can now import this into e.g. Excel ([which by the way now have Python support](https://techcommunity.microsoft.com/t5/excel-blog/announcing-python-in-excel-combining-the-power-of-python-and-the/ba-p/3893439)), or optionally convert it to csv first:

```python
import json


def main():
    # Read repos.json
    with open("repos.json", "r") as f:
        data = json.load(f)

    print(f"Repo count: {len(data)}")

    # Export repos.csv
    with open("repos.csv", "w") as f:
        keys = data[0].keys()
        f.write(",".join(keys) + "\n")

        for repo in data:
            values = [str(repo[key]) for key in keys]
            f.write(",".join(values) + "\n")


if __name__ == "__main__":
    main()
```
