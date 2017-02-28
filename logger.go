package loggers

import (
	"net/http"
	"time"

	logrus "github.com/Sirupsen/logrus"
	"github.com/steve-winter/logrus-stack"
)

func init() {
	callerLevels := logrus.AllLevels
	stackLevels := []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}
	ignoreList := []string{"loggers/logger"}
	logrus.AddHook(logrus_stack.NewHook(callerLevels, stackLevels, ignoreList))
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(s string, args ...interface{}) {
	logrus.Errorf(s, args...)
}

func Infof(s string, args ...interface{}) {
	logrus.Infof(s, args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(s string, args ...interface{}) {
	logrus.Fatalf(s, args...)
}

func LogWebRequest(inner http.Handler, path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		Infof(
			"%s  %s  %s  %s",
			r.Method,
			r.RequestURI,
			path,
			time.Since(start),
		)
	})
}
