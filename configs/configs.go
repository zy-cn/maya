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
