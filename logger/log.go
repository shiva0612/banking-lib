package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func init() {
	config := zap.NewProductionConfig()

	config.Level.SetLevel(zapcore.DebugLevel)
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.Encoding = "console"

	zlog, _ := config.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)) //prints stacktrace only for error level and above
	log = zlog.Sugar()
}

func Fatal(args ...interface{}) {
	log.Fatalln(args)
}
func Panic(args ...interface{}) {
	log.Panicln(args)
}
func Info(args ...interface{}) {
	log.Infoln(args)
}
func Debug(args ...interface{}) {
	log.Debugln(args)
}
func Error(args ...interface{}) {
	log.Errorln(args)
}
