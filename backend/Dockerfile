FROM golang AS builder
ENV LANG=C.UTF-8 CGO_ENABLED=0 GO111MODULE=on
WORKDIR /tasks/backend
RUN go get github.com/pilu/fresh && \
  wget -q https://github.com/schemalex/schemalex/releases/latest/download/schemalex_linux_amd64.tar.gz && \
  tar -xzf schemalex_linux_amd64.tar.gz && \
  cp schemalex_linux_amd64/schemalex /go/bin/schemalex && \
  rm schemalex_linux_amd64.tar.gz && \
  rm -rf schemalex_linux_amd64
ADD go.mod go.sum /tasks/backend/
RUN go mod download
ADD . /tasks/backend/
RUN ["go", "install"]

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/bin/tasks /root/
CMD ["./tasks"]
