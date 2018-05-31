---
title: Pod
weight: 10
menu: true
---

Pod

Pod is a group of one or more tightly related containers that will run together on the same worker node and in the same Linux namespace(s).

Each pod has its own IP, Hostname, processes etc. running a single application and/or additional supporting processes, each running in its own container.

Pods are spread out on different worker nodes.

```shell

workshop $ kubectl apply -f https://polarsquad.github.io/kubernetes-workshop/examples/pod.yaml
pod "example-app" created

workshop $ kubectl get pods
NAME                       READY     STATUS    RESTARTS   AGE
example-app-f6h54   1/1         Running   0          19s

workshop $ kubectl delete -f https://polarsquad.github.io/kubernetes-workshop/examples/pod.yaml
pod "example-app" deleted

```