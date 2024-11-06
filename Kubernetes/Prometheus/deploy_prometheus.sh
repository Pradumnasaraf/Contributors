#!/bin/bash

# Check if Helm is installed
if ! command -v helm &> /dev/null; then
  echo "Error: Helm is not installed. Please install Helm and try again."
  exit 1
fi

# Set variables
NAMESPACE="prometheus"
RELEASE_NAME="prometheus"
VALUES_FILE="Kubernetes/Prometheus/values.yml"
HELM_REPO_NAME="prometheus-community"
HELM_REPO_URL="https://prometheus-community.github.io/helm-charts"

# Add the Helm repository if not already added
if ! helm repo list | grep -q "$HELM_REPO_NAME"; then
  echo "Adding Helm repository: $HELM_REPO_NAME..."
  helm repo add "$HELM_REPO_NAME" "$HELM_REPO_URL"
fi

# Update Helm repositories
echo "Updating Helm repositories..."
helm repo update

# Install or upgrade Prometheus
echo "Deploying Prometheus in the $NAMESPACE namespace using Helm..."
helm upgrade --install "$RELEASE_NAME" "$HELM_REPO_NAME/$RELEASE_NAME" \
  -f "$VALUES_FILE" \
  --create-namespace \
  --namespace "$NAMESPACE" \
--wait

echo "Prometheus has been deployed to the $NAMESPACE namespace."
