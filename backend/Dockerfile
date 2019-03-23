FROM golang AS builder
ENV CGO_ENABLED=0 GO111MODULE=on
WORKDIR /tasks
RUN go get github.com/pilu/fresh
ADD go.mod go.sum /tasks/
RUN go mod download
ADD . /tasks/
RUN ["go", "install"]

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/bin/tasks /root/
CMD ["./tasks"]
