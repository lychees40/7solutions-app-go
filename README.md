# 7solutions-app-go

## Overview

This is a simple Golang application and kustomize templates to deploy on GCP

## Prerequisites

- Docker installed on your local machine.
- Docker Hub account with a repository created.
- Docker Hub credentials added to GitHub repository secrets (with Write and Read permissions):
  - `DOCKER_USERNAME`
  - `DOCKER_PASSWORD`

## GitHub Actions Workflow

This repository uses GitHub Actions to automate the build and push process for Docker images. The workflow is triggered on pushes to the `develop` branch and on tags matching the pattern `v[1]+.[0-9]+.[0-9]+`.

## Acknowledgements

- The CI step for compiling the Go application is handled during the Docker multi-stage build process.
