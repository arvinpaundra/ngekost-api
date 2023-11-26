package log

import "github.com/sirupsen/logrus"

func Logging(msg ...string) *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-02-02 15:04:05",
		PrettyPrint:     true,
	})
	l.SetReportCaller(true)

	return l
}
