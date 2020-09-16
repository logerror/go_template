package log

import (
	"time"
	stringutil "welights.net/go_template/pkg/utils/stringutils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	var options LoggerOptions
	options.Level = level
	config := initConfig(options)

	var err error
	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("Init logger failed:" + err.Error())
	}

	logger = zapLogger
	zap.ReplaceGlobals(zapLogger)

	return logger
}

type LoggerOptions struct {
	Level            string
	OutputPaths      string
	ErrorOutputPaths string
}

func InitLoggerWithOptions(options LoggerOptions) *zap.Logger {
	config := initConfig(options)

	var err error
	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("Init logger failed:" + err.Error())
	}

	logger = zapLogger
	zap.ReplaceGlobals(zapLogger)

	return logger
}

func initConfig(options LoggerOptions) zap.Config {
	config := zap.NewProductionConfig()
	//默认 error
	if stringutil.IsBlank(options.Level) {
		options.Level = "error"
	}

	config.Level = getLevel(options.Level)

	std := []string{"stdout"}
	err := []string{"stderr"}

	if stringutil.IsNotBlank(options.OutputPaths) {
		std = append(std, options.OutputPaths)
	}

	if stringutil.IsNotBlank(options.OutputPaths) {
		err = append(err, options.ErrorOutputPaths)
	}

	config.OutputPaths = std
	config.ErrorOutputPaths = err

	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.NameKey = "name"
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		encodeTimeLayout(t, "2006-01-02 15:04:05.000", enc)
	}

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
