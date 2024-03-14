FROM alpine:3
RUN mkdir /app
COPY ./bin/brokerApp /app

CMD ["/app/brokerApp"]