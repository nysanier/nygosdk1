gopath=$(pwd)
echo $gopath
export GOPATH=$gopath
# GOPATH="$p" go build -v  # 无效
go build -v
