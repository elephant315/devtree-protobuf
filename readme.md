# Device Report Service

This project is a Golang application that provides a report of all devices connected to a Linux system. 
The report is returned in protobuf format and includes information about the device type, device path, vendor ID, and product ID. 

There are two options to deliver device report: 
1. gRPC Service (TCP 50051)
2. HTTP Endpoint (TCP 8080)


## Project Structure

**cmd**: Contains the main entry point of the application.

**internal**: Contains the core application logic, adapters, domain models, and services.

**proto**: Contains the protobuf definitions and generated files.

**tests**: Contains additional test clients for the project.

## Requirements

1. Golang version 1.21.5 or similar
2. Libudev (ubuntu/debian: sudo apt install libudev-dev)

## Run, build and test

Use Makefile

STEP 1. Build server apps:
```
make test 
make build
```
Execute `build/app_http_server` or `build/app_grpc_server`

STEP 2. Build sample client apps use:
```
make build_client
```
Execute `build/app_http_client` or `build/app_grpc_client`

