## How to apply manifests

This directory contains manifests for Kubernetes workshop cluster.

To deploy these resources to Kubernetes, you need to use [kontemplate](https://github.com/tazjin/kontemplate)

```shell
cp dex-secrets.example.yml dex-secrets.yml
# Add the secrets

# Ensure you're using the k8s-demo cluster!
kubectx

# Apply manifests..
kontemplate apply k8s-demo-cluster.yml