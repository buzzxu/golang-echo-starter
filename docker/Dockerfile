FROM buzzxu/docker-ci:golang as build

WORKDIR $GOPATH/src/golang-echo-starter
ADD . $GOPATH/src/golang-echo-starter

# 设置 GitHub 令牌
ARG GITHUB_TOKEN
RUN echo "machine github.com login YOUR_GITHUB_USERNAME password ${GITHUB_TOKEN}" > ~/.netrc

# 下载私有模块
ENV GOPRIVATE github.com/buzzxu
RUN go get -d github.com/buzzxu/yuanmai-platform-go

RUN cd && \
    cd $GOPATH/src/formflow-admin-go/cmd && \
    go build -o app; \
    mv app  /opt/app;

FROM buzzxu/docker:debian

LABEL org.opencontainers.image.source="https://github.com/buzzxu/golang-echo-starter"
LABEL org.opencontainers.image.authors="downloadxu@163.com"
LABEL org.opencontainers.image.description="这是一个基于 Golang 的表单管理平台"

WORKDIR /app
COPY --from=build /opt/app /app

ADD docker/env.toml /app/env.toml


EXPOSE 3000
CMD ["./app"]