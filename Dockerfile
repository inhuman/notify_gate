FROM alpine
LABEL maintainer="msgexec@gmail.com"

COPY ./notify-gate /usr/local/bin

RUN chmod +x /usr/local/bin/notify-gate
RUN apk update && apk add ca-certificates
ENTRYPOINT /usr/local/bin/notify-gate