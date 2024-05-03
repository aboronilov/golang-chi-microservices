FROM alpine:3.19

RUN mkdir /app

COPY listenerApp /app

CMD ["/app/listenerApp"]