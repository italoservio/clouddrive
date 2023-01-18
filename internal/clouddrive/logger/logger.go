package logger

import (
	"fmt"
	"time"
)

func Info(msg string) {
	fmt.Printf("%s[INFO][%s]%s %s\n", Green, getDateTime(), Reset, msg)
}

func Error(msg string) {
	fmt.Printf("%s[ERROR][%s]%s %s\n", Red, getDateTime(), Reset, msg)
}

func Warn(msg string) {
	fmt.Printf("%s[WARN][%s]%s %s\n", Yellow, getDateTime(), Reset, msg)
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
