#!/bin/bash

# Set an error flag
set -e

# Deploy API components
echo "Creating namespace for API and deploying API components..."
kubectl apply -f Kubernetes/API/Api.Namespace.yml
kubectl apply -f Kubernetes/API/Api.ConfigMap.yml
kubectl apply -f Kubernetes/API/Api.Deployment.yml
kubectl apply -f Kubernetes/API/Api.Ingress.yml
kubectl apply -f Kubernetes/API/Api.Secret.yml
kubectl apply -f Kubernetes/API/Api.Service.yml

# Deploy Mongo
echo "Creating namespace for Mongo and deploying Mongo components..."
kubectl apply -f Kubernetes/Mongo/Mongo.Namespace.yml
kubectl apply -f Kubernetes/Mongo/Mongo.HeadlessService.yml
kubectl apply -f Kubernetes/Mongo/Mongo.Statefulset.yml

# Deploy Redis
echo "Creating namespace for Redis and deploying Redis components..."
kubectl apply -f Kubernetes/Redis/Redis.Namespace.yml
kubectl apply -f Kubernetes/Redis/Redis.HeadlessService.yml
kubectl apply -f Kubernetes/Redis/Redis.Statefulset.yml

# Deploy Prometheus
echo "Deploying Prometheus using Helm..."
bash Kubernetes/Prometheus/deploy_prometheus.sh

# Deploy Grafana
echo "Deploying Grafana using Helm..."
bash Kubernetes/Grafana/deploy_grafana.sh

echo "All components deployed successfully!"
