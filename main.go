package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type image struct {
	Id     int    `json:"id""`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Base64 string `json:"base64"`
}

func getFiles(path string) (images []image, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	pthSep := string(os.PathSeparator)

	for id, fi := range dir {
		file, err := os.Open("./saoif/" + fi.Name())
		if err != nil {
			return nil, err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		stat, err := file.Stat()
		if err != nil {
			return nil, err
		}
		data := make([]byte, stat.Size())
		_, err = file.Read(data)
		if err != nil {
			return nil, err
		}
		image := image{Id: id, Name: fi.Name(), Url: path + pthSep + fi.Name(), Base64: base64.StdEncoding.EncodeToString(data)}
		images = append(images, image)
	}

	return images, nil
}

func main() {
	router := gin.Default()
	//
	router.GET("/saoif", func(context *gin.Context) {
		images, err := getFiles("./saoif")
		if err != nil {
			context.IndentedJSON(http.StatusNotFound, nil)
		} else {
			context.IndentedJSON(http.StatusOK, images)
		}

	})
	router.GET("/saoif/:key", func(context *gin.Context) {
		key := context.Param("key")
		file, err := ioutil.ReadFile("./saoif/" + key)
		if err != nil {
			return
		}

		if err != nil {
			return
		}
		fileContentDisposition := "attachment;filename=\"" + key + "\""
		context.Header("Content-Type", "application/x-png")
		context.Header("Content-Disposition", fileContentDisposition)
		context.Data(http.StatusOK, "application/x-png", file)
	})

	router.Run("localhost:8990")
}
