---
title: Deployment
weight: 40
menu: true
---

Deployment

Deployment is used for higher-level resource to deploy applications and update them declaratively.

Creating Deployment a ReplicaSet resource is created "underneath". Actual pods are created and managed by ReplicaSets not by Deployment directly.

```shell

workshop $ kubectl apply -f examples/deployment-frontend.yaml
service "frontend" created

workshop $ kubectl apply -f examples/deployment-backend.yaml 
service "backend" created

workshop $ kubectl get deployments
NAME       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
backend    1         1         1            1           32s
frontend   1         1         1            1           36s


workshop $ kubectl rollout status deployment frontend
deployment "frontend" successfully rolled out

workshop $ kubectl rollout status deployment backend
deployment "backend" successfully rolled out

workshop $ kubectl get pods -o wide
NAME                        READY     STATUS    RESTARTS   AGE       IP           NODE
backend-759bc76ddc-czxlp    1/1       Running   0          2m        10.1.0.182   docker-for-desktop
frontend-589b8768cf-z9x7l   1/1       Running   0          2m        10.1.0.181   docker-for-desktop

workshop $ kubectl delete deployment backend
workshop $ kubectl delete deployment frontend


```