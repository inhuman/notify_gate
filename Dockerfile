FROM centos
LABEL maintainer="msgexec@gmail.com"

RUN mkdir -p /opt/notify-gate

COPY ./notify-gate /opt/notify-gate/

RUN chmod +x /opt/notify-gate/notify-gate
RUN yum update && yum install -y ca-certificates

ENTRYPOINT /opt/notify-gate/notify-gate