# Deployment


##  rolling update 

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

update version

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

get history of replicaSet that have rolling update 

```sh
kubectl get rs -l app=guestbook
```

## rollback

deploy new version that make application crashing

```sh
kubectl apply -f crashapp.yml
```

lookup crash app

```sh
kubectl get pod -l app=guestbook
```

### roll back with command

```sh
kubectl rollout undo deploy/guestbook
```

or using --to-revision this way you can specific number of revision to rollback

```sh
kubectl rollout history deploy/guestbook
```

specific revision

```sh
kubectl rollout --to-revision=1 deploy/guestbook
```

### roll back with config yml

```sh
kubectl apply -f gbv2.yml
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