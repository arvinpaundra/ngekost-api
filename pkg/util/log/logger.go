package log

import "github.com/sirupsen/logrus"

type logger struct {
	log     *logrus.Logger
	message string
}

func Logging(msg string) *logger {
	l := logger{}
	l.log = logrus.New()
	l.log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-02-02 15:04:05",
		PrettyPrint:     true,
	})
	l.log.SetReportCaller(true)
	l.message = msg

	return &l
}

func (l *logger) Info() {
	l.log.Info(l.message)
}

func (l *logger) Error() {
	l.log.Error(l.message)
}

func (l *logger) Warn() {
	l.log.Warn(l.message)
}
