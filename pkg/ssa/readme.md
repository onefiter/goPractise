# 查看ssa
## 1. 设置环境变量
```shell
# windows
$env:GOSSAFUNC="main"

#linux
export GOSSAFUNC=main

go build main.go #会生成对应的xxx.html，用浏览器打开对应的html文件，查看对应的ssa
```
# 查看Plan9汇编代码

```shell
go build -gcflags -S main.go
```
