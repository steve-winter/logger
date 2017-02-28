package loggers

import "github.com/Sirupsen/logrus"

type logEntry struct {
	entry logrus.Entry
}

func (e *logEntry) WithTitle(s string) *logEntry {

	return e.WithString("title", s)
}

func (e *logEntry) WithString(title string, field string) *logEntry {

	e.entry.WithField(title, field)

	return e
}

func (e *logEntry) ToEntry() logrus.Entry {
	return e.entry
}
