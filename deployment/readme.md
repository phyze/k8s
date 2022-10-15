# Deployment

```sh
kubectl apply -f gbv1.yml
```

lookup deployment

```sh
kubectl get deployment -l app=guestbook
```

lookup pods

```sh
kubectl get pods -l app=guestbook
```

rolling update 

```sh
kubectl apply -f gbv2.yml
```

lookup pods

```sh
kubectl get pod -l app=guestbook
```

lookup image version

```sh
kubectl describe deploy/guestbook -o jsonpath="{.spec.template.spec.containers[?(@.name=='guestbook')].image}"
```

## stategy deployment

```sh
kubectl apply -f stategyDeploy.yml
```

lookup 

```sh
kubectl get pod -l app=guestbook
```

## liveneess

```sh
kubectl apply -f liveness.yml
```

lookup

```sh
kubectl get pod -l app=liveness
```


## readiness

```sh
kubectl apply -f readiness.yml
```

lookup

```sh
kubectl get pod -l app=readiness
```

## liveness and readiness

```sh
kubectl apply -f readlive.yml
```

lookup

```sh
kubectl get pod -l app=readlive
```

# clean

```sh
kubectl delete -f stategyDeploy.yml
kubectl delete -f liveness.yml
kubectl delete -f readiness.yml
kubectl delete -f readlive.yml
```