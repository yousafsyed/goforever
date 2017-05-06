package main

import (
    "fmt"
    "log"
    "os"
)

var LogFile *os.File
var logger *log.Logger

func logOut(out string) {

    LogFile, err := os.OpenFile(config.LogsFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
        os.Exit(1)
    }
    logger = log.New(LogFile, "", log.Lshortfile|log.LstdFlags)
    logger.Println(out)
    defer LogFile.Close()
}
