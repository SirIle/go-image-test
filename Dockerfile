FROM scratch
EXPOSE 80
COPY go-image-test /
ENTRYPOINT ["/go-image-test"]
