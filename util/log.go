package util

import (
	"fmt"
	"time"
)

// LogInfo prints statement to screen with timestamp.
func LogInfo(msg string) {
	fmt.Printf("[INFO][%s] %s\n", time.Now(), msg)
}

// LogErr prints error statement to screen with timestamp.
func LogErr(trace string, err error) {
	if err == nil {
		fmt.Printf("[EROR][%s] %s \n", time.Now(), trace)
	} else {
		fmt.Printf("[EROR][%s] %s - %s\n", time.Now(), trace, err.Error())
	}
}
