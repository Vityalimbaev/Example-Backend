FROM harbor.alfatell.ru/dockerhub/library/golang:1.17.2 AS builder
COPY ./app /go/src/
WORKDIR /go/src/
RUN go mod vendor

RUN cd cmd/main/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o archive-backend .

FROM harbor.alfatell.ru/dockerhub/library/alpine:latest
RUN apk --no-cache add ca-certificates && mkdir -p /opt/Archive-Backend/db/migrations
WORKDIR /opt/Archive-Backend
COPY /db/migrations/ /opt/Archive-Backend/db/migrations
COPY --from=builder /go/src/cmd/main/archive-backend .
COPY --from=builder /go/src/.env .

ENTRYPOINT ["./archive-backend"]