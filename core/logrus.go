package core

import (
	"bytes"
	"fmt"
	"go_admin_api/config"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 颜色常量，对应终端ANSI颜色码
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// LogFormatter 自定义日志格式化结构体
type LogFormatter struct{}

// Format 实现 logrus.Formatter 接口，自定义日志格式
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同日志级别设置颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	// 初始化缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日期格式（示例：2006-01-02 15:04:05）
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	// 处理调用者信息（文件名和行号）
	fileName := "???"
	line := 0
	if entry.HasCaller() {
		fileName = path.Base(entry.Caller.File)
		line = entry.Caller.Line
	}

	// 拼接带颜色的日志内容（格式：[时间] [颜色化级别] [文件名:行号] 消息）
	// \x1b[0m 用于重置颜色，避免后续输出受影响
	fmt.Fprintf(b, "\x1b[%dm[%s] [%s] [%s:%d] %s\x1b[0m\n",
		levelColor,
		timestamp,
		entry.Level.String(),
		fileName,
		line,
		entry.Message,
	)

	return b.Bytes(), nil
}

// InitLogger 初始化日志
func InitLogger() *logrus.Logger {
	mLog := logrus.New()      // 新建一个实例
	mLog.SetOutput(os.Stdout) // 设置输出类型为标准输出
	// 开启返回函数名和行号（从配置获取开关）
	mLog.SetReportCaller(config.Config.Logger.ShowLine)
	mLog.SetFormatter(&LogFormatter{}) // 设置自己定义的Formatter

	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // 设置最低的Level
	InitDefaultLogger()  // 不注释即启用全局log
	return mLog
}

// InitDefaultLogger 初始化全局log
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout) // 设置输出类型为标准输出
	// 开启返回函数名和行号（从配置获取开关）
	logrus.SetReportCaller(config.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{}) // 设置自己定义的Formatter

	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低的Level
}
