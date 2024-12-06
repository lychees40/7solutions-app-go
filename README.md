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

## Security

To ensure the security of the application, the following steps are integrated into the CI/CD pipeline:

### Dependency Check

- **govulncheck**: This step uses the `golang/govulncheck-action@v1` to check for known vulnerabilities in the Go dependencies.
  ```yaml
  - name: govulncheck check
    uses: golang/govulncheck-action@v1
    with:
      go-version-input: 1.21
      go-package: ./...
  ```

### Static Code Analysis

- **gosec**: This step uses the `securego/gosec@master` action to perform static code analysis and identify potential security issues in the codebase.
  ```yaml
  - name: gosec check
    uses: securego/gosec@master
    with:
      args: ./...
  ```
## Acknowledgements


- The CI step for compiling the Go application is handled during the Docker multi-stage build process.
- Now `govulncheck` and `gosec` run only during build time to check for vulnerabilities and security issues. While this helps catch issues early, we recommend enhancing security by also incorporating runtime security features. For example, using Google Kubernetes Engine (GKE) features like `security_posture_mode` and `security_posture_vulnerability_mode` can help monitor and enforce security policies during runtime. Additionally, any vulnerabilities detected at runtime should trigger a feedback loop back to the CI pipeline to ensure continuous improvement and security.