# Namespace and quota


## namespace

create namespace 

```sh
kubectl apply -f ns.yml -n training
```

## resource quota

```sh
kubectl apply -f priority.yml
kubectl apply -f rq.yml -n training
```

ตรวจสอบ  resource quota

```sh
kubectl describe resourcequota -n training
```

deploy app

```sh
kubectl apply -f bookstore-h.yml
kubectl apply -f bookstore-m.yml
kubectl apply -f bookstore-l.yml
```

## limit resource 

```sh
kubectl apply -f limit.yml
```

```sh
kubectl apply -f pod-over-resource.yml
```