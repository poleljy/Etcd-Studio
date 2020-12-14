FROM centos:centos7

WORKDIR /usr/local/bin
COPY ./EtcdStudio /usr/local/bin/


EXPOSE 8080

ENTRYPOINT /usr/local/bin/EtcdStudio -p 8080
