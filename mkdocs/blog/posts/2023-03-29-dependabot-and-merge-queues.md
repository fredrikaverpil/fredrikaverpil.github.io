---
date: 2023-03-29
draft: false
tags:
- github-actions
---

# Using GitHub merge queue to ease the Dependabot churn

Good morning, project!

```bash
$ gh pr list

#1626  chore(deps-dev): bump black from 23.1.0 t...  dependabot/pip/black-23.3.0                  about 7 hours ago
#1625  chore(deps-dev): bump types-requests from...  dependabot/pip/types-requests-2.28.11.17     about 7 hours ago
#1624  chore(deps-dev): bump mkdocs-include-mark...  dependabot/pip/mkdocs-include-markdown-p...  about 7 hours ago
#1623  chore(deps-dev): bump types-redis from 4....  dependabot/pip/types-redis-4.5.3.1           about 7 hours ago
#1622  chore(deps-dev): bump pre-commit from 3.1...  dependabot/pip/pre-commit-3.2.1              about 7 hours ago
#1621  chore(deps-dev): bump types-deprecated fr...  dependabot/pip/types-deprecated-1.2.9.2      about 7 hours ago
#1620  chore(deps-dev): bump types-python-dateut...  dependabot/pip/types-python-dateutil-2.8...  about 7 hours ago
#1619  chore(deps-dev): bump types-redis from 4....  dependabot/pip/types-redis-4.5.3.0           about 7 hours ago
#1618  chore(deps-dev): bump moto from 4.1.4 to ...  dependabot/pip/moto-4.1.6                    about 7 hours ago
```

-Spits out coffee-

<!-- more -->

## Automating :simple-dependabot: dependabot PR merging

!!! warning "Don't try this at home, kids"

    You should always read through the changelog of Dependabot PRs and have at least a basic understanding of what _changes_ you introduce to your projects, your colleagues/co-authors and yourself, before merging.

    But once you've done that, maybe let's see if we can ease the pain a bit here...

So I'm fortunate to work at a company who owns a GitHub organization, and right now [merge queues](https://github.blog/changelog/2023-02-08-pull-request-merge-queue-public-beta/) is in beta for GitHub organizations. So by enabling this (in the repo settings[^1]) I can queue up all these dependabot PRs for merging in one go!

[^1]:
    By the way, you can set up merge queues to employ a "rebase and merge" method :material-star-face:.

Let's write a little script!

!!! example "`dependabot-merge.sh`"

    ```bash
    #!/bin/bash -e

    # Get the list of numbers
    pr_numbers=$(gh pr list "$@" --app dependabot --json number --jq '.[].number')

    # Iterate over each number and approve and merge the corresponding PR
    for pr_number in $pr_numbers; do
        gh pr review --approve $pr_number
        gh pr merge $pr_number
    done
    ```

    You'll need to install and authenticate the [GitHub CLI](https://cli.github.com/) to make the `gh` command accessible, which is invoked by this script.


The script will take arguments and forward to `gh`. This can be useful to filter out certain PRs you want to merge.

!!! example "Excecution examples"

    ```bash
    # Merge PRs that successfully passed CI
    ./dependabot-merge.sh --search "status:success"

    # Merge PRs that successfully passed CI and are categorized as developer dependencies
    ./dependabot-merge.sh --search "status:success chore(deps-dev) in:title " 
    ```

    See `gh pr list --help` for more examples and help on `--search`.

!!! note "My dependabot setup"

    To be able to search for `chore(deps-dev)`, you might have to add something like this to your dependabot settings:

    ```yaml
    version: 2

    updates:
      - package-ecosystem: "pip"
        commit-message:
          prefix: "chore"
          include: "scope"
    ```
