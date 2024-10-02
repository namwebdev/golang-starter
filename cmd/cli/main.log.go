package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello!!")

	// logger := zap.NewExample()
	// logger.Info("Hello!!", zap.String("name", "Me"), zap.Int("age", 18))
	encoder := getEncodeLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core)

	logger.Info("Hello!!", zap.String("name", "Me"), zap.Int("age", 18))
	logger.Error("Error!!", zap.String("name", "Me"), zap.Int("age", 18))
}

func getEncodeLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.Lock(os.Stdout)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
