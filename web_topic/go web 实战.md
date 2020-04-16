## go web 实战

### 接口设计规范

1. API有版本信息

```
/v1/getInfo
/v2/getInfo
```

2. 尽可能使用复数，且含义明确，最好是名词

```
/v1/topics
/v1/users
/v1/getusers   // 不推荐
```

3. 使用GET参数规划数据展示规则

```
/v1/users  // 显示全部或者默认条数
/v1/users?limit=10    // 只显示10条
```

### 值得注意的地方

#### 1. 路由组和{}

使用理由组可以很好的分层，使用{}来对代码进行分块，使得代码有很好阅读效果

```go
package main

import (
	"github.com/gin-gonic/gin"
	. "practice_project/web_topic/src/dao"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/topics")
	// v1的请求放在代码块里，更加容易理解和阅读
	{
		v1.GET("", GetTopicList)

		v1.GET("/:topic_id", GetTopicDetail)

		v1.Use(MustLogin())
		{
			v1.POST("", NewTopic)

			v1.DELETE("/:topic_id", DeleteTopic)
		}

	}

	router.Run(":8080")
}
```

#### 2. 数据绑定使用form

再Model里面使用数据绑定，一般情况下需要用`form`

```go
type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
```

#### 3. Model字段开头大写

```go
type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
```

如果将字段开头小写，则不会被识别

#### 4. required使得参数必须不为空

```go
type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
```

#### 5. Json映射

```go
type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
```

这里的json对应的是接收json或者传递json时候的映射



## 关于go 1.9 导包问题

### 本地导包

首先看一下我的项目结构

![](https://cdn.jsdelivr.net/gh/nizonglong/oss@master/2020-04-16%2015:28:56-uPic-Snipaste_2020-04-16_15-28-39.png)

1. 首先，查看自己代码所在位置`pwd`

```
我的位置是这个地方
~/workspace/goland_projects/src/practice_project
```

然后，确定自己的GOPATH是`~/workspace/goland_projects`，其中的src是默认的因此不用添加，否则路径报错

2. 配置自定义GOPATH

`export GOPATH=~/workspace/goland_projects`

3. 安装包

`go install practice_project/web_topic`

如图，安装好以后就可以用了，然后再导入包就可以

![](https://cdn.jsdelivr.net/gh/nizonglong/oss@master/2020-04-16%2015:29:42-uPic-Snipaste_2020-04-16_15-29-36.png)

导入->可以使用了

![](https://cdn.jsdelivr.net/gh/nizonglong/oss@master/2020-04-16%2015:33:18-uPic-Snipaste_2020-04-16_15-30-27.png)

