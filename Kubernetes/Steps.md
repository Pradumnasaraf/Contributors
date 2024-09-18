### Steps

Steps to deploy the application on a Kubernetes cluster.

Prerequisites:

- A Kubernetes cluster (Cloud provider)
- Helm v3
- kubectl

Make sure you are inside the `Kubernetes` directory. 

1. Install Nginx Ingress Controller

As soon you install on a cluster, it will create a LoadBalancer service on the cloud provider and will be assigned a public IP address. This IP address will be used to access the services running on the cluster.

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install --namespace ingress-nginx --create-namespace nginx-ingress ingress-nginx/ingress-nginx --set controller.publishService.enabled=true
```

2. Update the [Api.Ingress.yaml](./Api.Ingress.yaml) file with the domain name or the LoadBalancer URL to get the traffic routed to the service.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: api-ingress
  namespace: api
spec:
  ingressClassName: nginx
  rules:
  - host: <your-domain/loadbalancer url
    http:
      paths:
        - path: /
          pathType: Prefix
          backend:
            service:
              name: api-service
              port:
                number: 8080
```

4. Apply Database resources

- **Database.Namespace.yaml**: Create a namespace called `database`
- **Database.HeadlessService.yaml**: Create a headless service to expose the statefulset
- **Database.StatefulSet.yaml**: Create a statefulset to run the database. It will also create a persistent volume claim to store the data.

```bash
kubectl apply -f Database.Namespace.yaml
kubectl apply -f Database.HeadlessService.yaml
kubectl apply -f Database.StatefulSet.yaml
```

1. Apply the API resources

- **Api.Namespace.yaml**: Create a namespace called `api`
- **Api.Secret.yaml**: Create a secret to store the database url, authentication details, etc.
- **Api.ConfigMap.yaml**: Create a configmap to store the environment variables, like port, database name, etc.
- **Api.Deployment.yaml**: Create a deployment to run the API service
- **Api.Service.yaml**: Create a service ClusterIP to expose the deployment within the cluster
- **Api.Ingress.yaml**: Create an ingress to route the external traffic to the service using nginx ingress controller


```bash
kubectl apply -f Api.Namespace.yaml
kubectl apply -f Api.Secret.yaml
kubectl apply -f Api.ConfigMap.yaml
kubectl apply -f Api.Deployment.yaml
kubectl apply -f Api.Service.yaml
kubectl apply -f Api.Ingress.yaml
```

Now you can access the API service using the domain name or the LoadBalancer URL.

```
<your-domain/loadbalancer url>
<your-domain/loadbalancer url>/health
<your-domain/loadbalancer url>/query
```
