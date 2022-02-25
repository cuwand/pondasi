package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

type LogConfig struct {
	logger      zerolog.Logger
	serviceName string
}

type applicationInfo struct {
	appName    string
	appVersion string
	logRootDir string
}

var appLoggerConfig LogConfig

func init() {
	app := "app"

	appLoggerConfig = NewLogger(app)

	_, logDirectory := getDirectory(app)
	go rotation(logDirectory)
}

func appInfo() applicationInfo {
	appName, ok := os.LookupEnv("APP_NAME")

	if !ok {
		appName = fmt.Sprintf("PONDASI-%v", time.Now().Format("02012006"))
	}

	appVersion, ok := os.LookupEnv("APP_VERSION")

	if !ok {
		appVersion = "0.0.0-PONDASI"
	}

	logRootDir, ok := os.LookupEnv("LOG_ROOT_DIR")

	if !ok {
		logRootDir = "./logs"
	}

	return applicationInfo{
		appName:    appName,
		appVersion: appVersion,
		logRootDir: logRootDir,
	}
}

func getDirectory(serviceName string) (io.Writer, string) {
	appInfo := appInfo()

	appDir := fmt.Sprintf("%s-%s", appInfo.appName, appInfo.appVersion)
	logDir := fmt.Sprintf("%s/%s/%s.log", appInfo.logRootDir, appDir, serviceName)
	dir := path.Dir(logDir)

	// Generate Root Directory
	if _, err := os.Stat(appInfo.logRootDir); os.IsNotExist(err) {
		err = os.Mkdir(appInfo.logRootDir, os.ModePerm)

		if err != nil {
			panic(err.Error())
		}
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)

		if err != nil {
			panic(err.Error())
		}
	}

	f, err := os.OpenFile(logDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err.Error())
	}

	return io.MultiWriter(f, os.Stdout), fmt.Sprintf("%s/%s", appInfo.logRootDir, appDir)
}

func NewLogger(serviceName string) (logger LogConfig) {
	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		panic(err)
	}

	writer, _ := getDirectory(serviceName)

	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(loc)
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("| ")
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatCaller = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.Out = writer

	//multi := zerolog.MultiLevelWriter(output, os.Stdout)
	fileLogger := zerolog.New(output).
		With().
		CallerWithSkipFrameCount(3).
		Str("services", serviceName).
		Timestamp().
		Logger()

	logger = LogConfig{
		logger:      fileLogger,
		serviceName: serviceName,
	}

	return logger
}
