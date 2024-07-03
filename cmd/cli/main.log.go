package main

import (
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile(".log/log.txt", os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
