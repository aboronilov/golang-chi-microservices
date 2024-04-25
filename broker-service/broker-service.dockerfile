FROM alpine:3.19

RUN mkdir /app

COPY brokerApp /app

CMD ["/app/brokerApp"]