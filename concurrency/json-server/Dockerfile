FROM node:12-alpine
LABEL Maintainer=github.com/rnkoaa

RUN npm install -g json-server \
    && mkdir -p /data

EXPOSE 3000

VOLUME /data 

ENTRYPOINT ["json-server", "-H", "0.0.0.0", "--watch", "/data/db.json"]
