# 技术参考

## 1. Web框架Fiber
考虑Fiber优点，同时考虑到Gin的代码维护越来越少，但Fiber还是很活跃的，可以看到Fiber的前景还是不错的，所以使用Fiber作为web框架
源码：https://github.com/gofiber/fiber
文档：gofiber.io
实例：https://github.com/gofiber/recipes
## 2. 配置文件处理 go-toml/v2
配置文件采用TOML格式处理，使用 github.com/pelletier/go-toml/v2 库，参考文章：https://www.cnblogs.com/realcp1018/p/14863128.html

## 3. ORM框架， gorm.io/gorm
考虑到gorm经过很多应用的检验，在性能等方面还是靠得住的，同时现在也支持*string类指针类型，同时也支持直接手写sql模式，所以选择gorm

## 4. env
- joho/godotenv   A Go port of Ruby's dotenv library (Loads environment variables from .env files)     6K     2023-4-13
- caarlos0/env    A simple and zero-dependencies library to parse environment variables into structs.  3.4K   2023-4-13
fiber的实例库里也用godotenv，所以本框架使用godotenv

## 5. email（选用jordan-wright/email）
email是一个很重要的功能，特别是需要消息提醒，定时报告发送等都会用到，关于email的库，go并不多，其中有两个星数较多的库
- https://github.com/jordan-wright/email   2.3K  最后更新时间2021年1月
- https://github.com/go-gomail/gomail      3.9K  最后更新时间2016年4月
两个库都很不错，感觉第一个使用起来更简单，也有连接池的概念，当然第二个也有类似，但从简单性来说，第一个更好
除了官方的文档外，也可以快速阅读：https://blog.csdn.net/qq_45533884/article/details/106334063
所以选择使用第一个库

## 6. Log (选用zap)
log框架第三方库非常丰富，可供选择的主要有 logrus 和 zap，考虑zap大厂出品，性能十分优异，所以选择zap，并配合gopkg.in/natefinch/lumberjack.v2，(注意不是github.com/natefinch/lumberjack)使用,安装使用：
```shell
   go get -u go.uber.org/zap
   go get -u gopkg.in/natefinch/lumberjack.v2
```

- https://github.com/sirupsen/logrus  22.5K  功能强大，丰富
- https://github.com/uber-go/zap      18.5K  性能十分优异，uber大厂出品


## 7. cmd 命令行
   https://github.com/spf13/cobra，太优秀了，支持子命令，使用也很简单，hugo等很多工具都使用该库，没必要再考虑其他了