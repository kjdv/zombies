FROM debian:testing

RUN apt-get update && \
    apt-get install -y \
    golang && \
    apt-get clean

# print go version, should be at least 1.10
RUN go version

# go is opinionated about the directory structure:
ENV GOPATH /go
WORKDIR /go/src/github.com/klaasjacobdevries/zombies
COPY . /go/src/github.com/klaasjacobdevries/zombies

RUN go build
RUN go test ./...
RUN go test -v

CMD ./zombies
