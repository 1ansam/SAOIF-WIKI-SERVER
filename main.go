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
	router.GET("/banner", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK,
			"                 公 告                    \n"+
				"	刀剑神域关键斗士（WIKI）\n"+
				"	更新内容：\n"+
				"		1.记忆技能卡替换为中文卡面\n"+
				"		2.取消1-3星卡片展示\n"+
				"		3.增加能力、减益分类查询\n"+
				"					2022-08-04\n"+
				"	如想提交中文卡片、纠错或补充资料\n"+
				"	请联系QQ:1229921114\n"+
				"	感谢十六叶玩家提供的图片")
	})

	//单手直剑
	setRouter(router, "/saoif/sword/four_stars")
	setRouter(router, "/saoif/sword/rush")
	setRouter(router, "/saoif/sword/burst")
	setRouter(router, "/saoif/sword/connect")
	setRouter(router, "/saoif/sword/mod")

	//单手细剑
	setRouter(router, "/saoif/rapier/four_stars")
	setRouter(router, "/saoif/rapier/rush")
	setRouter(router, "/saoif/rapier/burst")
	setRouter(router, "/saoif/rapier/connect")
	setRouter(router, "/saoif/rapier/mod")

	//匕首
	setRouter(router, "/saoif/dagger/four_stars")
	setRouter(router, "/saoif/dagger/rush")
	setRouter(router, "/saoif/dagger/burst")
	setRouter(router, "/saoif/dagger/connect")
	setRouter(router, "/saoif/dagger/mod")

	//单手棍
	setRouter(router, "/saoif/club/four_stars")
	setRouter(router, "/saoif/club/rush")
	setRouter(router, "/saoif/club/burst")
	setRouter(router, "/saoif/club/connect")
	setRouter(router, "/saoif/club/mod")

	//双手斧
	setRouter(router, "/saoif/axe/four_stars")
	setRouter(router, "/saoif/axe/rush")
	setRouter(router, "/saoif/axe/burst")
	setRouter(router, "/saoif/axe/connect")
	setRouter(router, "/saoif/axe/mod")

	//双手枪
	setRouter(router, "/saoif/spear/four_stars")
	setRouter(router, "/saoif/spear/rush")
	setRouter(router, "/saoif/spear/burst")
	setRouter(router, "/saoif/spear/connect")
	setRouter(router, "/saoif/spear/mod")

	//双手弓
	setRouter(router, "/saoif/bow/four_stars")
	setRouter(router, "/saoif/bow/rush")
	setRouter(router, "/saoif/bow/burst")
	setRouter(router, "/saoif/bow/connect")
	setRouter(router, "/saoif/bow/mod")

	//盾
	setRouter(router, "/saoif/shield/four_stars")
	setRouter(router, "/saoif/shield/rush")
	setRouter(router, "/saoif/shield/burst")
	setRouter(router, "/saoif/shield/connect")
	setRouter(router, "/saoif/shield/mod")

	//能力
	setRouter(router, "/saoif/ability/four_stars")
	setRouter(router, "/saoif/ability/rush")
	//提升力量
	setRouter(router, "/saoif/ability/power_rise_3")
	setRouter(router, "/saoif/ability/power_rise_4")
	//先遣部队
	setRouter(router, "/saoif/ability/advance_trooper_5")
	setRouter(router, "/saoif/ability/advance_trooper_4")
	setRouter(router, "/saoif/ability/advance_trooper_3")
	//百战
	setRouter(router, "/saoif/ability/veteran_sword_technique_4")
	//百防
	setRouter(router, "/saoif/ability/light_magic_4")
	//防御提升
	setRouter(router, "/saoif/ability/vital_assention_4")
	setRouter(router, "/saoif/ability/vital_assention_3")
	//武器弱点攻击增加（打800/打500/刺800/刺500）
	setRouter(router, "/saoif/ability/weapon_weakness_attack_rise")

	//减益
	//刻印 小刻印 死兆印 烙印
	setRouter(router, "/saoif/debuff/mark_debuff")
	setRouter(router, "/saoif/debuff/small_mark_debuff")
	setRouter(router, "/saoif/debuff/death_mark_debuff")
	setRouter(router, "/saoif/debuff/stigma_debuff")
	//防御弱化5
	setRouter(router, "/saoif/debuff/weak_defense_debuff_5")
	//斩、次、打弱化
	setRouter(router, "/saoif/debuff/weak_slash_debuff_4")
	setRouter(router, "/saoif/debuff/weak_thrust_debuff_4")
	setRouter(router, "/saoif/debuff/weak_blunt_debuff_4")
	//暗、光、风、土、火、水弱化
	setRouter(router, "/saoif/debuff/dark_element_weak_debuff_4")
	setRouter(router, "/saoif/debuff/dark_element_weak_debuff_3")
	setRouter(router, "/saoif/debuff/holy_element_weak_debuff_4")
	setRouter(router, "/saoif/debuff/wind_element_weak_debuff_4")
	setRouter(router, "/saoif/debuff/earth_element_weak_debuff_4")
	setRouter(router, "/saoif/debuff/fire_element_weak_debuff_4")
	setRouter(router, "/saoif/debuff/fire_element_weak_debuff_5")
	setRouter(router, "/saoif/debuff/water_element_weak_debuff_4")
	setRouter(router, "/saoif/debuff/water_element_weak_debuff_3")
	//黑暗（降命中率）
	setRouter(router, "/saoif/debuff/dusk_debuff_4")
	//移除敌人攻击、防御 增加
	setRouter(router, "/saoif/debuff/remove_attack_increasing_buff")
	setRouter(router, "/saoif/debuff/remove_defense_increasing_buff")
	//切换量表增加更多
	setRouter(router, "/saoif/debuff/greater_than_usual_switch")
	//切换量表增加大量
	setRouter(router, "/saoif/debuff/greater_more_than_usual_switch")
	//中断量表增加更多
	setRouter(router, "/saoif/debuff/quicker_than_usual_break")

	err := router.Run(":7777")
	if err != nil {
		return
	}
}
