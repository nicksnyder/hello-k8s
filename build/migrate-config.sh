#!/bin/bash
# PersistentVolume (PV)
# PersistentVolumeClaim (PVC)

set -ex

FILTER="${FILTER:-gitserver}"

# Select all PVs that match the filter and make sure they have a retain policy.

# Pre-requisite
# kubectl get pv -o json | jq --raw-output  ".items | map(select(.spec.claimRef.name | contains(\"$FILTER\"))) | .[] | \"kubectl patch pv -p '{\\\"spec\\\":{\\\"persistentVolumeReclaimPolicy\\\":\\\"Retain\\\"}}' \\(.metadata.name)\"" | bash

# Get the PV associated with the PVC 
$PV=$(kubectl get pv -o json | jq --raw-output ".items | map(select(.spec.claimRef.name == '$PVC')) | first | .metadata.name")

# Make sure reclaim policy is "retain" (it should already be due to pre-requisite).
kubectl patch pv -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}' $PV

# Delete the PVC bound to the PV
#kubectl delete pvc $PVC

# Update PV claim so it can be reused by new PVC in
kubectl patch pv -p '{"spec":{"claimRef":{"uid":null}}}' $PV
