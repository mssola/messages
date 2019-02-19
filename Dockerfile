FROM library/golang:1.11.5-alpine3.9
MAINTAINER Miquel Sabaté Solà <mikisabate@gmail.com>

COPY . /go/src/github.com/mssola/messages

RUN cd /go/src/github.com/mssola/messages; go install && \
    mkdir -p /srv/messages && \
    cp /go/src/github.com/mssola/messages/index.html /srv/messages/index.html && \
    rm -r /go/src

ENV MESSAGES_FILE_PATH /srv/messages
EXPOSE 3000
ENTRYPOINT ["messages"]
