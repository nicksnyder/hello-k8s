FROM base AS base 

# Create a minimal image with just the config service.
FROM alpine:3.10.3
RUN apk add --no-cache \
    ca-certificates \
    tini
COPY --from=base /go/bin/config /usr/local/bin/
RUN addgroup -S config && adduser -S -G config config 
USER config 
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["./usr/local/bin/config"]