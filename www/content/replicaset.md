---
title: ReplicaSet
weight: 30
menu: true
---

ReplicaSet

Usually one does not create pods directly but create other resources like ReplicaSet's which creates and manages pods.
(This because if node fails then those pods are lost.)

Creating ReplicaSet

```shell

workshop $ kubectl create -f examples/replicaset.yaml
replicaset "kubers" created


```

Examining ReplicaSet and Pods

```shell

workshop $ kubectl get rs
NAME      DESIRED   CURRENT   READY     AGE
kubers    3         3         3         55s

workshop $ kubectl get pods
NAME           READY     STATUS    RESTARTS   AGE
kubers-2bhns   1/1       Running   0          1m
kubers-xghq7   1/1       Running   0          1m
kubers-zg6kb   1/1       Running   0          1m

workshop $ kubectl get pods -o wide
workshop $ kubectl describe pod kubers-2bhns
workshop $ kubectl logs kubers-2bhns

```



Seeing what happens if pod is deleted in ReplicaSet

```shell

workshop $ kubectl kubectl delete pod kubers-2bhns 

workshop $ kubectl get pods
NAME           READY     STATUS        RESTARTS   AGE
kubers-2bhns   1/1       Terminating   0          5m
kubers-cmz6z   1/1       Running       0          5s
kubers-xghq7   1/1       Running       0          5m
kubers-zg6kb   1/1       Running       0          5m

```


Horizontally scaling pods

```shell

workshop $ kubectl scale rs kubers --replicas=10

workshop $ kubectl get pods

workshop $ kubectl scale rs kubers --replicas=3

```

Delete ReplicaSet

```shell

workshop $ kubectl delete rs replicaset-example

workshop $ kubectl get pods

workshop $ kubectl get rs

```
