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
 * @Description: 日志文件处理
 * @Version:
 */

const (
	FORMAT   = "20060102" // 日期格式
	LineFeed = "\r\n"     //LineFeed 换行
)

var (
	LogPath = global.GConfig.LogPath // 日志文件保存路径
)

// @title:    	  WriteLog
// @description:  将日志写入文件
// @auth:         liuzw3018
// @param:        fileName, msg string
// @return:       error
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
	_, err = io.WriteString(f, msg+LineFeed)

	defer func() {
		_ = f.Close()
	}()
	return err
}

// @title:    	  CreateDir
// @description:  创建日志文件夹
// @auth:         liuzw3018
// @param:        path string
// @return:       error
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	_ = os.Chmod(path, os.ModePerm)
	return nil
}

// @title:    	  IsExist
// @description:  判断文件是否存在
// @auth:         liuzw3018
// @param:        f string
// @return:       bool
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
