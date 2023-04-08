package log

import "fmt"

func Error(meta interface{}, err error, message string) {
	fmt.Printf("%v", map[string]interface{}{
		"err":      err,
		"metadata": meta,
		"message":  message,
	})
}

func Info(message string) {
	fmt.Printf("[INFO] %v\n", message)
}

func Warn(message string) {
	fmt.Printf("[WARN] %v\n", message)
}
