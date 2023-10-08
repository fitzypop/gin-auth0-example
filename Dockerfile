FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

RUN mkdir /build

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -a -installsuffix cgo -o goapi


FROM scratch

COPY --from=builder /build/goapi /bin/goapi
COPY --from=builder /build/.env .

CMD ["chmod", "+x", "/bin/goapi"]

ENTRYPOINT ["/bin/goapi"]
