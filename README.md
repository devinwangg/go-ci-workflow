# go ci workflow
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![CI](https://github.com/devinwangg/go-github-action/workflows/CI/badge.svg)](https://github.com/devinwangg/go-github-action/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/devinwangg/go-ci-workflow)](https://goreportcard.com/report/github.com/devinwangg/go-ci-workflow)

## Project Description
This repository showcases a Continuous Integration (CI) workflow for a Golang project using GitHub Actions and Codecov. The CI pipeline automatically builds the code, runs unit tests, and reports code coverage metrics via Codecov.

## Features
- Automated Code Build
  Utilizes GitHub Actions for automated code builds.
- Automated Unit Tests
  Automated unit tests to ensure code quality.
- Code Coverage Reporting
  Reports code coverage via Codecov.
- Easy-to-Follow Configuration
  Contains easy-to-follow GitHub Action YAML configuration files.
- Multi-Version Support
  Supports multiple versions using semantic versioning.

## Workflow
This repository demonstrates the implementation of four workflow jobs:

1. Unit Testing: Automatically runs unit tests to ensure code functionality.
2. Golang Style Checker: Uses the `golangci-lint` tool to check code style and quality.
3. Unit Testing Coverage Badge: Generates a badge indicating the coverage percentage of unit tests.
4. Semantic Versioning: Implements semantic versioning to manage versions of the project.
5. Security Scanning: Uses the `gosec` tools to scan golang source code for security problems 

## Pre-requisites
- Github Account Token
- Codecov Account Token

## Usage
1. Fork and Clone the Repository
   ```
    git clone git@github.com:devinwangg/go-ci-testing-workflow.git
   ``` 
2. Navigate to Project Directory
    ```
    cd go-ci-testing-workflow
   ```

3. Run Tests Locally
    ```
    make test
   ```

4. Generate Github Token and Codecov Token to Github Action secret
5. Push your changes to the GitHub repository
6. Check Code Coverage on Codecov
   Once the workflow is complete, check the code coverage report on your Codecov dashboard.

## GitHub Actions Configuration
The main configuration for the GitHub Actions workflow is stored in `.github/workflows` folder. Here you'll find the steps for building the code and running tests.



