FROM golang:1.18 AS builder

WORKDIR /usr/local/go/src/learning-go

COPY . .

RUN CGO_ENABLE=0 go build -o ingress-manager ./kubernetes/client-go/custom-controller/main.go

FROM alpine:3.15.3

WORKDIR /apps

COPY --from=builder /usr/local/go/bin/ingress-manager .

CMD ["./ingress-manager"]