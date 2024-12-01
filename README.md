# 7solutions-app-go

## Overview

This project sets up the infrastructure for the 7 Solutions assignment, including networking and a GKE private cluster with Argo CD.

## Prerequisites

- Docker installed on your local machine.
- Docker Hub account with repository created.
- Docker Hub credentials added to GitHub repository secrets (Write and Read Permission):
  - `DOCKER_USERNAME`
  - `DOCKER_PASSWORD`

## GitHub Actions Workflow

This repository uses GitHub Actions to automate the build and push process for Docker images. The workflow is triggered on pushes to the `develop` branch and on tags matching the pattern `v[1]+.[0-9]+.[0-9]+`.

### Workflow Steps

1. **Checkout repository**: This step checks out the repository to the GitHub Actions runner.

## Acknowledgements

- The CI step for compiling the Go application is handled during the Docker multi-stage build process.
