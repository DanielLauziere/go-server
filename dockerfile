FROM scratch
 COPY ./go-server /go-server
 ENTRYPOINT ["/go-server"]

