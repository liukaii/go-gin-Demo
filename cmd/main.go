package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/pkg"
	"go-gin/mdw"
)


func main() {
	//创建一个路由handler
	router := gin.Default()

	v1 := router.Group("/api/v1")
	//绑定路由规则和路由函数
	{
		v1.GET("/", pkg.Demo1)
		//Param方法获取name
		v1.GET("/user/:name", pkg.DemoWithParam1)
		//利用*匹配的规则更多,但是实际测试未能成功,后发现此处不能与user相同，故改为user1
		v1.GET("/user1/:name/*action", pkg.DemoWithParam2)
		//Query方法获得参数，注意请求的特殊符号前要用\转义形式为\?firstname\=liu\&lastname\=kai,可以出现中文,urlencode
		v1.GET("/welcome", pkg.DemoWithQuery1)
		//PostForm获得参数，请求方法为curl -X POST http://219.245.186.205:8080/api/v1/form_post
		// -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=liukai" | python -m json.tool
		v1.POST("/form_post", pkg.DemoWithPostForm1)
		//同时使用query string 和 body 发送参数给服务器
		v1.PUT("/post", pkg.DemoWithQueryPostForm1)
		//上传单个文件,上传方式为curl -X POST http://219.245.186.205:8080/api/v1/upload
		// -F "upload=@/home/lk/Desktop/1.jpg" -H "Content-Type:multipart/form-data"
		v1.POST("/upload", pkg.DemoWithUpLoad1)
		//上传多个文件，利用多个-F "upload=@/home/lk/Desktop/1.jpg"形式
		v1.POST("/multi/upload", pkg.DemoWithMultiUpload1)
		//参数绑定,目前请求application/content-type为x-www-form-urlencoded且请求为key=value形式时成功了
		//content-type为application/json且请求为'{"key":"value"}'形式(注意json是有数据类型)成功了，注意最外面是单引号而不是`
		v1.POST("/login", pkg.DemoParamBinding)
		//用bind绑定参数，自动推断是bind表单还是json的参数
		v1.POST("loginwithbind", pkg.DemoWithBind)
		//多格式渲染c.XML,响应也可以使用不同的content-type,通常有html,text,json,xml,plain等
		v1.GET("/render",pkg.DemoWithXML)
		//重定向的请求
		v1.GET("/redict/baidu", pkg.DemoWithRedict)
		//分组路由，让代码逻辑更加模块化，同时分组也易于定义中间件的使用范围（可以了解一下Flask）
		v1.GET("/auth/signin", pkg.DemoWithCookie)
		//异步协程
		v1.GET("/sync", pkg.DemoWithSync)
		//同步的逻辑可以看到，服务的进程睡眠了。异步的逻辑则看到响应返回了。然而我暂时并没有搞懂
		v1.GET("/async", pkg.DemoWithAsync)
	}
	//群组的中间件,AuthMiddleware是一个简易的中间件，执行AuthMiddleware的逻辑，再到/home
	authorized := router.Group("/", mdw.AuthMiddleware())
	{
		authorized.POST("/home", pkg.LoginEndpoint)
	}

	//单个路由中间件,对指定路由函数进行注册，写在router.Use(mdw.Middleware())之前
	router.GET("/before", mdw.Middleware(), pkg.DemoWithSingleMiddleware)
	//尽管是全局中间件，注册中间件过程之前设置的路由，将不会受注册的中间件所影响，只有注册了中间件以下代码的路由函数规则，才会被中间件装饰
	router.Use(mdw.Middleware())
	//该花括号只是一种代码规范，只要使用router进行路由的路由函数都等于被装饰了，区分权限，可以使用组返回的对象注册中间件
	{
		//使用router装饰中间件，然后在/middleware即可读取request的值
		router.GET("/middleware", pkg.DemoWithMiddleware)
	}

	//启动路由，监听端口
	router.Run(":8080")
	//另一种开启方式
	//http.ListenAndServe(":8080", router)
	//自定义router
	/*
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	*/
}
