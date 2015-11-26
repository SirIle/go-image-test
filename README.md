# go-image-test, a test service

This simple program is meant to be used to demonstrate scaling using Docker containers. When called it returns a ying/yang image with differing colors based on the hash of the hostname (container) that serves the image.

For building the container [golang-builder](https://github.com/CenturyLinkLabs/golang-builder) is used. The container can be built (assuming you have a working Docker environment) with:

```
docker run --rm   -v "$(pwd):/src"   -v /var/run/docker.sock:/var/run/docker.sock   centurylink/golang-builder
```

If you want to tag the image straight after the build the command is

```
docker run --rm   -v "$(pwd):/src"   -v /var/run/docker.sock:/var/run/docker.sock   centurylink/golang-builder sirile/go-image-test
```

After that running the container can be done with

```
docker run --rm -p 80:80 go-image-test
```

Testing the service can be done with

```
curl http://<ip_of_docker_node>
```
