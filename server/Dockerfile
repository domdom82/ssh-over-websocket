FROM ubuntu

RUN apt-get update && \
    apt-get install -y iproute2 iputils-ping curl net-tools ssh

RUN useradd testuser && \
    echo "testuser:yourpassword" | chpasswd

COPY server /server

ENTRYPOINT /server

EXPOSE 8080