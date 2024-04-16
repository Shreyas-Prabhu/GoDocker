FROM golang:1.22.2-alpine3.18
WORKDIR /app
COPY . /app
RUN go build /app
EXPOSE 4000
ENTRYPOINT [ "./goDocker" ]