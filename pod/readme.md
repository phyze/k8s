# Pod

```sh
kubectl apply -f pod.yml
```

deploy หลาย ๆ pods

```sh
kubectl apply -f pod-list.yml
```

กำหนด resource request และ limit 

```sh
kubectl apply -f podresource.yml
```

# Clean

```sh
kubectl delete  -f pod.yml
kubectl delete -f pod-list.yml
kubectl delete -f podresource.yml
```