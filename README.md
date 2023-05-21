# Docker

Build with:

    docker build -t app/api .

Run with

    docker run --rm -it -p 3080:3080 --name api app/api

Healthiness

    curl -v localhost:3080/healthz


# K8s

## Push New Image

create new commit
```sh
git commit -a -m "message"
```

Create new mage

```sh
TAG=$(git rev-parse HEAD)
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
docker push gmhafiz/migrate:"${TAG}"
```

## Database

```sh
k apply -f db/local-pv.yaml
k apply -f db/postgresql-data-claim.yaml
helm install postgresql-dev -f db/values.yaml bitnami/postgresql

k apply -f db-secret.yaml
```

## Deployment

```sh
k apply -f server.yaml
watch -n 1 kubectl get deploy
```

# Service

```sh
k apply -f server-service.yaml
watch -n 1 kubectl get svc
```

# Port Forward

Database
```sh
export POSTGRES_ADMIN_PASSWORD=$(kubectl get secret --namespace default postgresql-dev -o jsonpath="{.data.postgres-password}" | base64 -d)

export POSTGRES_PASSWORD=$(kubectl get secret --namespace default postgresql-dev -o jsonpath="{.data.password}" | base64 -d)

kubectl port-forward --namespace default svc/postgresql-dev 15432:5432 &
    PGPASSWORD="$POSTGRES_PASSWORD" psql --host 127.0.0.1 -U app1 -d app_db -p 5432
```

Change incoming port to an unused port, in this case `15432`.


Server

```sh
kubectl port-forward deployment/server 3080:3080
```

test

```sh
curl -v http://localhost:3080/healthz
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