FROM alpine:latest

LABEL maintainer="zhuyasen zhuyasen@126.com"

# 修改时区，需要翻墙
RUN apk update \
   && apk add tzdata \
   && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
   && echo "Asia/Shanghai" > /etc/timezone

# 复制app二进制文件
COPY ./eiblog /eiblog/eiblog
RUN chmod +x /eiblog/eiblog
# 复制静态文件
COPY ./conf /eiblog/conf
COPY ./static /eiblog/static
COPY ./views /eiblog/views

WORKDIR /eiblog/
# 运行app
CMD ./eiblog
