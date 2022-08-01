namespace=shopper-api

if [ $1 = 'up' ]
then
  kubectl apply -f ./$namespace/namespace.yaml

  # Copy ConfigMap from shopper-db namespace
  kubectl get cm postgres-config-map -n shopper-db -o yaml \
    | sed 's/namespace: shopper-db/namespace: shopper-api/' \
    | kubectl create -f -

  kubectl apply -f ./$namespace/deployment.yaml
  kubectl apply -f ./$namespace/service.yaml
  kubectl apply -f ./$namespace/ingress.yaml

  # Apply nginx ingress controller
  kubectl create namespace m
  helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
  helm repo update
  echo "Installing nginx ingress controller..."
  helm install nginx ingress-nginx/ingress-nginx --namespace m -f ./nginx-ingress.yaml

  # Get minikube ip
  echo "Add line: '$(minikube ip) arch.homework' at the bottom of your /etc/hosts file"
elif [ $1 == 'down' ]
then
  kubectl delete namespace $namespace
fi
