package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"github.com/cloudflare/cfssl/log"
	"io"
	"github.com/gin-gonic/gin/binding"
	"time"
)

func Demo1(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func DemoWithParam1(c *gin.Context) {
	//拿到name参数
	name := c.Param("name")
	//返回信息
	c.String(http.StatusOK, "Hello %s", name)
}

func DemoWithParam2(c *gin.Context) {
	//拿到name参数
	name := c.Param("name")
	//拿到action参数
	action := c.Param("action")
	//返回信息
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func DemoWithQuery1(c *gin.Context) {
	//请求形式 ？key1=value1&key2=value2
	//得到key1, 默认值Guset
	firstname := c.DefaultQuery("firstname","Guest")
	//得到key2
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func DemoWithPostForm1(c *gin.Context) {
	//拿到message对应的信息
	message := c.PostForm("message")
	//拿到nick的信息，默认值为anonymous
	nick := c.DefaultPostForm("nick", "anonymous")
	//返回json信息
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"status_code": http.StatusOK,
			"status": "ok",
		},
		"message": message,
		"nick": nick,
	})
}

func DemoWithQueryPostForm1(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")
	fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
	})
}

func DemoWithUpLoad1(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println("name",name)
	//解析客户端文件name属性，如果不传文件则会报错
	file, header, err := c.Request.FormFile("upload")
	//fmt.Println("file", file)
	//fmt.Println("header", header)
	if err != nil {
		//http.StatusBadRequest = 400
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//filename图片名
	filename := header.Filename
	//fmt.Println("filename", filename)

	//fmt.Println(file, err, filename)
	//out输出路径
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	//fmt.Println("out", out)
	//从file复制到out
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	//http.StatusCreated = 201
	c.String(http.StatusCreated, "upload successful")
}

func DemoWithMultiUpload1(c *gin.Context) {
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		log.Fatal(err)
	}
	//得到文件句柄
	formdata := c.Request.MultipartForm
	//得到文件列表
	files := formdata.File["upload"]
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(files[i].Filename)

		defer file.Close()

		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(out, file)

		if err != nil {
			log.Fatal(err)
		}
		//http.StatusCreated = 201
		c.String(http.StatusCreated, "upload successful")

	}
}
func DemoWithWebForm1(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}

func DemoParamBinding(c *gin.Context) {
	var user User
	var err error
	contentType := c.Request.Header.Get("Content-Type")

	switch contentType {
	//新的web表单页的content-type类型
	case "application/json":
		err = c.BindJSON(&user)
	//旧的web表单页的content-type类型
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&user, binding.Form)
	}

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.Username,
		"passwd": user.Passwd,
		"age": user.Age,
	})
}

func DemoWithBind(c *gin.Context) {
	var user User
	//用Bind绑定参数
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"passwd": user.Passwd,
		"age": user.Age,
	})
}

func DemoWithXML(c *gin.Context) {
	//拿到contentType
	contentType := c.DefaultQuery("content_type", "json")
	if contentType == "json" {
		c.JSON(http.StatusOK, gin.H{
			"user": "liukai",
			"passwd": "123",
		})
	} else if contentType == "xml" {
		c.XML(http.StatusOK, gin.H{
			"user": "liukai",
			"passwd": "123",
		})
	}
}

func DemoWithRedict(c *gin.Context) {
	//重定向，http.StatusMovedPermanently = 301
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}

func DemoWithMiddleware(c *gin.Context) {
	//如果没有注册中间件就使用MustGet方法读取c的值可能会抛错，可以使用Get方法取而代之
	request := c.MustGet("request").(string)
	req, _ := c.Get("request")

	c.JSON(http.StatusOK, gin.H{
		"middle_request": request,
		"request": req,
	})
}

func DemoWithSingleMiddleware(c *gin.Context) {
	request := c.MustGet("request").(string)

	c.JSON(http.StatusOK, gin.H{
		"middle_request": request,
	})
}

//设置cookie
func DemoWithCookie(c *gin.Context) {
	//cookie名为session_id。需要指定path为/，不然gin会自动设置cookie的path为/auth
	cookie := &http.Cookie{
		Name: "session_id",
		Value: "123",
		Path: "/",
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)
	c.String(http.StatusOK, "Login successgul")
}

func LoginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "home",
	})
}

func DemoWithSync(c *gin.Context) {
	time.Sleep(5 * time.Second)
	fmt.Println("Done! in path: %s" + c.Request.URL.Path)
}

func DemoWithAsync(c *gin.Context) {
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path: %s" + cCp.Request.URL.Path)
	}()
}