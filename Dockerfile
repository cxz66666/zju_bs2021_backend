FROM golang as builder
EXPOSE 4000
WORKDIR /workspace
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# src code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o backend .

FROM alpine:latest
COPY --from=builder /workspace/backend /backend
RUN chmod +x /backend
ENV TZ=Asia/Shanghai
ENTRYPOINT ["/backend"]