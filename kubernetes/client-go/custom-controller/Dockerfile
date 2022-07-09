FROM golang:1.18 AS builder

WORKDIR /apps

COPY . .

RUN export GO111MODULE=on && export GOPROXY=https://goproxy.cn,direct && CGO_ENABLE=0 go build -mod=readonly -o ingress-manager main.go

FROM debian:stable-20220622

WORKDIR /apps

COPY --from=builder /apps/ingress-manager .

CMD ["./ingress-manager"]