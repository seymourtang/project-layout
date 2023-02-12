FROM debian:bullseye-slim

LABEL maintainer="tangxianmeng@agora.io"

USER nobody

WORKDIR /app

COPY bin/app .

ENTRYPOINT ["./app"]