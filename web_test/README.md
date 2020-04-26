exampleapi
========

http web 后台项目规范示例

## docs

包含.sql文件和.toml文件

1. .sql文件内容为项目用到的数据表结构的sql语句
2. .toml文件内容为项目用到的配置，包括数据库配置、redis配置、请求的远程接口地址等

## pkg

项目主目录

### cache
缓存相关操作

### config
全局配置

### handlers
封装好的请求处理操作

1. 以handler_开头的文件表示具体的请求入口

2. request_param文件包含所有请求参数的结构体

3. constants文件包含handlers层用到的所有常量或全局变量，以中间件名称、配置参数等为主

5. utils包含handlers层用到的一些公共方法

6. 其他不以handler开头的文件则是一些响应参数的结构体或该层用到的一些公共结构体

### models
数据库相关操作

### remote_api
封装好的远程接口调用

### router
初始化路由、配置handler拦截规则、中间件注入等操作

## test
单元测试相关代码

1. handler_开头的文件表明对应于具体handler的单元测试

2. init文件包含单元测试路由环境的初始化操作

3. response文件包含各请求对应的响应参数结构

## main
程序入口

## .gitignore
忽略的文件配置

## README.md
项目介绍文档