package storage

import "os"

type LogWriter struct {
	file *os.File
}

func NewLogWriter(path string) (*LogWriter, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	return &LogWriter{file: file}, nil
}

func (l *LogWriter) Append(record []byte) error {
	_, err := l.file.Write(record)
	if err != nil {
		return err
	}
	return nil
}
