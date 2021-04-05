############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Export target arch arg
ARG TARGETARCH
# Install git to get go deps
RUN apk update && apk add --no-cache curl gcc musl-dev git ca-certificates tzdata && update-ca-certificates
# Set GO Build env var
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux
# Create appuser.
ENV USER=appuser
ENV UID=10001 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"
# Create build dir
WORKDIR /build
# COPY go.mod and go.sum files to the workspace
COPY go.* ./
# Caching vendored deps
RUN go mod download
# Check their integrity
RUN go mod verify
# COPY the source code as the last step
COPY . .
# Build the binary (dynamically linked).
RUN go build -ldflags="-w -s" -o bin/server .
############################
# STEP 2 build a small image
############################
FROM golang:alpine
WORKDIR /go
# Import missing files in scratch image like the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy dynamic lib
COPY --from=builder /lib/*.so.1 /lib/
ENV LD_LIBRARY_PATH=/lib/    
# Copy our static executable.
COPY --from=builder /build/bin/server /go/bin/server
# Create storage folders
RUN mkdir ./data || echo "directory already exist"
RUN mkdir ./data/documents || echo "directory already exist"
RUN mkdir ./data/photos || echo "directory already exist"
RUN mkdir ./data/models3d || echo "directory already exist"
RUN chown -R appuser:appuser ./data
# Use an unprivileged user.
USER appuser:appuser
# Run the server binary.
ENTRYPOINT ["/go/bin/server"]
