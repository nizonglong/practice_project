## gin study

### gin 使用

在需要使用的项目里加上

`import "github.com/gin-gonic/gin"`

即可使用gin框架



## gin router【路由】

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/get", func(c *gin.Context) {
        c.String(200, "get")
    })
    r.POST("/post", func(c *gin.Context) {
        c.String(200, "post")
    })
    r.Handle("DELETE", "/delete",
        func(c *gin.Context) {
            c.String(200, "delete")
        })
    r.Any("/any", func(c *gin.Context) {
        c.String(200, "any")
    })
    r.Run() // listen and serve on 127.0.0.1:8080
}
```

1. router请求方式

GET, POST, DELETE, PUT, PATCH, HEAD, OPTIONS, CONNECT, TRACE

2. 若要一个接口响应不同的请求方式，可以使用any，支持所有请求方式

#### router static 路由请求静态资源

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.Static("/assets","./assets")
    r.StaticFS("/static",http.Dir("static"))
    r.StaticFile("/favicon.icon","./favicon.icon")

    r.Run() // listen and serve on 127.0.0.1:8080
}
```

使用cmd： curl "http://127.0.0.1:8080/assets/b.html"

访问静态资源



#### router url 类似于restful

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/:name/:id", func(c *gin.Context) {
        c.JSON(200,gin.H{
            "name":c.Param("name"),
            "id":c.Param("id"),
        })
    })

    r.POST("/:name/:id", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "name_p":c.Param("name"),
            "id_p":c.Param("id"),
        })
    })
    r.Run() // listen and serve on 127.0.0.1:8080
}
```

关于post的用法与上面类似，换一个函数即可

#### router_generic 泛绑定

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/user/*action", func(c *gin.Context) {
        c.String(200, "hello world")
    })

    r.Run() // listen and serve on 127.0.0.1:8080
}
```

