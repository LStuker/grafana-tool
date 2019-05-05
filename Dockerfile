FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/lstuker/grafana-tool
ADD . .
RUN go build -o grafana-tool .

FROM alpine
LABEL maintainer="lucien.stuker+grafan-tool@gmail.com"
LABEL version="0.0.1"
LABEL description="A small helper for Grafana"
COPY --from=builder /go/src/github.com/lstuker/grafana-tool/grafana-tool /app/
WORKDIR /app
ENTRYPOINT ["./grafana-tool"]