FROM alpine
LABEL maintainer="msgexec@gmail.com"

RUN mkdir -p /opt/notify-gate/api

COPY ./notify-gate /opt/notify-gate/
COPY api/index.html /opt/notify-gate/api

RUN chmod +x /opt/notify-gate/notify-gate
RUN apk update && apk add ca-certificates

ENTRYPOINT /opt/notify-gate/notify-gate