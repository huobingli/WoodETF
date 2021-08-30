package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const BaseUploadPath = "C:\\ci_auto_publish\\"

//启动
func main() {
	router := gin.Default()
	router.POST("/upload_file", HandleUploadFile)
	router.POST("/upload_muti_file", HandleUploadMutiFile)
	router.GET("/download", HandleDownloadFile)
	router.GET("/getUrl", Get)
	router.GET("/file", HandleShowFile)
	//router.GET("/", HandleShowFile)
	router.Run(":7001")
}

//错误
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(c *gin.Context) {
	url := c.Param("url")
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	c.JSON(200, gin.H{"status": 1, "result": result.String(), "message": "Success"})
}

// HandleUploadFile 上传单个文件
func HandleUploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}
	fmt.Println(header.Filename)
	fmt.Println(string(content))
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}

// HandleUploadMutiFile 上传多个文件
func HandleUploadMutiFile(c *gin.Context) {
	// 限制上传文件大小
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 4<<20)
	// 限制放入内存的文件大小
	err := c.Request.ParseMultipartForm(4 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}
	formdata := c.Request.MultipartForm
	files := formdata.File["file"]
	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		fmt.Println(v.Filename)
		fmt.Println(string(content))
	}
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}

// HandleDownloadFile 下载文件
func HandleDownloadFile(c *gin.Context) {
	//content := c.Query("content")
	//content = "hello world, 我是一个文件，" + content

	filename := c.Query("content")
	path := "D:\\ci\\autoBuild\\cd-tool\\output\\build\\"
	filepath := path + filename
	c.Writer.WriteHeader(http.StatusOK)
	// 返回文件名
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/text/plain")
	c.File(filepath)
}

func HandleShowFile(c *gin.Context) {
	path := "D:\\ci\\autoBuild\\cd-tool\\output\\build\\"
	fileName := path + c.Query("name")
	fmt.Println(fileName)
	c.File(fileName)

	//http.FileServer(http.Dir(BaseUploadPath))
	//http.FileServer()

}
