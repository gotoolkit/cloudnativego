FROM alpine:latest as base
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV APP_PORT 8585
ENV APP_DATASTOREPATH /data
EXPOSE $APP_PORT

VOLUME $APP_DATASTOREPATH

COPY cloudnativego /
CMD [ "/cloudnativego" ]