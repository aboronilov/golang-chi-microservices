FROM alpine:3.19

RUN mkdir /app

COPY mailServiceApp /app

CMD ["/app/mailServiceApp"]