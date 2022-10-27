# configmap

apply configmap

```sh
kubectl apply -f cm.yml
```


apply nginx

```sh
kubectl apply -f nginx.yml
```

observe Environment variable that map via configmap

```sh
kubectl exec -it nginx -- bash
echo $MyKey
exit
```


forward port to nginx for observe index.html that was volume

```sh
kubectl port-forward pod/nginx 8080:80
```

call nginx

```sh
curl localhost:8080
```

# clean

```sh
kubectl delete -f cm.yml
kubectl delete -f nginx.yml
```