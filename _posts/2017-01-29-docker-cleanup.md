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
