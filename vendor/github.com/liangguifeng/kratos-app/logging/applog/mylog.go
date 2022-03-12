package applog

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	LOGGER_ERR_TYPE      = "err"
	LOGGER_ACCESS_TYPE   = "access"
	LOGGER_BUSINESS_TYPE = "business"

	FLAG_SERVICE_NAME = ":service_name"
	FLAG_LOGGER_TYPE  = ":logger_type"
	FLAG_EMPTY        = "/:-"
)

var (
	globalRootPath string // 日志统一存储根目录
	globalLevel    uint8  // 配置日志级别
	globalAppName  string // app name
)

var levelMap = map[string]uint8{
	"debug": LEVEL_DEBUG,
	"info":  LEVEL_INFO,
	"warn":  LEVEL_WARN,
	"error": LEVEL_ERROR,
}

func InitGlobalConfig(rootPath string, level string, appName string) error {
	if rootPath == "" {
		return errors.New("applog初始化全局配置失败,参数rootPath不能为空")
	}
	if _, ok := levelMap[level]; !ok {
		return errors.New("applog初始化全局配置失败,参数level错误")
	}
	if appName == "" {
		return errors.New("applog初始化全局配置失败,参數appName不能为空")
	}
	globalRootPath = rootPath
	globalLevel = levelMap[level]
	globalAppName = appName
	return nil
}

func contextCall(l *LoggerContext, ctx context.Context) LoggerContextIface {
	fields := make(Fields)

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		mmd := make(map[string][]string)
		for k, v := range md {
			if !strings.HasPrefix(k, "elasticapm") {
				mmd[k] = v
			}
		}

		fields["grpc.metadata"] = mmd
	}

	return l.WithCommonFields(fields)
}

func GetAccessLogger(loggerTag string) (*LoggerContext, error) {
	return getLogger(LOGGER_ACCESS_TYPE, loggerTag, false)
}
func GetBusinessLogger(loggerTag string) (*LoggerContext, error) {
	return getLogger(LOGGER_BUSINESS_TYPE, loggerTag, true)
}

func GetErrLogger(loggerTag string) (*LoggerContext, error) {
	return getLogger(LOGGER_ERR_TYPE, loggerTag, true)
}

func getLogger(loggerType string, loggerTag string, caller bool) (*LoggerContext, error) {
	logFilePath := fmt.Sprintf("%s/%s/%s/log", globalRootPath, loggerType, globalAppName)
	if strings.Contains(globalRootPath, FLAG_SERVICE_NAME) && strings.Contains(globalRootPath, FLAG_LOGGER_TYPE) {
		filePath := strings.Replace(globalRootPath, FLAG_LOGGER_TYPE, loggerType, -1)
		filePath = strings.Replace(filePath, FLAG_SERVICE_NAME, globalAppName, -1)
		logFilePath = fmt.Sprintf("%s/log", filePath)
	} else if strings.Contains(globalRootPath, FLAG_EMPTY) {
		filePath := strings.Replace(globalRootPath, FLAG_EMPTY, "", -1)
		if strings.Contains(filePath, FLAG_SERVICE_NAME) {
			filePath = strings.Replace(filePath, FLAG_SERVICE_NAME, globalAppName, -1)
		}
		if strings.Contains(filePath, FLAG_LOGGER_TYPE) {
			filePath = strings.Replace(filePath, FLAG_LOGGER_TYPE, loggerType, -1)
		}
		logFilePath = fmt.Sprintf("%s/log", filePath)
	}

	log, err := CreateLogger(LoggerConf{
		LogFilePath: logFilePath,
		AppName:     globalAppName,
		Level:       globalLevel,
		Caller:      caller,
		ContextCall: contextCall,
	})
	if err != nil {
		return nil, err
	}
	logger := log.Tag(loggerTag).WithCommonFields(Fields{
		"app": globalAppName,
	})
	return logger.(*LoggerContext), nil
}
