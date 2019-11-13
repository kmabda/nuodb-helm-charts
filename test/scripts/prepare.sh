#!/usr/bin/env bash

# set -o errexit
# set -o nounset
# set -o pipefail

# optionally set to 'nuodb' or nuodb-incubator'
# to test published artifacts in GCS-backed repos
: ${REPO_NAME:="stable"}
: ${VALUE_CLASS:="tiny"}

ME=`basename $0`
SCRIPT_DIR="$( cd "$(dirname "$0")" ; pwd -P )"

: ${SELF_ROOT:=${SCRIPT_DIR%test/scripts}}
pushd ${SELF_ROOT} >/dev/null
SELF_ROOT=`pwd`
popd >/dev/null

. ${SELF_ROOT}/test/scripts/profile.sh

values_option="--values ${SELF_ROOT}/samples/${VALUE_CLASS}.yaml"

values_overrides="--set cloud.provider=${PLATFORM}"

# =========================================================

# setup cluster scoped tiller (and storage classes)...

helm install storage-class ${REPO_NAME}/storage-class \
    --namespace kube-system \
    ${values_option} \
    ${values_overrides}

helm list --namespace kube-system

# =========================================================

kubectl create namespace ${TARGET_NAMESPACE}
kubectl config set-context $(kubectl config current-context) --namespace=${TARGET_NAMESPACE}

if [ "${PLATFORM}" == "google" ]; then
    GCP_ACCOUNT=`gcloud auth list --filter=status:ACTIVE --format="value(account)"`
    kubectl create clusterrolebinding user-admin --clusterrole=cluster-admin --user=${GCP_ACCOUNT}
elif [ "${PLATFORM}" == "azure" ]; then
    # required for the azure kubernetes dashboard...
    kubectl create clusterrolebinding kubernetes-dashboard --clusterrole=cluster-admin --serviceaccount=kube-system:kubernetes-dashboard
fi

# n.b. azure aks ships with its own install of the kubernetes dashboard
if [ ! "${PLATFORM}" == "azure" ]; then
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
    kubectl apply -f ${SELF_ROOT}/test/files/dashboard-adminuser.yaml
fi

if [ "${PLATFORM}" == "google" ]; then
    kubectl delete clusterrolebinding user-admin
fi

sleep 30

echo "note bearer token to log into dashboard:"
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')
