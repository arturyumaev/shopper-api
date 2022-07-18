# apply namespace shopper-api
kubectl apply -f ./shopper-api/namespace.yaml
kubectl apply -f ./shopper-api/deployment.yaml
kubectl apply -f ./shopper-api/service.yaml
kubectl apply -f ./shopper-api/ingress.yaml

# apply nginx ingress controller
kubectl create namespace m
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
helm repo update
echo "Installing nginx ingress controller..."
helm install nginx ingress-nginx/ingress-nginx --namespace m -f ./nginx-ingress.yaml

# get minikube ip
echo "Add line: '$(minikube ip) arch.homework' at the bottom of your /etc/hosts file"
