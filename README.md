# go ci workflow
[![CI](https://github.com/devinwangg/go-github-action/workflows/CI/badge.svg)](https://github.com/devinwangg/go-github-action/workflows/CI/badge.svg)

## Project Description
This repository showcases a Continuous Integration (CI) workflow for a Golang project using GitHub Actions and Codecov. The CI pipeline automatically builds the code, runs unit tests, and reports code coverage metrics via Codecov.

## Features
Automated code build using GitHub Actions
Automated unit tests
Code coverage reporting via Codecov
Easy-to-follow GitHub Action YML configuration files
Support for multiple Go versions

## Pre-requisites
- Github Account Token
- Codecov Account Token

# Usage
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



