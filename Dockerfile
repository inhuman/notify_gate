FROM alpine
LABEL maintainer="msgexec@gmail.com"

RUN mkdir -p /opt/notify-gate/templates

COPY ./notify-gate /opt/notify-gate/
COPY templates/index.html /opt/notify-gate/templates

RUN chmod +x /opt/notify-gate/notify-gate
RUN apk update && apk add ca-certificates

ENTRYPOINT /opt/notify-gate/notify-gate