# 编译程序和镜像
CGO_ENABLED=0 go build # 禁止 CGO 以获得静态链接 binary ，方便使用 FROM scratch，缩小大小
docker build -t my_version:latest .
# 传输镜像
docker save my_version:latest | gzip > image.tar.gz
cat image.tar.gz | ssh ice-lab5 docker load
cat image.tar.gz | ssh ice-p1620 docker load
