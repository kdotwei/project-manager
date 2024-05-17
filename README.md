# Project Management System with Microservices

## Introduction
This coursework final project is a project management system implemented with a microservices architecture. The project is developed using Go, with Gine and GORM as the primary Go packages. The system is divided into three services: identifier, user manager, and project manager. Nginx is used as the API Gateway. All services are deployed using Docker Compose.

## Services
- **Identifier Service**: Generates and manages unique identifiers.
- **User Manager Service**: Manages user information and authentication.
- **Project Manager Service**: Manages project-related information.
- **API Gateway**: Uses Nginx as a reverse proxy and load balancer.

Architecture Graph

<img src="https://imgur.com/FVETwCp.jpg" alt="Architecture" width="70%">

## Installation
Follow these steps to install and deploy the project:

1. Clone the repository:
   ```bash
   git clone https://github.com/kdotwei/project-manager.git
   ```

2. Navigate to the services directory and deploy using Docker Compose:
   ```bash
   cd project-manager/services && docker compose up --build -d
   ```
