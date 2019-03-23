FROM golang
ENV GO111MODULE=on
WORKDIR /tasks
ADD go.mod go.sum /tasks/
RUN go mod download
ADD . /tasks/
CMD ["go", "run", "server.go"]
