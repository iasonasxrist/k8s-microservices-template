# Microservices Project

This project demonstrates a simple microservices architecture using Go, Docker, and Kubernetes. It includes a backend service and a frontend service.

## Table of Contents

- [Microservices Project](#microservices-project)
  - [Table of Contents](#table-of-contents)
  - [Project Overview](#project-overview)
  - [Architecture](#architecture)
  - [Setup and Deployment](#setup-and-deployment)
    - [Using Minikube](#using-minikube)

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
