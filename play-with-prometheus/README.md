For running prometheus:

`$ docker run --name prometheus -d -p 127.0.0.1:9090:9090 prom/prometheus`

or you can use my docker compose

port 2112 - my app which created own metric

docker-compose run two app, postgres, redis, cadvisor.  One of the app uses github.com/dlmiddlecote/sqlstats for adding some SQL stat to Prometheus.

