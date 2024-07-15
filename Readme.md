# Microservices Project

This project demonstrates a simple microservices architecture using Go, Docker, and Kubernetes. It includes a backend service and a frontend service.

## Table of Contents

- [Microservices Project](#microservices-project)
  - [Table of Contents](#table-of-contents)
  - [Project Overview](#project-overview)
  - [Architecture](#architecture)
  - [Setup and Deployment](#setup-and-deployment)
    - [Using Minikube](#using-minikube)
    - [Service Explanation](#service-explanation)

## Project Overview

The goal of this project is to illustrate the deployment of microservices using Docker and Kubernetes. The project consists of two services:

- **Backend Service**: A simple Go application that responds with a greeting message.
- **Frontend Service**: A simple Go application that forwards requests to the backend service and displays the response.

## Architecture

The architecture of the project is as follows:


1. **Backend Service**:
    - Written in Go.
    - Provides a REST endpoint that responds with a greeting message.
    - Packaged into a Docker container.

2. **Frontend Service**:
    - Written in Go.
    - Provides a REST endpoint that forwards requests to the backend service.
    - Packaged into a Docker container.

3. **Kubernetes Deployment**:
    - Both services are deployed as separate deployments in a Kubernetes cluster.
    - Backend service is exposed internally within the cluster.
    - Frontend service is exposed externally to the internet.

4. **Nginx Load Balancer**:
    - Deployed as a service in the Kubernetes cluster.
    - Load balances traffic between multiple instances of the frontend service.

## Setup and Deployment

### Using Minikube

1. **Start Minikube**:
   Minikube is a tool that lets you run Kubernetes locally. Start Minikube by running:
   ```sh
   minikube start
2. **Build Docker Images Locally**:
    ```sh 
    docker build -t iasonasxrist/backend:latest ./backend
    docker build -t iasonasxrist/no-app:latest ./frontend

3. **Load Docker Images into Minikube**:
   ```sh
   minikube image load iasonasxrist/backend:latest
   minikube image load iasonasxrist/no-app:latest

4. **Apply Kubernetes Manifests**:
   ```sh
   kubectl apply -f backend-deployment.yaml
   kubectl apply -f frontend-deployment.yaml
5. **Access the Frontend Service**:
    ```sh
    minikube service frontend --url

## Service Explanation
Explanation of Kubernetes Services
When you deploy services in Kubernetes, they are accessible within the cluster through different types of services. Below is an explanation of the listed services:


<img width="832" alt="image" src="https://github.com/user-attachments/assets/a947f46f-3fad-4a5b-b147-b264a4dc2399">


# backend
Name: backend
Type: ClusterIP
Cluster IP: 10.108.65.182
External IP: <none>
Ports: 3000/TCP
Age: 12m

Explanation:
- Type: ClusterIP: This service is only accessible within the Kubernetes cluster. It does not have an external IP address, meaning it cannot be accessed from outside the cluster.
- Cluster IP: The internal IP address of the service within the cluster. Other services within the cluster can use this IP address to communicate with the backend service.
- Ports: This service listens on port 3000 using the TCP protocol. Other services within the cluster can communicate with the backend service on this port.

# commandservice-clusterip-srv
Name: commandservice-clusterip-srv
Type: ClusterIP
Cluster IP: 10.103.202.29
External IP: <none>
Ports: 80/TCP
Age: 33d

Explanation:
- Type: ClusterIP: This service is also only accessible within the Kubernetes cluster.
- Cluster IP: The internal IP address of the service within the cluster. Other services within the cluster can use this IP address to communicate with the commandservice-clusterip-srv service.
- Ports: This service listens on port 80 using the TCP protocol. Other services within the cluster can communicate with the commandservice-clusterip-srv service on this port.

# frontend
Name: frontend
Type: LoadBalancer
Cluster IP: 10.103.63.94
External IP: localhost
Ports: 80:32485/TCP
Age: 115s

Explanation:
- Type: LoadBalancer: This service is exposed externally, allowing access from outside the Kubernetes cluster. It uses a load balancer to distribute incoming traffic to the available pods.
- Cluster IP: The internal IP address of the service within the cluster.
- External IP: localhost indicates that the service is accessible via the localhost address. This might be specific to the local Kubernetes setup (e.g., Minikube or kind), where the LoadBalancer service is mapped to localhost for external access.
- Ports: This service listens on port 80 externally and maps it to port 32485 on the node. The traffic to localhost:80 is forwarded to 32485 on the cluster, which is then directed to the appropriate pods.

# kubernetes
Name: kubernetes
Type: ClusterIP
Cluster IP: 10.96.0.1
External IP: <none>
Ports: 443/TCP
Age: 57d

Explanation:
- Type: ClusterIP: This is the default service created by Kubernetes, which is used to manage the Kubernetes cluster.
- Cluster IP: The internal IP address of the Kubernetes API server.
- Ports: This service listens on port 443 using the TCP protocol, which is the default port for HTTPS. This is used for secure communication with the Kubernetes API server.

