# 若运行环境是 linux 则，需要换成 debian
# 使用最新版 alpine 作为基础镜像，alpine 是一个超轻量级 Linux 发行版本，专为 Docker 中 Web 应用而生。
# 它能保证绝大多数 web 应用可以正常运行， 即使它只包含必要的 run-time 文件和依赖，镜像大小只有 4 MB
FROM debian:latest
# 在容器的根目录下创建 app 目录
RUN mkdir /app
# 将工作目录切换到 /app 下
WORKDIR /app
# 将微服务的运行文件拷贝到 /app 目录下
ADD consignment-service /app/consignment-service

# 运行服务端程序
CMD ["./consignment-service"]
