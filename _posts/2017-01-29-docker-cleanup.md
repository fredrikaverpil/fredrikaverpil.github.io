---
layout: post
title: Docker cleanup
tags: [docker]
---

Quick and easy way to remove all containers (and their volumes) as well as all images:

```bash
# Remove containers and their volumes
docker stop $(docker ps -a -q)
docker rm -v $(docker ps -a -q)

# Remove images
docker rmi -f $(docker images -q)
```

Update:

```bash
docker system prune
```

...will result in:

```text
WARNING! This will remove:
        - all stopped containers
        - all networks not used by at least one container
        - all volumes not used by at least one container
        - all dangling images
        - all build cache
Are you sure you want to continue? [y/N] y
```
