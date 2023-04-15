// 日志处理
// by zhangying 2023.4.15
// 参考 https://github.com/uber-go/zap

package logger

import (
	"errors"
	"maya/configs"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func InitLogger(config *configs.Config) (*zap.Logger, *zap.SugaredLogger) {
	writeSyncer := getLogWriter(config)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
	defer logger.Sync()
	sugar = logger.Sugar()

	return logger, sugar
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getLogWriter(config *configs.Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Log.FileName,   // 日志文件位置
		MaxSize:    config.Log.MaxSize,    // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: config.Log.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     config.Log.MaxAge,     // 保留旧文件的最大天数
		Compress:   config.Log.Compress,   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func test() {
	sugar.Debug("这是一条日志", zap.String("name", "zhangSang"), zap.Int("age", 18)) // zap.NewProduction() 默认不输出该级别日志
	sugar.Info("这是一条日志", zap.String("name", "zhangSang"), zap.Int("age", 18))
	sugar.Error("这是一条日志", zap.String("name", "zhangSang"), zap.Error(errors.New("错误信息")))

	logger.Debug("这是一条日志", zap.String("name", "zhangSang"), zap.Int("age", 18)) // zap.NewProduction() 默认不输出该级别日志
	logger.Info("这是一条日志", zap.String("name", "zhangSang"), zap.Int("age", 18))
	logger.Error("这是一条日志", zap.String("name", "zhangSang"), zap.Error(errors.New("错误信息")))

	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "http://yyyy.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "http://xxxx.com")

	sugar.Debugf("查询用户信息开始 id:%d", 1)
	sugar.Infof("查询用户信息成功 name:%s age:%d", "zhangSan", 20)
	sugar.Errorf("查询用户信息失败 error:%v", "未该查询到该用户信息")
}

//参考文章
//https://blog.csdn.net/qq_44011116/article/details/125661345
//https://blog.csdn.net/weixin_52000204/article/details/126651319
