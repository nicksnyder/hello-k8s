kubectl get pv -o json | jq --raw-output  ".items | map(select(.spec.claimRef.name | contains(\"config-\"))) | .[] | \"kubectl patch pv -p '{\\\"spec\\\":{\\\"persistentVolumeReclaimPolicy\\\":\\\"Retain\\\"}}' \\(.metadata.name)\"" | bash
kubectl get deployment -o json | jq --raw-output ".items | map(select(.metadata.name | contains(\"config-\"))) | .[] | \"kubectl delete deployment \\(.metadata.name)\"" | bash
kubectl get pvc -o json | jq --raw-output ".items | map(select(.metadata.name | contains(\"config-\"))) | .[] | \"kubectl delete pvc \\(.metadata.name)\"" | bash
kubectl get pv -o json | jq --raw-output ".items | map(select(.spec.claimRef.name | contains(\"config-\"))) | .[] | \"kubectl patch pv -p '{\\\"spec\\\":{\\\"claimRef\\\":{\\\"uid\\\":null,\\\"name\\\":\\\"config-file-\\(.spec.claimRef.name)\\\"}}}' \\(.metadata.name)\"" | bash