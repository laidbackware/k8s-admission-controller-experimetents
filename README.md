# k8s-admission-controller-experiments

Based on https://github.com/kubernetes/sample-controller

## (optional) Regenerate certs
Run `scripts/gen_certs.sh` from the root of the repo.

## Kind prep
Run `kind_setup.sh` in the root of the repo.

## Kind tidy up
Delete the kind cluster and network:
```sh
kind delete cluster -n tilt && docker network rm kind-tilt
```

## Set the Webhook
```sh
export CA_CERT="$(cat secrets/cert.crt | base64 -w 0)"
cat ${script_dir}/manifests/webhook.yaml | envsubst | kubectl apply -f -
```
