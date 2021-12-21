FROM damianoneill/golang-alpine-builder
EXPOSE 4000
WORKDIR /workspace


RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
RUN apk add yasm && apk add ffmpeg
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer

WORKDIR /app
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# src code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o backend .
    RUN chmod +x backend

#COPY --from=builder /workspace/backend /backend
ENV TZ=Asia/Shanghai
ENTRYPOINT ["./backend"]