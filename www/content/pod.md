---
title: Pod
weight: 20
menu: true
---

Pod

Pod is a group of one or more tightly related containers that will run together on the same worker node and in the same Linux namespace(s).

Each pod has its own IP, Hostname, processes etc. running a single application and/or additional supporting processes, each running in its own container.

Pods are spread out on different worker nodes.

Good practise is to have frontend and backend containers in different pods. This way utilization of your infrastructure is more efficient.

Example pod:

```yaml
#example pod
apiVersion: v1
kind: Pod
metadata:
  name: example-app
spec:
  containers:
  - name: example-app
    image: nginx:1.7.9
    ports:
    - containerPort: 80
```

Creating Pod from example yaml.

```shell

workshop $ kubectl create -f examples/pod.yaml
pod "example-app" created

workshop $ kubectl get pods
NAME                       READY     STATUS    RESTARTS   AGE
example-app   1/1         Running   0          19s

```

Examining and Running commands inside container.

```shell

workshop $ kubectl exec example-app env

workshop $ kubectl exec -it example-app bash

```

Deleting Pod

```shell

workshop $ kubectl get pods
NAME                       READY     STATUS    RESTARTS   AGE
example-app   1/1         Running   0          19s

workshop $ kubectl delete -f examples/pod.yaml
pod "example-app" deleted

```