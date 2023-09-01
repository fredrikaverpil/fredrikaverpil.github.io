---
date: 2023-09-01
draft: false
tags:
  - github
---

# Export repo data from GitHub organization

This is a quickie, using the [GitHub CLI](https://cli.github.com/) and optionally some Python.

Replace `YOUR_ORG` with your organization and export all your organization's repo data to `repos.json`:

```bash
gh api \
        -H "Accept: application/vnd.github+json" \
        -H "X-GitHub-Api-Version: 2022-11-28" \
        /orgs/YOUR_ORG/repos \
        --paginate > repos.json
```

!!! info "Exported data"

    At the time of writing this post, this is the kind of data which will be exported from the organization's repo(s):

    - id
    - node_id
    - name
    - full_name
    - private
    - owner
    - html_url
    - description
    - fork
    - url
    - forks_url
    - keys_url
    - collaborators_url
    - teams_url
    - hooks_url
    - issue_events_url
    - events_url
    - assignees_url
    - branches_url
    - tags_url
    - blobs_url
    - git_tags_url
    - git_refs_url
    - trees_url
    - statuses_url
    - languages_url
    - stargazers_url
    - contributors_url
    - subscribers_url
    - subscription_url
    - commits_url
    - git_commits_url
    - comments_url
    - issue_comment_url
    - contents_url
    - compare_url
    - merges_url
    - archive_url
    - downloads_url
    - issues_url
    - pulls_url
    - milestones_url
    - notifications_url
    - labels_url
    - releases_url
    - deployments_url
    - created_at
    - updated_at
    - pushed_at
    - git_url
    - ssh_url
    - clone_url
    - svn_url
    - homepage
    - size
    - stargazers_count
    - watchers_count
    - language
    - has_issues
    - has_projects
    - has_downloads
    - has_wiki
    - has_pages
    - has_discussions
    - forks_count
    - mirror_url
    - archived
    - disabled
    - open_issues_count
    - license
    - allow_forking
    - is_template
    - web_commit_signoff_required
    - topics
    - visibility
    - forks
    - open_issues
    - watchers
    - default_branch
    - permissions

<!-- more -->

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
