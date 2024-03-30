# 切换成Linux环境
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
# 构建Linux下可运行文件
go build -o nourish_orchard_server
# 还原Windows环境
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=windows
go env -w GOOS=windows