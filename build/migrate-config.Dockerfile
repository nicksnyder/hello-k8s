FROM alpine:3.7
RUN apk add --no-cache \
    ca-certificates \
    tini
COPY ./build/migrate-config.sh /usr/local/bin/
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["./usr/local/bin/migrate.sh"]