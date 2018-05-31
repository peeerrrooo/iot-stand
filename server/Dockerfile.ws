FROM golang

WORKDIR /opt

COPY ws-service .

RUN chmod +x ws-service

EXPOSE 9120

CMD ["/opt/ws-service"]