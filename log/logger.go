package log

import (
	"go.uber.org/zap"
	"sync"
)

var loggerConf *BaseLogger
var loggerOnce sync.Once

func GetLogInstance() *BaseLogger {
	loggerOnce.Do(func() {
		loggerConf = &BaseLogger{}
	})
	return loggerConf
}

type BaseLogger struct {
	accessLogger   *zap.Logger
	dbAccessLogger *zap.Logger
	debugLogger    *zap.Logger
	warnLogger     *zap.Logger
	errorLogger    *zap.Logger
	infoLogger     *zap.Logger
	upstreamLogger *zap.Logger
	zapLogger *zap.Logger
	zapSugaredLogger      *zap.SugaredLogger
}

func (p *BaseLogger) Init() {
	sugarLogger:=zap.NewExample().Sugar()
	logger,_:=zap.NewProduction()
	p.zapSugaredLogger=sugarLogger
	p.zapLogger=logger
	/*
	conf := config.GetConfInstance()
	var basedir string
	if conf.IsNormalMode {
		basedir = conf.LogPath
	} else {
		cfg := tars.GetServerConfig()
		basedir = fmt.Sprintf("%s/%s/%s", cfg.LogPath, cfg.App, cfg.Server)
	}*/

	/*
	rogger.SetLevel(getServerLogLevel())

	p.accessLogger = rogger.GetLogger("access")
	p.accessLogger.SetDayRoller(basedir, 3)

	p.dbAccessLogger = rogger.GetLogger("db_access")
	p.dbAccessLogger.SetDayRoller(basedir, 3)

	p.debugLogger = rogger.GetLogger("debug")
	p.debugLogger.SetDayRoller(basedir, 3)

	p.warnLogger = rogger.GetLogger("warn")
	p.warnLogger.SetDayRoller(basedir, 3)

	p.errorLogger = rogger.GetLogger("error")
	p.errorLogger.SetDayRoller(basedir, 3)

	p.infoLogger = rogger.GetLogger("info")
	p.infoLogger.SetDayRoller(basedir, 3)

	// 加.是为了方便es日志系统分词
	p.upstreamLogger = rogger.GetLogger("upstream.")
	p.upstreamLogger.SetDayRoller(basedir, 3)

	Debug("basedir:" + basedir)
	Debug("Init BaseLogger success")*/
}

func Debug(v ...interface{}) {
	GetLogInstance().zapLogger.Debug("Debug",zap.Any("Debug",v))
	GetLogInstance().zapSugaredLogger.Debug(v)
}

func Error(v ...interface{}) {
	GetLogInstance().zapLogger.Error("ERROR",zap.Any("ERROR",v))
	GetLogInstance().zapSugaredLogger.Error(v)
}

func Warn(v ...interface{}) {
	GetLogInstance().zapLogger.Warn("Warn",zap.Any("Warn",v))
	GetLogInstance().zapSugaredLogger.Warn(v)
}

func Info(v ...interface{}) {
	GetLogInstance().zapLogger.Info("Info",zap.Any("Info",v))
	GetLogInstance().zapSugaredLogger.Info(v)
}

func Access(v ...interface{}) {
	GetLogInstance().zapSugaredLogger.Info(v)
}

func DBAccess(v ...interface{}) {
	GetLogInstance().zapSugaredLogger.Info(v)
}

func Upstream(v ...interface{}) {
	GetLogInstance().zapSugaredLogger.Info(v)
}
