FROM golang:1
WORKDIR /usr/src/app
COPY . .
WORKDIR /usr/src/app/epbf
COPY ./ebpf .
RUN go build -o ebpf .

ENTRYPOINT ["./ebpf"]

