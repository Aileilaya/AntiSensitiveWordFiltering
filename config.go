package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type _Config struct {
	Version    string
	Separetor  string
	Dictionary []string
}

var Config _Config
var iversion = "20210929"
var preset = []string{
	"裸体",
	"发情",
	"婊子",
	"娼妇",
	"调教",
	"快感",
	"赤裸",
	"拔出",
	"性骚扰",
	"猥亵",
	"插入",
	"半裸",
	"贱奴",
	"全裸",
	"淫乱",
	"舔舐",
	"自慰",
	"奸淫",
	"淫母",
	"吮吸",
	"贫乳",
	"巨乳",
	"情欲",
	"淫靡",
	"性爱",
	"奸杀",
	"强奸",
	"裸露",
	"性欲",
	"黄段子",
	"色情",
	"插进",
	"幼女",
	"乳首",
	"情趣",
	"猥琐",
	"性奴",
	"射出",
	"凌辱",
	"下流",
	"妓女",
	"处女",
	"插的",
	"激射",
	"射进",
	"贱人",
	"自残",
	"榨干",
	"爱爱",
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

var _config_name = "./config.json"

func readConfig() bool {
	if !CheckFileExist(_config_name) {
		return false
	}
	fd, _ := os.Open(_config_name)
	defer fd.Close()
	d, _ := ioutil.ReadAll(fd)
	err := json.Unmarshal(d, &Config)
	if err != nil {
		return false
	}
	return true
}

func SaveConfig() {
	removeDuplicateItem()
	save, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		fmt.Println("[SaveConfig] " + err.Error())
	}
	fd, _ := os.Create(_config_name)
	defer fd.Close()
	fd.Write(save)
}

func setDefaultConfig() {
	Config.Version = iversion
	Config.Dictionary = append(Config.Dictionary, preset...)
	Config.Separetor = "^"
	SaveConfig()
}

func InitConfig() {
	if !readConfig() {
		fmt.Println("[!] 配置文件无效")
		setDefaultConfig()
		os.Exit(1)
	} else {
		if Config.Version != iversion {
			fmt.Println("[!] 配置文件需要更新")
			Config.Version = iversion
			setDefaultConfig()
			os.Exit(1)
		}
	}
}

func removeDuplicateItem() {
	var hashset = make(map[string]bool)
	for i := range Config.Dictionary {
		hashset[Config.Dictionary[i]] = true
	}
	Config.Dictionary = []string{}
	for k, _ := range hashset {
		Config.Dictionary = append(Config.Dictionary, k)
	}
}
