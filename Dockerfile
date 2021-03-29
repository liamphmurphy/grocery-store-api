FROM golang:latest
LABEL maintainer="liamphmurphy@gmail.com"

# set working dir
WORKDIR $GOPATH/src/github.com/murnux/grocery-store-api

# copy files over to working dir
COPY . .

# download and install dependencies
RUN go get -d -v ./... && go install -v ./...

# compile the executable
RUN go build

# takes gin out of the default debug mode
ENV GIN_MODE=release

CMD ["grocery-store-api"]