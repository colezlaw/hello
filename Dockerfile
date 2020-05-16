FROM golang:alpine AS build-env
WORKDIR /app
COPY . /app
RUN go build -o hello

FROM alpine
COPY --from=build-env /app/hello /app/hello
WORKDIR /app
EXPOSE 8080
ENTRYPOINT [ "/app/hello" ]