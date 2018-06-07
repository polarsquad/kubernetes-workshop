# kubernetes-workshop
This repository contains sources for Polar Squad [Kubernetes](https://kubernetes.io) workshop guide and material.

Clone this repository (`git clone --recursive git://github.com/polarsquad/kubernetes-workshop.git`) and start following the guide at [https://polarsquad.github.io/kubernetes-workshop](https://polarsquad.github.io/kubernetes-workshop)

## Development
** NOTE: You don't need to do these steps, [read the guide!](https://polarsquad.github.io/kubernetes-workshop) These are only for developers who would like to run the guide locally.

### Prerequisites
- Clone this repository with submodules `git clone --recursive git://github.com/polarsquad/kubernetes-workshop.git`
- Install Go
- `make setup`

### Serve
To view the guide locally, run:
```shell
make serve
```

### Build
```shell
make build
```