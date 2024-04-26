FROM alpine:3.19

RUN mkdir /app

COPY authApp /app

CMD ["/app/authApp"]