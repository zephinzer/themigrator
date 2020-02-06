FROM alpine:3.11.3
RUN apk update --no-cache \
  && apk upgrade --no-cache
COPY ./bin/themigrator_linux_amd64 /usr/bin/themigrator
ENTRYPOINT ["themigrator"]
