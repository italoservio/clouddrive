package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func Info(msg string) {
	_, filename, line, _ := runtime.Caller(1)

	fmt.Printf(
		"%s[INFO][%s][%s:%d]%s %s\n",
		green, getDateTime(),
		cutFilename(filename), line,
		reset, msg,
	)
}

func Error(msg string) {
	_, filename, line, _ := runtime.Caller(1)

	fmt.Printf(
		"%s[ERROR][%s][%s:%d]%s %s\n",
		red, getDateTime(),
		cutFilename(filename), line,
		reset, msg,
	)
}

func Warn(msg string) {
	_, filename, line, _ := runtime.Caller(1)

	fmt.Printf(
		"%s[WARN][%s][%s:%d]%s %s\n",
		yellow, getDateTime(),
		cutFilename(filename), line,
		reset, msg,
	)
}

func cutFilename(filename string) string {
	filename_arr := strings.Split(filename, "/")
	filename = strings.Join(filename_arr[len(filename_arr)-3:], "/")
	return filename
}

func getDateTime() string {
	time := time.Now()
	date_time := fmt.Sprintf(
		"%s/%s/%d %s:%s:%s",
		fmt.Sprintf("%02d", time.Day()),
		fmt.Sprintf("%02d", time.Month()),
		time.Year(),
		fmt.Sprintf("%02d", time.Hour()),
		fmt.Sprintf("%02d", time.Minute()),
		fmt.Sprintf("%02d", time.Second()),
	)
	return date_time
}
