from alpine
run sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk --no-cache add tzdata
add app /app/
add etc/docker-config.yaml /app/etc/config.yaml
workdir /app
ENV REDIS_TYPE="node" \
    REDIS_PASS="" \
    REDIS_TLS=false \
    WebSiteURL="http://127.0.0.1:8888" \
    ShortKeyLength=7 \
    ShortKeyTTL=604800
ENTRYPOINT ["./app"]