package main

import (
	"fmt"
	"log"
	"maya/configs"
	"maya/global"
	"maya/internal/dao"
	"maya/internal/middleware"
	"maya/internal/routers"
	mayaEmail "maya/pkg/email"
	mayaLogger "maya/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	//全局设置时区为东8区
	cst8zone := time.FixedZone("CST", 8*3600)
	time.Local = cst8zone

	app := fiber.New()
	app.Static("/", "./public")
	routers.SetupRoutes(app)

	var whitePaths []string = nil
	app.Use(middleware.JwtProtected(whitePaths))

	go func() {
		app.Listen((":9002"))
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	//Shutdown gracefully shuts down the server without interrupting any active connections. Shutdown works by first closing all open listeners and then waits indefinitely for all connections to return to idle before shutting down.
	// ShutdownWithTimeout will forcefully close any active connections after the timeout expires.
	//if err :=   app.Shutdown() ; err != nil {
	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func init() {
	//读取配置文件
	var err error
	_config, err := configs.GetConfig()
	if err != nil {
		fmt.Println("read config failed:", err)
		return
	}

	fmt.Printf("Config: %d", _config.Server.Port)
	global.Config = _config

	//初始化db
	//global.DBEngine, err = model.NewDBEngine(_config)
	err = dao.NewDBEngine(_config)
	if err != nil {
		fmt.Println("get DBEngine failed:", err)
		return
	}

	//初始化email
	if _config.SMTPInfo.Enabled {
		mayaEmail.InitEmail(_config)
	}

	//初始化logger
	global.Logger, global.Sugar = mayaLogger.InitLogger(_config)

	global.Logger.Info("系统启动成功", zap.String("name", "maya system"))

}
