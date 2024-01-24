# qlub-golang-app-challange
Coding challenge for calculating max sum on binary tree.

## Used Libraries
 Viper - Manages reading, writing, and managing configuration files.
 Gin - A fast and minimalistic web framework for Go.
 Ginkgo - Provides a structure to organize your tests and helps in creating readable and understandable test scenarios.

## API Endpoints

### 1. Calculate Max Sum (/calculateMaxSum)

- **Endpoint:** `/calculateMaxSum`
- **Method:** POST
- **Example Usage:**
  ```bash
  curl --location 'localhost:8080/calculateMaxSum' \
--header 'Content-Type: application/json' \
--data '{
"tree": {
"nodes": [
{"id": "1", "left": "2", "right": "3", "value": 1},
{"id": "3", "left": null, "right": null, "value": 3},
{"id": "2", "left": null, "right": null, "value": 2}
],
"root": "1"
}
}'


## Run Locally

### Option 1
Add `SPEC_FILE_PATH="../../config"` ENV variable in anyway and run `main.go` file.

### Option 2 - Docker Container 
- Build the image

`docker build -t qlub-calculator-api:latest . || make build`

- Run the container

`docker run -p 8080:8080 qlub-calculator-api:lates || make run`

## Run Unit Test
You can run the unit test with coverage using the `make unit-test-with-coverage` command.