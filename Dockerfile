FROM google/golang:latest
MAINTAINER Juan Quintana <juan.quintana@luxuriem.com>
ADD . /go/src/github.com/juanjoqmelian/go-rest/users
RUN go get github.com/julienschmidt/httprouter
RUN go get gopkg.in/mgo.v2
RUN go install github.com/juanjoqmelian/go-rest/users
ENTRYPOINT /go/bin/users
