FROM golang:1.19
RUN mkdir /build
ADD .. /build/
WORKDIR /build
COPY .. /build
RUN go build -tags=jsoniter -o crud -mod vendor cmd/main.go
EXPOSE 8080
CMD sh /build/start.sh
