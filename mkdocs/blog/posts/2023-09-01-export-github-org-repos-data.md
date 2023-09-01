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

    - allow_forking
    - archive_url
    - archived
    - assignees_url
    - blobs_url
    - branches_url
    - clone_url
    - collaborators_url
    - comments_url
    - commits_url
    - compare_url
    - contents_url
    - contributors_url
    - created_at
    - default_branch
    - deployments_url
    - description
    - disabled
    - downloads_url
    - events_url
    - fork
    - forks
    - forks_count
    - forks_url
    - full_name
    - git_commits_url
    - git_refs_url
    - git_tags_url
    - git_url
    - has_discussions
    - has_downloads
    - has_issues
    - has_pages
    - has_projects
    - has_wiki
    - homepage
    - hooks_url
    - html_url
    - id
    - is_template
    - issue_comment_url
    - issue_events_url
    - issues_url
    - keys_url
    - labels_url
    - language
    - languages_url
    - license
    - merges_url
    - milestones_url
    - mirror_url
    - name
    - node_id
    - notifications_url
    - open_issues
    - open_issues_count
    - owner
    - permissions
    - private
    - pulls_url
    - pushed_at
    - releases_url
    - size
    - ssh_url
    - stargazers_count
    - stargazers_url
    - statuses_url
    - subscribers_url
    - subscription_url
    - svn_url
    - tags_url
    - teams_url
    - topics
    - trees_url
    - updated_at
    - url
    - visibility
    - watchers
    - watchers_count
    - web_commit_signoff_required

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
