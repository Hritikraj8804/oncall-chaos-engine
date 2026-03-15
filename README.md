# 🚨 Oncall Chaos Engine

A **production-style DevOps demo application** that simulates random DevOps incidents through an API.
The project demonstrates a **containerized microservice architecture** using Docker, Docker Compose, Redis, and Nginx, deployed on **AWS EC2**.

This project focuses on **real DevOps workflows**, including container builds, reverse proxy architecture, service networking, and cloud deployment.

---

# 🧱 Architecture

```
Internet
   │
   ▼
Nginx (Reverse Proxy)
   │
   ▼
Go API Service
   │
   ▼
Redis (Visit Counter)
```

### Components

| Service    | Purpose                       |
| ---------- | ----------------------------- |
| **Nginx**  | Reverse proxy and entry point |
| **Go API** | DevOps incident generator     |
| **Redis**  | Stores visit counter          |

---

# ⚙️ Tech Stack

* Go (net/http)
* Redis
* Nginx
* Docker
* Docker Compose
* AWS EC2
* Docker Hub
* Distroless container image

---

# 📦 Features

* Random DevOps incident generator
* Visit counter using Redis
* Reverse proxy architecture
* Multi-container orchestration
* Multi-stage Docker build
* Distroless production container
* Cloud deployment on AWS

Example API response:

```json
{
  "incident": "terraform applied in production",
  "tool": "kubernetes",
  "fix": "check the logs",
  "visits": 21
}
```

---

# 📂 Project Structure

```
oncall-chaos-engine
│
├── app
│   ├── main.go
│   ├── handlers.go
│   ├── redis.go
│   └── data
│       └── incidents.json
│
├── docker
│   └── Dockerfile
│
├── nginx
│   └── nginx.conf
│
├── docker-compose.yml
└── README.md
```

---

# 🐳 Docker Architecture

The application uses a **multi-stage Docker build**.

### Build Stage

```
golang:1.23-alpine
↓
compile Go binary
```

### Runtime Stage

```
distroless base image
↓
run compiled binary
```

Benefits:

* Small image size
* Reduced attack surface
* Faster startup

---

# 🚀 Running Locally

### Clone the repository

```
git clone https://github.com/Hritikraj8804/oncall-chaos-engine.git
cd oncall-chaos-engine
```

### Start the stack

```
docker compose up --build
```

### Test the API

```
http://localhost/incident
```

---

# 🌐 Deployment (AWS EC2)

The application is deployed on an **AWS EC2 instance**.

### Architecture

```
Internet
   ↓
AWS EC2
   ↓
Docker Compose
   ↓
Nginx → Go API → Redis
```

### Steps

1. Launch EC2 instance (Ubuntu)
2. Install Docker and Docker Compose
3. Clone repository
4. Run

```
docker compose up -d
```

### Access

```
http://<EC2_PUBLIC_IP>/incident
```

---

# 🐳 Docker Image

The Go service is published to Docker Hub:

```
docker.io/<dockerhub-username>/oncall-chaos-engine
```

---

# 📈 Future Improvements

Planned DevOps upgrades:

* CI/CD pipeline using GitHub Actions
* Prometheus metrics
* Grafana dashboards
* Kubernetes deployment
* Infrastructure as Code (Terraform)

---

# 🎯 Learning Goals

This project demonstrates practical DevOps concepts:

* Containerized microservices
* Docker networking
* Reverse proxy architecture
* Environment configuration
* Production Docker images
* Cloud deployment workflows

---

# 👨‍💻 Author

**Hritik Raj**

DevOps | Cloud | Backend Systems