泛绑定：意思个通配符差不多，类似于/user/**

### param get 参数传递get方式

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/test", func(c *gin.Context) {
        firstName := c.Query("first_name")
        lastName := c.DefaultQuery("last_name", "last_default_name")

        c.String(http.StatusOK, "%s,%s", firstName, lastName)
    })
    r.Run(":8080") // listen and serve on 127.0.0.1:8080
}
```

c.Query("first_name")：默认获取传递的参数，没有默认值

c.DefaultQuery("last_name", "last_default_name")：获取传递的参数，若为空则传递默认值last_default_name

**POST方式和get类似**

#### param body获取body的参数

```go
package main

import (
    "bytes"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
)

func main() {
    r := gin.Default()
    r.POST("/test", func(c *gin.Context) {
        bodyByts, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.String(http.StatusBadRequest, err.Error())
            c.Abort()
        }

        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))
        firstName := c.Query("first_name")
        lastName := c.DefaultQuery("last_name", "last_default_name")
        c.String(http.StatusOK, string(bodyByts))
        c.String(http.StatusOK, "%s,%s", firstName, lastName)
    })
    r.Run() // listen and serve on 127.0.0.1:8080
}
```

注意：`bodyByts, err := ioutil.ReadAll(c.Request.Body)`这句话是使用ioutil将c.Request.Body的内容导入bodyByts，因此c.Request.Body里面的数据就『剪切』走了。

如果要重新让c.Request.Body有数据，就要用c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))重新赋值回去

#### param struct

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

type Person struct {
    Name     string    `form:"name"`
    Address  string    `form:"address"`
    Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
    r := gin.Default()
    r.GET("/testing", testing)
    r.POST("/testing", testing)
    r.Run(":8080") // listen and serve on 127.0.0.1:8080
}

func testing(c *gin.Context) {
    var person Person
    // 这里是根据请求的content type来做不同binding操作
    if err := c.ShouldBind(&person); err == nil {
        c.String(http.StatusOK, "%v", person)
    } else {
        c.String(http.StatusInternalServerError, "person bind error : %v", err)
    }
}
```



#### gin bind参数绑定

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

type Person struct {
    Name     string    `form:"name"`
    Address  string    `form:"address"`
    Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
    r := gin.Default()
    r.GET("/testing", testing)
    r.POST("/testing", testing)
    r.Run(":8080") // listen and serve on 127.0.0.1:8080
}

func testing(c *gin.Context) {
    var person Person
    // 这里是根据请求的content type来做不同binding操作
    if err := c.ShouldBind(&person); err == nil {
        c.String(http.StatusOK, "%v", person)
    } else {
        c.String(http.StatusInternalServerError, "person bind error : %v", err)
    }
}
```

#### valid binding 绑定校验

```
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Person struct {
    Age     int    `form:"age" binding:"required,gt=10"`
    Name    string `form:"name" binding:"required"`
    Address string `form:"address" binding:"required"`
}

func main() {
    r := gin.Default()
    r.GET("/testing", func(c *gin.Context) {
        var person Person
        if err := c.ShouldBind(&person); err != nil {
            c.String(http.StatusInternalServerError, "%v", err)
            c.Abort()
            return
        }
        c.String(http.StatusOK, "%v", person)
    })
    r.Run()
}
```

go doc文档地址:https://godoc.org/gopkg.in/go-playground/validator.v9

校验绑定是对数据格式的校验，确保数据格式正确

#### vaild custom自定义校验

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "gopkg.in/go-playground/validator.v8"
    "net/http"
    "reflect"
    "time"
)

type Booking struct {
    CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
    CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// Structure
func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

    if date, ok := field.Interface().(time.Time); ok {
        today := time.Now()
        if date.Unix() > today.Unix() {
            return true
        }
    }

    return false
}

func main() {
    r := gin.Default()

    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        _ = v.RegisterValidation("bookabledate", bookableDate)
    }

    r.GET("/bookable", func(c *gin.Context) {
        var b Booking
        if err := c.ShouldBind(&b); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "ok!", "booking": b})
    })
    r.Run()
}
```

这里可以自定义校验方式，使用go doc内提供的统一校验function样式去定义自己的校验即可，详情可在代码中查看

#### valid 多语言化

```go
package main

import (
    "github.com/gin-gonic/gin"
    en2 "github.com/go-playground/locales/en"
    zh2 "github.com/go-playground/locales/zh"
    "github.com/go-playground/universal-translator"
    "gopkg.in/go-playground/validator.v9"
    en_translations "gopkg.in/go-playground/validator.v9/translations/en"
    zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
    "net/http"
)

type Person struct {
    Age     int    `form:"age" validate:"required,gt=10"`
    Name    string `form:"name" validate:"required"`
    Address string `form:"address" validate:"required"`
}

var (
    Uni      *ut.UniversalTranslator
    Validate *validator.Validate
)

// 验证信息多语言化
func main() {
    Validate = validator.New()

    zh := zh2.New()
    en := en2.New()
    Uni := ut.New(zh, en)

    r := gin.Default()
    r.GET("/testing", func(c *gin.Context) {
        locale := c.DefaultQuery("locale", "zh")
        trans, _ := Uni.GetTranslator(locale)
        switch locale {
        case "zh":
            zh_translations.RegisterDefaultTranslations(Validate, trans)
        case "en":
            en_translations.RegisterDefaultTranslations(Validate, trans)
        default:
            zh_translations.RegisterDefaultTranslations(Validate, trans)
        }

        var person Person
        if err := c.ShouldBind(&person); err != nil {
            c.String(http.StatusInternalServerError, "%v", err)
            c.Abort()
            return
        }

        if err := Validate.Struct(person); err != nil {
            errs := err.(validator.ValidationErrors)
            sliceErrs := []string{}
            for _, e := range errs {
                sliceErrs = append(sliceErrs, e.Translate(trans))
            }

            c.String(http.StatusInternalServerError, "%v", sliceErrs)
            c.Abort()
            return
        }

        c.String(http.StatusOK, "%v", person)

        r.Run()
    })

    r.Run()
}
```

通过调整`locale := c.DefaultQuery("locale", "zh")`里面的zh可以使用不同的语言

### Gin 中间件

r:=gin.Default() 默认自带Logger和Recovery两个中间件

拆分查看使用

```go
package main

import (
    "github.com/gin-gonic/gin"
    "io"
    "net/http"
    "os"
)

func main() {
    f,_:=os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)
    gin.DefaultErrorWriter = io.MultiWriter(f)

    r := gin.New()

    r.Use(gin.Logger(), gin.Recovery())
    r.GET("/test", func(c *gin.Context) {
        name := c.DefaultQuery("name", "default_name")
        panic("test panic")
        c.String(http.StatusOK, "%s", name)
    })
    r.Run()
}
```

Logger如果不定义

```go
f,_:=os.Create("gin.log")
gin.DefaultWriter = io.MultiWriter(f)
gin.DefaultErrorWriter = io.MultiWriter(f)
```

则默认将日志打印在控制台输出，这里指定了写入到log文件中

#### 自定义白名单中间件

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func IPAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        ipList := []string{
            "127.0.0.2",
        }

        flag := false
        clientIP := c.ClientIP()
        for _, host := range ipList {
            if clientIP == host {
                flag = true
                break
            }
        }

        if !flag {
            c.String(http.StatusUnauthorized, "%s, not in iplist", clientIP)
        }
    }
}

func main() {
    r := gin.Default()
    r.Use(IPAuthMiddleware())
    r.GET("/test", func(c *gin.Context) {
        c.String(http.StatusOK, "hello test")
    })
    r.Run()
}
```

### 拓展：超时关闭服务器

正常情况下，如果要关闭服务器，立刻关闭的话会使得一部分数据丢失导致无法正常返回数据。那么超时关闭就是处理完最后的一批数据后再关闭应用。案例如下

```go
package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    r := gin.Default()
    r.GET("/test", func(c *gin.Context) {
        time.Sleep(10 * time.Second)
        c.String(http.StatusOK, "hello test")
    })

    srv := &http.Server{
        Addr:    ":8085",
        Handler: r,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("shutdown server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("server shutdown: ", err)
    }

    log.Println(":server exiting")
}
```

如果访问test会有10s超时，在这期间关闭服务器不会立刻关闭会有10s超时，在关闭的时候可以正常处理test的请求并返回然后关闭服务器。

### 拓展：模板渲染

通俗来说就后端传参给html界面，是属于前后端不分离的，不建议使用

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("template/*")
    r.GET("/index", func(c *gin.Context) {
        c.HTML(http.StatusOK,"index.html",gin.H{
            "title":"index.html",
        })
    })
    r.Run()
}
```

cmd： curl -X GET "http://127.0.0.1:8080/index"

可以看到相应的界面被传递了参数，值是"index.html"

### 拓展：自动化证书配置

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/autotls"
    "net/http"
)

func main() {
    r:=gin.Default()
    r.GET("/test", func(c *gin.Context) {
        c.String(http.StatusOK,"hello test")
    })

    autotls.Run(r, "www.itpp.tk")
}
```

