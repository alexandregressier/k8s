FROM golang:1-alpine AS build
WORKDIR /app/
COPY cmd/ cmd/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v cmd/hello/hello.go

FROM scratch
WORKDIR /app/
COPY --from=build /app/hello .
EXPOSE 8180
ENTRYPOINT ["./hello"]
