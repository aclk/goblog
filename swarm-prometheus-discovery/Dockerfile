FROM alpine:latest

ADD bin/swarm-prometheus-discovery-linux-amd64 /
ENTRYPOINT ["./swarm-prometheus-discovery-linux-amd64","-network", "my_network", "-ignoredServices", "goblog_prometheus,goblog_grafana"]
