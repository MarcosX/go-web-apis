FROM alpine:3
RUN mkdir /app
COPY ./authApp /app

CMD ["/app/authApp"]