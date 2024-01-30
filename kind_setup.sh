#!/bin/bash

set -u

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

networks="$(docker network inspect kind-tilt 2>&1)"
set -e
if [[ "$networks" == *"network kind-tilt not found"* ]]; then
  docker network create --gateway="100.127.0.1" --ip-range="100.127.0.0/24" --subnet="100.127.0.0/16" kind-tilt
else
  echo "kind-tilt network already exists"
fi

clusters="$(kind get clusters)"
if [[ "$clusters" == *"tilt"* ]]; then
  echo "Tilt cluster already exists"
else
  export KIND_EXPERIMENTAL_DOCKER_NETWORK=kind-tilt 
  kind create cluster -n tilt
  echo -e "\nChecking cluster pods are online\n"
  kubectl wait --namespace kube-system --for=condition=ready pod --all --timeout=90s
  sleep 5
  kubectl wait --namespace kube-system --for=condition=ready pod --selector=k8s-app=kube-proxy --timeout=90s
  kubectl wait --namespace kube-system --for=condition=ready pod --selector=k8s-app=kube-dns --timeout=90s
  kubectl wait --namespace kube-system --for=condition=ready pod --selector=k8s-app=kindnet --timeout=90s
  sleep 1
fi

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml
sleep 1
kubectl wait --namespace metallb-system --for=condition=ready pod --selector=app=metallb --timeout=90s
sleep 1

export IP_RANGE="100.127.255.2-100.127.255.250"
# cat ${script_dir}/metallb.yml | envsubst
cat ${script_dir}/deployments/metallb.yml | envsubst | kubectl apply -f -