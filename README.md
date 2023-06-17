# Docker

Build with:

```sh
docker build -f Dockerfile-api -t app/api .
```

Run with

```sh
docker run --rm -it -p 3080:3080 --name api app/api
```

Healthiness

```sh
curl -v localhost:3080/healthz
```

# K8s

## Push New Image

create new commit
```sh
git commit -a -m "message"
```

Create new image

```sh
TAG=$(git rev-parse HEAD)
echo $TAG
docker build -f Dockerfile-api -t app/api --build-arg GIT_COMMIT=${TAG} .
```

Push to a container registry

```sh
docker tag app/api gmhafiz/api:${TAG}
docker push gmhafiz/api:${TAG}
```
Migration needs to happen first

```sh
TAG=$(git rev-parse HEAD)
docker build -f Dockerfile-migrate -t app/migrate --build-arg GIT_COMMIT="${TAG}" .
docker tag app/migrate gmhafiz/migrate:${TAG}
docker push gmhafiz/migrate:"${TAG}"
```

## Apply

```sh
cd db1
kubectl apply -f configmap.yaml
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```


```sh
k get po
```

```sh
k logs -f postgres-7c6b976c95-wjls2
```

```
PostgreSQL Database directory appears to contain a database; Skipping initialization

2023-05-22 02:45:53.640 UTC [1] LOG:  starting PostgreSQL 15.1 (Debian 15.1-1.pgdg110+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
2023-05-22 02:45:53.641 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
2023-05-22 02:45:53.641 UTC [1] LOG:  listening on IPv6 address "::", port 5432
2023-05-22 02:45:53.646 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
2023-05-22 02:45:53.653 UTC [27] LOG:  database system was shut down at 2023-05-22 02:40:23 UTC
2023-05-22 02:45:53.659 UTC [1] LOG:  database system is ready to accept connections
```

# Port Forward

Connect

```sh
export POSTGRES_PASSWORD=$(kubectl get cm --namespace default db-secret-credentials -o jsonpath="{.data.POSTGRES_PASSWORD}")

kubectl run postgresql-dev-client --rm --tty -i --restart='Never' --namespace default --image postgres:15.3 --env="PGPASSWORD=$POSTGRES_PASSWORD" \
      --command -- psql --host postgres -U app1 -d app_db -p 5432
```

Port Forward

```sh
kubectl port-forward --namespace default svc/postgres 45432:5432 &
    PGPASSWORD="$POSTGRES_PASSWORD" psql --host 127.0.0.1 -U app1 -d app_db -p 5432
```

## Migrate

Apply secret for our app

```sh
k apply -f secret.yaml
```

Assign the values to environment variables

```sh
export DB_HOST=$(kubectl get secret --namespace default postgres-secret-config -o jsonpath="{.data.host}" | base64 -d)
export DB_PORT=$( kubectl get secret --namespace default postgres-secret-config -o jsonpath="{.data.port}" | base64 -d)
export DB_NAME=$( kubectl get secret --namespace default postgres-secret-config -o jsonpath="{.data.name}" | base64 -d)
export DB_USER=$( kubectl get secret --namespace default postgres-secret-config -o jsonpath="{.data.user}" | base64 -d)
export DB_PASS=$( kubectl get secret --namespace default postgres-secret-config -o jsonpath="{.data.password}" | base64 -d)
```

Run a one-off Pod

```sh
kubectl run api-migrate --stdin --tty --rm --restart=Never --namespace default --image gmhafiz/migrate:0c0765c78b308b31362b9c6093e7fc254be98d72 --env="DB_HOST=$DB_HOST" --env="DB_PORT=$DB_PORT" --env="DB_NAME=$DB_NAME" --env="DB_USER=$DB_USER" --env="DB_PASS=$DB_PASS" --command -- migrate
```

Returns
```
2023/05/22 08:55:19 starting migrate...
2023/05/22 08:55:19 reading env
2023/05/22 08:55:19 OK    01_random.sql
2023/05/22 08:55:19 goose: no migrations to run. current version: 1
2023/05/22 08:55:19 goose: version 1
pod "api-migrate" deleted

```

## API Server 

```sh
kubectl port-forward deployment/server 3080:3080
```

test

```sh
curl -v http://localhost:3080/healthz
curl -v http://localhost:3080/ready
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
    "id": 1,
    "name": "d1cf935e-803f-4bf0-a757-f38e649304e2"
  },
  {
    "id": 2,
    "name": "8968275a-f91b-485a-ae93-55e301343461"
  },
  {
    "id": 3,
    "name": "782919e3-1927-48b1-b3b8-ba7d85ce15a1"
  },
  {
    "id": 4,
    "name": "4de3f3d5-f8d8-4ecf-b3a6-903f90fea32f"
  },
  {
    "id": 5,
    "name": "72a03a4c-f028-45d5-a952-992115b42f6a"
  },

and so on
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

It is on 1 replica at 500MHz and 128MB ram each. Let us ramp up the number of replicas to 4.

While having the chart opened, run the following

```sh
k apply -f server.yaml
```