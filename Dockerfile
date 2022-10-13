FROM golang:1.18 as builder
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=1 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
#RUN apk update && apk add libc6-compat
WORKDIR /app/
RUN mkdir -p ./static/css ./templates
COPY --from=builder /build/main ./
COPY --from=builder /build/static/css/ ./static/css/
COPY --from=builder /build/templates/ ./templates/
CMD ["./main"]