## gin安装遇到问题

### 版本问题

问题：直接git clone 最新的gin框架使用，发生报错问题

cannot find package github.com/go-playground/validator/v10

原因：由于目前安装环境是go1.9，但是最新clone的gin是1.6.2，版本不兼容

解决：clone gin的1.4版本即可，相关搜索链接如下

https://github.com/go-playground/validator/issues/546



### cannot find package "golang.org/x/sys/unix" in any of:

原因：由于限制问题，国内使用 go get 安装 golang 官方包可能会失败

解决：

```
cd ~/go/src
mkdir -p golang.org/x
cd golang.org/x
git clone https://github.com/golang/sys.git
```

