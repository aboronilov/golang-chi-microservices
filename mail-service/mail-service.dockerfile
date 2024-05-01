FROM alpine:3.19

RUN mkdir /app

COPY mailServiceApp /app
COPY templates /templates

CMD ["/app/mailServiceApp"]