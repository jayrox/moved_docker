FROM golang as compiler 
RUN CGO_ENABLED=0 go get -a -ldflags '-s' \ 
github.com/jayrox/moved_docker 

FROM scratch 
COPY --from=compiler /go/bin/moved_docker . 

VOLUME ["/mnt/src", "/mnt/dst"]

CMD ["./moved_docker"]
