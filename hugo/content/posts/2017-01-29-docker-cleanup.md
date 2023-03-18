---
ShowToc: false
TocOpen: false
date: 2017-01-29 02:00:12+01:00
draft: false
tags:
- docker
title: Docker cleanup
---

Quick and easy way to remove all containers (and their volumes) as well as all images:

```bash
# Remove containers and their volumes
docker stop $(docker ps -a -q)
docker rm -v $(docker ps -a -q)

# Remove images
docker rmi -f $(docker images -q)

# Remove unused images
docker system prune --all
```

Combine filters, `'xargs` etc:

```bash
# Stop all containers of a certain name
docker stop $(docker ps -q --filter name=mycontainer)

# Run containers based on folder/file names (use '%' where you want to insert the value corresponding to the file/folder name)
touch c1 c2 c3
ls | xargs -I % docker run --rm --name % hello-world:latest
```