package logger

import (
	"github.com/salawatbro/raxmet/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var Logger *zap.Logger

func InitLogger(env string) {
	var logLevel zapcore.LevelEnabler
	var encoderConfig zapcore.EncoderConfig
	var output zapcore.WriteSyncer

	if strings.ToLower(env) == "production" {
		encoderConfig = zap.NewProductionEncoderConfig()
		logLevel = zapcore.InfoLevel
		file, err := os.Create("app.log")
		if err != nil {
			panic(err)
		}
		output = zapcore.AddSync(file)
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		logLevel = zapcore.DebugLevel
		output = zapcore.AddSync(os.Stdout)
	}

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(constants.TimestampFormat)

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, output, logLevel)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.FatalLevel))
}

func CloseLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
