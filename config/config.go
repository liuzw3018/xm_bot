package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/20 9:08
 * @File:
 * @Description: 机器人配置信息加载
 * @Version:
 */

// Config 机器人配置信息
type Config struct {
	Server `yaml:"Server"`
	TXBot  `yaml:"TxBot"`
	CQHttp CQHttp `yaml:"CQHttp"`
}

// Server 机器人服务器配置
type Server struct {
	Host       string   `yaml:"Host"`
	Port       int      `yaml:"Port"`
	SuperUsers []string `yaml:"SuperUsers"`
	NickName   []string `yaml:"NickName"`
	LocalProxy string   `yaml:"LocalProxy"`
	LogPath    string   `yaml:"LogPath"`
}

// CQHttp CQHttp服务器配置
type CQHttp struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

// TXBot 天行机器人配置，可不使用
type TXBot struct {
	TxBotKey  string `yaml:"TxBotKey"`
	TxBotUrl  string `yaml:"TxBotUrl"`
	TxMod     int    `yaml:"TxMod"`
	TxPriv    int    `yaml:"TxPriv"`
	TxResType int    `yaml:"TxResType"`
}

// @title:    	  LoadConfig
// @description:  加载配置
// @auth:         liuzw3018
// @param:        nil
// @return:       *global.GConfig
func LoadConfig() *Config {
	f, err := ioutil.ReadFile("./config_env.yaml")
	if err != nil {
		log.Panicf("读取配置文件错误：%s\n", err)
	}

	c := Config{}
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		log.Panicf("加载配置文件错误：%s\n", err)
	}
	return &c
}
