## gin安装遇到问题

### 版本问题

问题：直接git clone 最新的gin框架使用，发生报错问题

cannot find package github.com/go-playground/validator/v10

原因：由于目前安装环境是go1.9，但是最新clone的gin是1.6.2，版本不兼容

解决：clone gin的1.4版本即可，相关搜索链接如下

https://github.com/go-playground/validator/issues/546



### 