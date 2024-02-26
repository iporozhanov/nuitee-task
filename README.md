# Nuitee LiteAPI


## Table of Contents

- [Description](#description)
- [Usage](#usage)
  - [Running Locally](#running-locally)
  - [Running with Docker](#running-with-docker)
  - [Run tests](#run-tests)
  - [Authenticate](#authenticate)
- [ApiDoc](openapi.yaml)

## Description

This app fetches the cheapest hotel rates from Hotelbeds API based on availability and required occupancies. It also converts the exchange rate to the desired currency (using the Coinbase APi to get the rates and cache them).

[More information on the enpoint](./openapi.yaml)

## Usage

### Running Locally
1. Setup  file

   ```shell
   cp config.yaml.example config.yaml
   ```
And fill it with the correct settings

2. Run it

   ```shell
   make run
   ```

### Running with Docker

1. Build the Docker image:

   ```shell
   make API_PORT=8080 build-docker
   ```

2. Run the Docker container with docker compose:

   ```shell
   make API_PORT=8080 run-docker
   ```
Notice: The api port you pass to the above commands should match the on in the config.yaml

### Run tests

   ```shell
   make test
   ```