FROM base AS base 

FROM alpine:3.7
RUN apk add --no-cache ca-certificates tini
COPY --from=base /go/bin/config /usr/local/bin/
RUN addgroup -S appgroup && adduser -S -G appgroup appuser
USER appuser 
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["./usr/local/bin/config"]