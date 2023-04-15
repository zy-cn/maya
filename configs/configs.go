// Copyright 2023 zhangying. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 配置文件采用TOML

/*
Package configs 配置数据信息
*/
package configs

import (
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type Config struct {
	App struct {
		Title string `toml:"titel"`
	} `toml:"app"`

	Server struct {
		RunMode string `toml:"runMode"`
		Host    string `toml:"host"`
		Port    int32  `toml:"port"`
	}

	Jwt struct {
		Secret string `toml:"secret"`
		Expire int    `toml:"expire"`
	}

	Database struct {
		DbType       string `toml:"dbType"`
		Host         string `toml:"host"`
		DbName       string `toml:"dbName"`
		UserName     string `toml:"userName"`
		Password     string `toml:"password"`
		TablePrefix  string `toml:"tablePrefix"`
		Charset      string `toml:"charset"`
		ParseTime    bool   `toml:"parseTime"`
		MaxIdleConns int    `toml:"maxIdleConns"`
		MaxOpenConns int    `toml:"maxOpenConns"`
	}

	SMTPInfo struct {
		Enabled  bool   `toml:"enabled"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		IsSSL    bool   `toml:"isSSL"`
		UserName string `toml:"userName"`
		Password string `toml:"password"`
		From     string `toml:"from"`
		PoolSize int    `toml:"poolSize"`
	}

	Log struct {
		FileName   string `toml:"fileName"`
		MaxSize    int    `toml:"maxSize"`    // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups int    `toml:"maxBackups"` // 保留旧文件的最大个数
		MaxAge     int    `toml:"maxAge"`     // 保留旧文件的最大天数
		Compress   bool   `toml:"compress"`   // 是否压缩/归档旧文件
	}
}

func GetConfig() (*Config, error) {
	config := new(Config)
	readFile, err := os.OpenFile("./configs/config.toml", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err.Error())
		panic("read config.toml failed")
	}
	defer readFile.Close()

	decoder := toml.NewDecoder(readFile)
	//defer decoder.Close()
	decoder.Decode(config)

	return config, nil
}
