package log

import (
	"os"
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	GLogger *logrus.Logger
)

func init() {
	GLogger = logrus.New()
	GLogger.SetFormatter(&logrus.TextFormatter{})
	GLogger.SetLevel(logrus.InfoLevel)

	logFile := filepath.Base(os.Args[0])
	ConfigLocalFilesystemLogger("./logfile", logFile, time.Hour*240, time.Hour*1)

	//logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY, 0)
	//if err != nil {
	//	panic(err)
	//}
	//GLogger.SetOutput(logFile)
}

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	infowriter, err := rotatelogs.New(
		baseLogPaht+"_info_"+"%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(baseLogPaht+"_info_"+"%Y%m%d%H%M.log"), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                                  // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime),                      // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	writer, err := rotatelogs.New(
		baseLogPaht+"%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(baseLogPaht+"%Y%m%d%H%M.log"), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                         // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime),             // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  infowriter,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	GLogger.AddHook(lfHook)
}
