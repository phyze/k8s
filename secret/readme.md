# Secret

## environment and file

```sh
kubectl apply -f secret.yml
kubectl apply -f pod.yml
```

get PASSWORD in  file that mount secret volume 

```sh
kubectl exec -it secret -- cat /etc/mysecretvolume/PASSWORD
```

get PASSWORD from environment variable

```sh
#macos or linux
kubectl exec -it secret -- env | grep mysecret 

#windows
kubectl exec -it secret -- env | findstr mysecret 
```

## secret with private registry

create secret for authentication

```sh
kubectl create secret docker-registry  k8scourse \
  --docker-server=ghcr.io/phyze/k8s \
  --docker-username=phyze \
  --docker-password=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 
```

apply container that pull image from private registry

```sh
kubectl apply -f private.yml
```

observe pod

```sh
kubectl get pod -l app=private
```

# Clean

```sh
kubectl delete -f secret.yml
kubectl delete -f pod.yml
kubectl delete -f private.yml
```