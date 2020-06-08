package Helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const SYS_CONFIG_PATH = "/app.yaml" //全局配置文件
const SLASH = "/"

var WorkDir string

func init() {
	wd, _ := os.Getwd()
	WorkDir = strings.Replace(wd, "\\", "/", -1)
}

//判断目录是否存在
func IsFilePathExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//判断api文件是否存在
func IsFileExist(filepath string) bool {
	fi, err := os.Stat(filepath)
	if err != nil {
		log.Println(err)
		return false
	}
	if fi.IsDir() {
		return false
	}
	return true
}

//判断接口文件是否存在
func IsApiFileExist(filename string) bool {
	fp := GetApiFile(filename)
	return IsFileExist(fp)
}

//辅助函数
func GetApiFile(filename string) string {
	return filename + ".go"
}

//调用os.MkdirAll递归创建文件夹
func CreateMutiDir(dir string, target string) (string, error) {
	targetPath := strings.TrimRight(WorkDir+"/"+dir, "/")
	if !PathExists(targetPath) {
		err := os.MkdirAll(targetPath, 066)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return "", err
		}
	}
	return fmt.Sprintf("%s%s%s.go", targetPath, SLASH, target), nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func PathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//读取文件
func ReadFile(f string) []byte {
	file, err := os.OpenFile(f, os.O_RDWR, 0666)
	if err != nil {
		log.Println("open file err:", err)
		return nil
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read file err:", err)
		return nil
	}
	return b
}

//遍历文件夹，把静态文件读取出来
//map["SERVICE_TPL"]=文件里面的内容
func LoadResource(dir string) map[string]string {
	ret := make(map[string]string)
	dirlist, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("read dir err", err)
	}
	for _, fi := range dirlist {
		if fi.IsDir() {
			continue //目前只处理一级 ，我们这个项目只需要这样
		}
		//这里统一把.换成下划线 并且大写
		keyName := strings.ToUpper(strings.Replace(fi.Name(), ".", "_", -1))
		ret[keyName] = string(ReadFile(dir + "/" + fi.Name()))
	}
	return ret
}
