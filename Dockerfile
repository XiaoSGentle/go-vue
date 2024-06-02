
# 构建后端
FROM golang:1.22.0-alpine AS GOBUILD

WORKDIR /app

COPY ./go-admin/xadmin ./xadmin/
COPY ./go-admin/xcore ./xcore/


ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

WORKDIR /app/xcore
RUN go mod tidy

WORKDIR /app/xadmin
RUN go mod tidy
RUN go work sync

RUN go build -o BACK_API ./main.go
WORKDIR /app


# 构建前端

FROM node:18 AS FRONTBUILD

WORKDIR /app

COPY ./soybean-admin .

RUN npm install -g pnpm
RUN pnpm config set store-dir /root/.local/share/pnpm/store/v3 --global
RUN pnpm config set registry http://mirrors.cloud.tencent.com/npm/
# Remove the node_modules directory
RUN rm -rf node_modules
RUN pnpm install
RUN pnpm run build

# 运行环境
FROM alpine:latest

WORKDIR /apps

COPY --from=GOBUILD /app/xadmin/BACK_API .

# 拷贝静态文件
COPY --from=GOBUILD /app/xadmin/public ./public

# 拷贝配置文件
COPY --from=GOBUILD /app/xadmin/config ./config

# 拷贝前端文件
COPY --from=FRONTBUILD /app/dist ./public/front/

CMD ["./BACK_API"]