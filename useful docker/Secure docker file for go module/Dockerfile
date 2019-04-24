############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
RUN adduser -D -g '' appuser

COPY . $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/

# Fetch dependencies.

# Using go get.
RUN go get -d -v

# Using go mod.
# RUN go mod download

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/backoffice

############################
# STEP 2 build a small image
############################
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable.
COPY --from=builder /go/bin/hello /go/bin/hello

# Use an unprivileged user.
USER appuser

# Run the hello binary.
ENTRYPOINT ["/go/bin/hello"]