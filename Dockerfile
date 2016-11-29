# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/hectorgool/gomicrosearch

# Setting up working directory
WORKDIR /go/src/github.com/hectorgool/gomicrosearch

# Get godeps for managing and restoring dependencies
ENV ELASTICSEARCH_HOSTS=elasticsearch:9200

ENV GOROOT=/usr/local/go

# Restore godep dependencies
RUN go get ./...

# Build the gomicrosearch command inside the container.
RUN go install github.com/hectorgool/gomicrosearch

# Run the gomicrosearch command when the container starts.
ENTRYPOINT /go/bin/gomicrosearch

# Service listens on port 8080.
EXPOSE 8080