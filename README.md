# Helm Microservice Project

This project demonstrates the deployment of two microservices (`greet-service` and `user-service`) using **Helm** on **Kubernetes**. It showcases how Helm can simplify the management of Kubernetes resources, such as deployments and services, while also allowing for easy configuration and versioning of microservices.

## 🚀 Project Overview

The project includes:

* **greet-service**: A Go-based microservice that greets users via a `/greet` endpoint.
* **user-service**: A Go-based microservice that returns user details via a `/user/{id}` endpoint.

Both services also expose the following endpoints:

* `/health`: A health check endpoint for Kubernetes readiness and liveness probes.
* `/metrics`: Prometheus-compatible metrics endpoint for monitoring.


### Key Helm Chart Files

* **Chart.yaml**: Metadata for the Helm chart.
* **values.yaml**: Contains default configuration values for the chart (e.g., image names, replica count).
* **templates/**: Kubernetes manifest templates (e.g., `deployment.yaml`, `service.yaml`) with Helm templating syntax.

## 🔧 Getting Started

### Prerequisites

Before you begin, ensure you have:

* A Kubernetes cluster (either local or remote).
* **Helm** installed.
* **kubectl** installed and configured to access your cluster.

### 1. Clone the Repository

```bash
git clone https://github.com/nsahil992/Helm-Microservice-Project.git
cd Helm-Microservice-Project
```

### 2. Deploying Microservices Using Helm

1. **Install `greet-service`**:

```bash
helm install greet-release ./greet-service
```

2. **Install `user-service`**:

```bash
helm install user-release ./user-service
```

### 3. Verify the Deployment

Check the deployments and services to confirm everything is running:

```bash
kubectl get deployments
kubectl get services
```

### 4. Access the Services

To access the services locally, you can use `kubectl port-forward`:

```bash
kubectl port-forward svc/greet-service 8080:8080
kubectl port-forward svc/user-service 8081:8081
```

Now, you can access the services in your browser:

* Greet Service: `http://localhost:8080/greet?name=YourName`
* User Service: `http://localhost:8081/user/{id}`


## 📦 Helm Chart Structure

* **values.yaml**: Configuration for the services (image, replica count, ports).
* **deployment.yaml**: Kubernetes Deployment manifest for each service, templated using Helm variables.
* **service.yaml**: Kubernetes Service manifest for each service, also templated.


## 📦 Packaging and Publishing Helm Charts

To package a Helm chart into a `.tgz` file:

```bash
helm package greet-service/
helm package user-service/
```

You can upload these `.tgz` files to Helm repositories like **Artifact Hub** for easy sharing and usage.


## 💬 Contact

* GitHub: [nsahil992](https://github.com/nsahil992)
* LinkedIn: [nsahil992](https://www.linkedin.com/in/nsahil992)
* Hashnode: [nsahil992](https://sahilnaik.hashnode.dev/)

---


