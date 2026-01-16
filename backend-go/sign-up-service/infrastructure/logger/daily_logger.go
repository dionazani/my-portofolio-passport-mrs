package infrastructure_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type DailyFileWriter struct {
	mu          sync.Mutex
	logDir      string
	currentDate string
	file        *os.File
}

func NewDailyFileWriter(dir string) *DailyFileWriter {
	return &DailyFileWriter{logDir: dir}
}

// Write implements the io.Writer interface
func (dw *DailyFileWriter) Write(p []byte) (n int, err error) {
	dw.mu.Lock()
	defer dw.mu.Unlock()

	// 1. Check current date
	newDate := time.Now().Format("2006-01-02")

	// 2. If date changed or file is nil, rotate the file
	if newDate != dw.currentDate {
		if dw.file != nil {
			dw.file.Close() // Safely close the PREVIOUS day's file
		}

		fileName := fmt.Sprintf("%s_logs.txt", newDate)
		path := filepath.Join(dw.logDir, fileName)

		// Create dir if missing
		os.MkdirAll(dw.logDir, 0755)

		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
		if err != nil {
			return 0, err
		}

		dw.file = f
		dw.currentDate = newDate
	}

	return dw.file.Write(p)
}
