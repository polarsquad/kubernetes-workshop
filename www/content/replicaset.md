---
title: ReplicaSet, Services and LoadBalancer
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


## Services

Service is a resource and created to make single point of entry to group of pods providing for example same frontend services.


```shell

workshop $ kubectl create -f examples/service.yaml

workshop $ kubectl get svc


```

Service is not yet available to external clients. Ping fails.

```shell

workshop $ curl localhost:8080
curl: (7) Failed to connect to localhost port 8080: Connection refused

```

Service is running and checking that pinging is possible within cluster. Curl command is performed inside Pod towards Service. Service redirects connection to frontend pod and http response is given back.


```shell

workshop $ kubectl get pods

workshop $ kubectl get services kubers

workshop $ kubectl exec kubers-6n4kz -- curl -s http://10.96.206.237


```

Running commands inside container.

```shell


workshop $ kubectl exec kubers-6n4kz env 

workshop $ kubectl exec kubers-6n4kz bash

```


## Load Balancer

Services can be exposed to external client by using Load Balancer. LoadBalancer redirects traffic to the node port across all the nodes. Clients can connect to the service through the load balancer's IP.


```shell

workshop $ kubectl create -f examples/loadbalancer.yaml

workshop $ kubectl get svc kubers-loadbalancer

workshop $ curl localhost


```






## Delete what we have created


Delete ReplicaSet

```shell

workshop $ kubectl delete rs kubers
workshop $ kubectl delete svc kubers-loadbalancer

workshop $ kubectl get pods

workshop $ kubectl get rs

workshop $ kubectl get svc

```
