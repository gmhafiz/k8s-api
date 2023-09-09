# Run

Fill in the settings into environment variables according to your postgres credentials

```sh
export DB_HOST=0.0.0.0
export DB_NAME=postgres
export DB_USER=postgres
export DB_PASS=
export DB_PORT=5432
```

Run migration to create the tables before running the api

```sh
go run cmd/api/migrate.go
go run cmd/api/main.go
```

Test if api is working

```sh
curl -v http://localhost:3080/healthz # tests if api is up
curl -v http://localhost:3080/ready   # tests if can connect to database
```

Returns

```
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 21 May 2023 07:54:42 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```

To test database:

```sh
curl -v http://localhost:3080/ready # and
curl -v http://localhost:3080/randoms | jq
```

Returns

```
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 21 May 2023 07:56:13 GMT
< Content-Type: text/plain; charset=utf-8
< Transfer-Encoding: chunked
< 
{ [5606 bytes data]
100  5593    0  5593    0     0  1092k      0 --:--:-- --:--:-- --:--:-- 1092k
* Connection #0 to host localhost left intact
[
  {
    "id": 100,
    "name": "e256f93d-e47e-47a8-8159-6ddd1ae12640"
  },
  {
    "id": 99,
    "name": "d91a5c34-d82e-43a9-acdb-403a3d3d040e"
  },
  {
    "id": 98,
    "name": "40f0e53f-c841-4a4e-b37e-1015123177dc"
  },
  and so on
```

# Benchmark

```sh
wrk -t2 -c8 -d 10s http://localhost:3080/randoms
```

returns

```
Running 10s test @ http://localhost:3080/randoms
  2 threads and 8 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   221.10us   67.03us   1.84ms   86.30%
    Req/Sec    17.94k   666.12    18.85k    74.26%
  360656 requests in 10.10s, 295.11MB read
Requests/sec:  35708.37
Transfer/sec:     29.22MB
```

# Load Testing

Install

```sh
sudo apt install pip
pip3 install locust
```

Run

```sh
locust -f ./locustfile.py --host=http://localhost:3080
```

Open  http://0.0.0.0:8089, run the load testing with 16 users and 1 spawn rate. Click "Start Swarming" and click on the chart.


# Kubernetes Deployment

See http://localhost:1314/blog/deploy-applications-in-kubernetes/#backend