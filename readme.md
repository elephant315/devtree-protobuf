# Device Report Service

This project is a Golang application that provides a report of all devices connected to a Linux system. 
The report is returned in protobuf format and includes information about the device type, device path, vendor ID, and product ID. 

There are two options to deliver this service via:
1. gRPC Service
2. HTTP Endpoint


## Project Structure

cmd: Contains the main entry point of the application.
internal: Contains the core application logic, adapters, domain models, and services.
proto: Contains the protobuf definitions and generated files.
tests: Contains additional tests clients for the project.


## Run, build and test

Use Makefile
Example:
```
make build
make run
make protobuf
make test 
```