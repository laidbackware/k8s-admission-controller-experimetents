# k8s-admission-controller-experiments

Based on https://github.com/kubernetes/sample-controller

## Kind prep
Run `kind_setup.sh` in the root of the repo.

## Kind tidy up
Delete the kind cluster and network:
```sh
kind delete cluster -n tilt && docker network rm kind-tilt
```