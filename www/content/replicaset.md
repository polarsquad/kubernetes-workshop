---
title: ReplicaSet, Services and LoadBalancer
weight: 30
menu: true
---

In this example we have Node.js web application which accept's HTTP requests and responses with the hostname of the machine it is running in.

The directory [apps/kubers/](https://github.com/polarsquad/kubernetes-workshop/tree/master/apps/kubers) in the workshop Git repo contains the Node.JS source code for the apps, Dockerfile to build the Docker images, and Kubernetes deployment configuration files. You can find pre-built images from [Polar Squad Docker Hub](https://hub.docker.com/r/polarsquad/kubers/)



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

ReplicaSet is controlling 3 Pods but Pods are not yet available for external use.


## Services

Service is a resource created for example to make single point of entry to group of pods providing same frontend services.


```shell

workshop $ kubectl create -f examples/service.yaml

workshop $ kubectl get svc


```

Service is not yet available to external clients. Curl fails.

```shell

workshop $ curl localhost:8080
curl: (7) Failed to connect to localhost port 8080: Connection refused

```

But service is running and you can check that requests are possible to make within cluster. Curl command is performed inside Pod towards Service. Service redirects connection to frontend Pod and http response is given back by the Pod.

Remember to check correct IP from Services and correct Pod id:


```shell

workshop $ kubectl get pods

workshop $ kubectl get services kubers

workshop $ kubectl exec kubers-6n4kz -- curl -s http://10.96.206.237


```


## Load Balancer

Services can be exposed to external client by using Load Balancer. LoadBalancer redirects traffic to the node port across all the nodes. Clients can connect to the service through the load balancer's IP.


```shell

workshop $ kubectl create -f examples/loadbalancer.yaml

workshop $ kubectl get svc kubers-loadbalancer

workshop $ curl localhost
You've hit pod kubers-2t9nx

workshop $ curl localhost
You've hit pod kubers-l57z4


```

External access through LoadBalancer should work now and when sending request to your service you should get responces from all three pods.




## Delete what we have created


Delete ReplicaSet

```shell

workshop $ kubectl delete rs kubers
workshop $ kubectl delete svc kubers
workshop $ kubectl delete svc kubers-loadbalancer

workshop $ kubectl get pods

workshop $ kubectl get rs

workshop $ kubectl get svc

```
