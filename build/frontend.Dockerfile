FROM base AS base 

# Create a minimal image with just the frontend service.
FROM alpine:3.7
RUN apk add --no-cache \
    ca-certificates \
    tini
COPY --from=base /go/bin/frontend /usr/local/bin/
RUN addgroup -S frontend && adduser -S -G frontend frontend
USER frontend 
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["./usr/local/bin/frontend"]