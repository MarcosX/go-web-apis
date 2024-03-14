FROM alpine:3
RUN mkdir /app
COPY ./bin/authApp /app

CMD ["/app/authApp"]