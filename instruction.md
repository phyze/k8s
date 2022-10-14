# Pod 

kubectl apply -f pod.yml   
kubectl apply -f pod-list.yml

### clean
kubectl delete -f pod.yml 
kubectl delete -f pod-list.yml

# Limit pod resource

kubectl apply -f pod_resource.yml

ดู resource ที่กำหนด   
kubectl describe pod myapp

### clean
kubectl delete -f pod.yml

# Deployment 

kubectl apply -f deployment.yml

### clean
kubectl delete -f deployment.yml