FROM alpine

RUN apk add bash curl

WORKDIR /webapptest

COPY test.sh .

ENTRYPOINT [ "/webapptest/test.sh" ]
