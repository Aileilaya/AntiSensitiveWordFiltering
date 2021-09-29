package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var rpath string
var spath string
var mode int

func init() {
	rand.Seed(time.Now().UnixNano())
	InitConfig()
	flag.StringVar(&rpath, "i", "", "目标路径")
	flag.StringVar(&spath, "o", "", "保存路径")
	flag.IntVar(&mode, "m", 0, "0 单文件 1 文件夹")
	flag.Parse()
	if !CheckFileExist(spath) {
		os.MkdirAll(spath, 0600)
	}
}

func main() {
	switch mode {
	case 0:
		processFile(rpath)
	case 1:
		filepath.Walk(rpath, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				if strings.HasSuffix(path, ".txt") {
					processFile(path)
				}
			}
			return nil
		})
	}

}

func processFile(p string) {
	f, err := os.OpenFile(p, os.O_RDWR, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	shaTang(f)
}

func shaTang(file *os.File) {
	fmt.Println("正在处理 " + file.Name())
	content, _ := ioutil.ReadAll(file)
	scontent := string(content)
	for i := range Config.Dictionary {
		scontent = strings.Replace(scontent, Config.Dictionary[i], neutralize(Config.Dictionary[i]), -1)
	}
	stat, _ := file.Stat()
	err := ioutil.WriteFile(spath+"\\"+stat.Name(), []byte(scontent), 0600)
	if err != nil {
		fmt.Println(err)
	}
}

func neutralize(s string) string {
	s2 := []rune(s)
	i := roll(1, len(s2))
	sep2 := []rune(Config.Separetor)
	newstrarr := []rune{}
	newstrarr = append(newstrarr, s2[:i]...)
	newstrarr = append(newstrarr, sep2...)
	newstrarr = append(newstrarr, s2[i:]...)
	return string(newstrarr)
}

func roll(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
