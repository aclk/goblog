FROM alpine:latest

EXPOSE 6767

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --update add \
    curl \
	tzdata \
	ca-certificates \
	&& rm -rf /var/cache/apk/*

ENV TZ=Asia/Shanghai

ADD bin/accountservice-linux-amd64 /

HEALTHCHECK --interval=10s --timeout=5s --start-period=10s --retries=5 CMD curl -sSf http://127.0.0.1:6767/health
ENTRYPOINT ["./accountservice-linux-amd64"]
