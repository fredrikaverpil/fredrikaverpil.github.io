---
date: 2022-12-17
draft: true
tags:
- kubernetes
title: "\U0001F3B1 Kubernetes"
---

# Kubernetes

## Pod

Enter pod:

```bash
POD=name-of-pod
kubectl exec --stdin --tty $POD -- /bin/bash
```

## Troubleshooting

```bash
kubectl describe deploy $DEPLOYMENT
kubectl get deploy $DEPLOYMENT
kubectl get pod $POD
```

## Sorting

```bash
kubectl get pods --sort-by=.metadata.creationTimestamp
kubectl get pods -lapp=$POD --sort-by=.metadata.creationTimestamp
```

## Scale down no of replicas

```bash
kubectl get deployments
kubectl scale --replicas=1 deployments/<my-microservice>
```

## Cron jobs

```bash
kubectl get cj
```

### Manual run

```bash
$ kubectl get cj | grep <my-name>

some-cronjob-name                           10 * * * *     False     1        10m             3h2m
```

```bash
$ kubectl create job --from=cronjob/some-cronjob-name fredrik-manual-test-1
```

### Suspension of cronjob

Suspension was done but when we wanted to resume it, it didn’t work. Turns out we had to remove the old one and re-deploy (for every cluster).

```bash
kubectl delete cj/some-cronjob-name
```