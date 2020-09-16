package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var levelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

// InitLogger creates a new logger.
func InitLogger(level string) *zap.Logger {
	config := initConfig(level)

	var err error
	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("Init logger failed:" + err.Error())
	}

	logger = zapLogger
	zap.ReplaceGlobals(zapLogger)

	return logger
}

func initConfig(level string) zap.Config {
	config := zap.NewProductionConfig()

	config.Level = getLevel(level)
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.NameKey = "name"
	std := []string{"stdout", "$HOME/std.log"}
	err := []string{"stderr","$HOME/err.log"}
	config.OutputPaths = std
	config.ErrorOutputPaths = err
	logLevel = config.Level.String()

	return config
}

func getLevel(level string) zap.AtomicLevel {
	zapLevel := levelMapping[level]
	return zap.NewAtomicLevelAt(zapLevel)
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}
