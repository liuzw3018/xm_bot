package xmlog

import (
	"io"
	"os"
	"time"
	"ximan/global"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/20 11:21
 * @File:
 * @Description: //TODO $
 * @Version:
 */

const (
	FORMAT = "20060102"
	//LineFeed 换行
	LineFeed = "\r\n"
)

var (
	LogPath = global.GConfig.LogPath
)

//WriteLog return error
func WriteLog(fileName, msg string) error {
	//以天为基准,存日志
	var path = LogPath + time.Now().Format(FORMAT) + "/"
	//fmt.Println(path)
	if !IsExist(path) {
		return CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)

	f, err = os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, LineFeed+msg)

	defer func() {
		_ = f.Close()
	}()
	return err
}

//CreateDir  文件夹创建
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	_ = os.Chmod(path, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
