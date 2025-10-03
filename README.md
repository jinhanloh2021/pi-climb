<a id="readme-top"></a>
<br />
<div align="center">
<h3 align="center"><a href="https://dev.piclimb.com">Pi Climb</a></h3>
  <p align="center">
    A social media app for climbers. Share, discover, connect.
    <br />
  </p>
</div>

## About The Project

A social media application for climbers. The backend REST API is built with Go, gin-gonic and Supabase postgres. All requests are authenticated and RLS applies to relevant tables to prevent unauthorised access. Monolithic architecture.

The post feed is done algorithmically, recommending posts based on friends and on similar gym and difficulty.

### Built With

- Go
- gin-gonic
- Next.js
- Supabase
- Postgres
- Flutter

<!-- ## Getting Started

### Prerequisites

### Installation -->

## Infrastructure

![Pi Climb architecture diagram](/assets/pi-climb-architecture-diagram.png)

Pi Climb is deployed on AWS Elastic Container Service, in a single cluster, service and task. It can be scaled by increasing the number of tasks should utilisation increase.

Route 53 resolves name to ALB IP address. ALB routes traffic to appropriate container based on path rules `/api/*` to BE else FE. Traffic from 80/http is redirected to 443/https.

ECS service is on private-subnet with security group that only allows inbound traffic from ALB, which protects it from external excess. As Go BE requires Supabase connection, outbound traffic from the private-subnet is routed to the NAT gateway in the public-subnet, which has internet gateway access and allows access to external resources, while protecting internal services.

## CI/CD

On push to release branch, the build-and-publish job is triggered. Web and server are built with [dockerfile](/server/dockerfile) and pushed to Elastic Container Registry.

Deployment in next job uses ECS task definition to deploy containers with correct configurations.

View full [workflow.yaml](/.github/workflows/workflow.yaml) and ECS [task-definition.json](/task-definition.json).


## Roadmap

- [ ] MVP
- [ ] Mobile Application
- [ ] Machine learning based feed
- [ ] Ads

<p align="right">(<a href="#readme-top">back to top</a>)</p>
