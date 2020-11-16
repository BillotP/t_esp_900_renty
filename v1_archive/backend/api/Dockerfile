############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git to get go deps
RUN apk update && apk add --no-cache curl git ca-certificates tzdata && update-ca-certificates
# Setup Go private package
ARG CI_USER=${CI_USER}
ARG CI_TOKEN=${CI_TOKEN}
ENV CI_TOKEN=$CI_TOKEN \
    CI_USER=$CI_USER
RUN echo -e "\n\
 machine gitlab.com\n\
   login $CI_USER\n\
   password $CI_TOKEN\n\
 " >> /root/.netrc
# Set GO Build env var
ENV GOPRIVATE=gitlab.com/ddng \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
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
# Build the binary (staticaly linked).
RUN go build -ldflags="-w -s" -o bin/server .
# # Create lock file
# RUN echo "running" > .lock
############################
# STEP 2 build a small image
############################
FROM scratch
WORKDIR /go
# Import missing files in scratch image like the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /bin/cat /bin/cat
# Copy our static executable.
COPY --from=builder /build/bin/server /go/bin/server
# Setup kube health check file
# COPY --from=builder /build/.lock /tmp/.lock
# Use an unprivileged user.
USER appuser:appuser
## healthcheck
# HEALTHCHECK --interval=12s --timeout=12s --start-period=30s \  
# CMD /bin/cat /tmp/.lock
# Run the server binary.
ENTRYPOINT ["/go/bin/server"]
