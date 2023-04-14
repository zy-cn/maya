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
