#!/bin/bash

# Check if Helm is installed
if ! command -v helm &> /dev/null; then
  echo "Error: Helm is not installed. Please install Helm and try again."
  exit 1
fi

# Set variables
NAMESPACE="grafana"
RELEASE_NAME="grafana"
VALUES_FILE="Kubernetes/Grafana/values.yml"
HELM_REPO_NAME="grafana"
HELM_REPO_URL="https://grafana.github.io/helm-charts"

# Add the Helm repository if not already added
if ! helm repo list | grep -q "$HELM_REPO_NAME"; then
  echo "Adding Helm repository: $HELM_REPO_NAME..."
  helm repo add "$HELM_REPO_NAME" "$HELM_REPO_URL"
fi

# Update Helm repositories
echo "Updating Helm repositories..."
helm repo update

# Install or upgrade Grafana
echo "Deploying Grafana in the $NAMESPACE namespace using Helm..."
helm upgrade --install "$RELEASE_NAME" "$HELM_REPO_NAME/grafana" \
  -f "$VALUES_FILE" \
  --create-namespace \
  --namespace "$NAMESPACE" \
  --wait

echo "Grafana has been deployed to the $NAMESPACE namespace."
