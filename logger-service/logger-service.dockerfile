FROM alpine:3.19

RUN mkdir /app

COPY loggerServiceApp /app

CMD ["/app/loggerServiceApp"]