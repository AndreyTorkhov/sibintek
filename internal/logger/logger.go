package logger

import (
    "log"
    "os"
    "path/filepath"
)

var logger *log.Logger

func init() {
    projectRoot, _ := os.Getwd()
    logPath := filepath.Join(projectRoot, "logs", "app.log")

    if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err == nil {
        logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err == nil {
            logger = log.New(logFile, "", log.LstdFlags)
            logger.Println("Log initialized successfully")
            return
        }
    }
    logger = log.New(os.Stdout, "", log.LstdFlags)
}

func LogError(err error) {
    logger.Printf("ERROR: %v", err)
}

func LogInfo(msg string) {
    logger.Printf("INFO: %s", msg)
}
