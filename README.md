# GoSortServer

# Go Server for Sorting Arrays

This Go server implements two endpoints, `/process-single` and `/process-concurrent`, to demonstrate sequential and concurrent sorting of arrays.

## Server Setup

- The Go server listens on port 8000.
- `/process-single` endpoint for sequential processing.
- `/process-concurrent` endpoint for concurrent processing.

## Input Format

JSON payload with the following structure:

```json
{
  "to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
}
Task Implementation
/process-single: Sorts each sub-array sequentially using the sort package.
/process-concurrent: Sorts each sub-array concurrently using goroutines and channels.
Response Format
JSON response with the following structure:

json

{
  "sorted_arrays": [[sorted_sub_array1], [sorted_sub_array2], ...],
  "time_ns": "<time_taken_in_nanoseconds>"
}
Performance Measurement
Time taken to sort all sub-arrays in each endpoint is measured in nanoseconds using the time package.

Dockerization
Dockerfile is provided for building the Docker image.
Docker image is pushed to Docker Hub.
Running the Server
Build the Docker image:
bash

docker build -t <your-username>/go-sort-server .
Run the Docker image:
bash

docker run -p 8000:8000 <your-username>/go-sort-server
Send requests to the server:
Sequential:
bash

curl -X POST http://localhost:8000/process-single -H "Content-Type: application/json" -d '{"to_sort":[ [1, 2, 3], [4, 5, 6], [7, 8, 9]]}'
Concurrent:
bash

curl -X POST http://localhost:8000/process-concurrent -H "Content-Type: application/json" -d '{"to_sort":[ [1, 2, 3], [4, 5, 6], [7, 8, 9]]}'
The server will respond with a JSON object containing the sorted arrays and the execution time.
