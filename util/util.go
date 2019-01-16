package util

import (
	"os"
	"strconv"
)

type File = os.File

// Int —> Str
func IntToString(intVar int) string {
	return strconv.Itoa(intVar)
}

// 获取当前路径
func GetCurrentDir() (string, error) {
	return os.Getwd()
}

// 打开文件
func OpenFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
}

// 重命名文件
func RenameFile(oldFileName, newFileName string) {
	os.Rename(oldFileName, newFileName)
}

// 检测文件夹/文件是否存在
func CheckPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建文件夹
func CreateDir(path string) (err error) {
	return os.Mkdir(path, os.ModePerm)
}

// 获取系统GOPATH
func GetSystemGoPATH() string {
	return os.Getenv("GOPATH")
}
