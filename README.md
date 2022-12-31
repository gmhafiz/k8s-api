# Docker

Build with:

    docker build -t app/api .

Run with

    docker run --rm -it -p 3080:3080 --name api app/api

Healthiness

    curl -v localhost:3080/healthz