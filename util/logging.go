package util

import (
	"fmt"
	"os"
)

const (
	loggingSwitchEnvVar = "GO-GROUPSESSION-LOGGING-SWITCH"
	loggingSwitchOn     = "LOGGING-ON"
)

// LoggingSwitchOn ロギングオン
func LoggingSwitchOn() error {
	return os.Setenv(loggingSwitchEnvVar, loggingSwitchOn)
}

// LoggingSwitchOff ロギングオフ
func LoggingSwitchOff() error {
	return os.Setenv(loggingSwitchEnvVar, "")
}

// IsLoggingOn ロギングするか
func IsLoggingOn() bool {
	return os.Getenv(loggingSwitchEnvVar) == loggingSwitchOn
}

// Log ログ出し
func Log(log string) {
	if !IsLoggingOn() {
		return
	}
	fmt.Println(log)
}
