package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type image struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getFiles(path string) (images []image, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	pthSep := string(os.PathSeparator)

	for id, fi := range dir {
		file, err := os.Open(path + "/" + fi.Name())
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
		image := image{Id: id, Name: fi.Name(), Url: path + pthSep + fi.Name()}
		images = append(images, image)
	}

	return images, nil
}

func setRouter(router *gin.Engine, path string) {
	router.GET(path, func(context *gin.Context) {
		images, err := getFiles("." + path)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound, nil)
		} else {
			context.IndentedJSON(http.StatusOK, images)
		}

	})
	router.GET(path+"/:key/", func(context *gin.Context) {
		key := context.Param("key")
		file, err := ioutil.ReadFile("." + path + "/" + key)
		if err != nil {
			return
		}

		if err != nil {
			return
		}
		fileContentDisposition := "attachment;filename=\"" + key + "\""
		context.Header("Content-Disposition", fileContentDisposition)
		context.Data(http.StatusOK, "application/x-png", file)
	})
}

func main() {
	router := gin.Default()
	//能力
	setRouter(router, "/saoif/ability/four_stars")
	setRouter(router, "/saoif/ability/three_stars")
	setRouter(router, "/saoif/ability/two_stars")
	setRouter(router, "/saoif/ability/one_star")
	setRouter(router, "/saoif/ability/rush")

	//单手直剑
	setRouter(router, "/saoif/sword/four_stars")
	setRouter(router, "/saoif/sword/three_stars")
	setRouter(router, "/saoif/sword/two_stars")
	setRouter(router, "/saoif/sword/one_star")
	setRouter(router, "/saoif/sword/rush")
	setRouter(router, "/saoif/sword/burst")
	setRouter(router, "/saoif/sword/connect")
	setRouter(router, "/saoif/sword/mod")

	//单手细剑
	setRouter(router, "/saoif/rapier/four_stars")
	setRouter(router, "/saoif/rapier/three_stars")
	setRouter(router, "/saoif/rapier/two_stars")
	setRouter(router, "/saoif/rapier/one_star")
	setRouter(router, "/saoif/rapier/rush")
	setRouter(router, "/saoif/rapier/burst")
	setRouter(router, "/saoif/rapier/connect")
	setRouter(router, "/saoif/rapier/mod")

	//匕首
	setRouter(router, "/saoif/dagger/four_stars")
	setRouter(router, "/saoif/dagger/three_stars")
	setRouter(router, "/saoif/dagger/two_stars")
	setRouter(router, "/saoif/dagger/one_star")
	setRouter(router, "/saoif/dagger/rush")
	setRouter(router, "/saoif/dagger/burst")
	setRouter(router, "/saoif/dagger/connect")
	setRouter(router, "/saoif/dagger/mod")

	//单手棍
	setRouter(router, "/saoif/club/four_stars")
	setRouter(router, "/saoif/club/three_stars")
	setRouter(router, "/saoif/club/two_stars")
	setRouter(router, "/saoif/club/one_star")
	setRouter(router, "/saoif/club/rush")
	setRouter(router, "/saoif/club/burst")
	setRouter(router, "/saoif/club/connect")
	setRouter(router, "/saoif/club/mod")

	//双手斧
	setRouter(router, "/saoif/axe/four_stars")
	setRouter(router, "/saoif/axe/three_stars")
	setRouter(router, "/saoif/axe/two_stars")
	setRouter(router, "/saoif/axe/one_star")
	setRouter(router, "/saoif/axe/rush")
	setRouter(router, "/saoif/axe/burst")
	setRouter(router, "/saoif/axe/connect")
	setRouter(router, "/saoif/axe/mod")

	//双手枪
	setRouter(router, "/saoif/spear/four_stars")
	setRouter(router, "/saoif/spear/three_stars")
	setRouter(router, "/saoif/spear/two_stars")
	setRouter(router, "/saoif/spear/one_star")
	setRouter(router, "/saoif/spear/rush")
	setRouter(router, "/saoif/spear/burst")
	setRouter(router, "/saoif/spear/connect")
	setRouter(router, "/saoif/spear/mod")

	//双手弓
	setRouter(router, "/saoif/bow/four_stars")
	setRouter(router, "/saoif/bow/three_stars")
	setRouter(router, "/saoif/bow/two_stars")
	setRouter(router, "/saoif/bow/one_star")
	setRouter(router, "/saoif/bow/rush")
	setRouter(router, "/saoif/bow/burst")
	setRouter(router, "/saoif/bow/connect")
	setRouter(router, "/saoif/bow/mod")

	//盾
	setRouter(router, "/saoif/shield/four_stars")
	setRouter(router, "/saoif/shield/three_stars")
	setRouter(router, "/saoif/shield/two_stars")
	setRouter(router, "/saoif/shield/one_star")
	setRouter(router, "/saoif/shield/rush")
	setRouter(router, "/saoif/shield/burst")
	setRouter(router, "/saoif/shield/connect")
	setRouter(router, "/saoif/shield/mod")

	err := router.Run(":7777")
	if err != nil {
		return
	}
}
