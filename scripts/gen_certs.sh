#!/bin/bash

repo_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )/.."

cert_dir="${repo_dir}/secrets"

echo -e "Creating certificates in ${cert_dir}.\nDirectory is ignored by git by default.\n"

mkdir -p "${cert_dir}"

CONFIG=$(cat << EOF
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[ req_distinguished_name ]
[ v3_req ]
subjectAltName=@alt_names
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth

[ alt_names ]
DNS.1 = lb-validator
DNS.2 = lb-validator.lb-validator
DNS.3 = lb-validator.lb-validator.svc
DNS.4 = lb-validator.lb-validator.svc.cluster.local
EOF
)

openssl genrsa -out "${cert_dir}/cert.key" 2048
openssl req -new -key "${cert_dir}/cert.key" -out "${cert_dir}/cert.csr" -subj "/CN=lb-validator.lb-validator.svc" -config <(echo -n "$CONFIG")
openssl x509 -req -days 3650 -in "${cert_dir}/cert.csr" -signkey "${cert_dir}/cert.key" -out "${cert_dir}/cert.crt"