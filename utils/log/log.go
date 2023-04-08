package log

import (
	"encoding/json"
	"fmt"
)

func Error(meta interface{}, err error, message string) {
	json, _ := json.Marshal(map[string]interface{}{
		"err":      err.Error(),
		"metadata": meta,
		"message":  message,
	})
	fmt.Printf("%v\n", string(json))
}

func Info(message string) {
	fmt.Printf("[INFO] %v\n", message)
}

func Warn(message string) {
	fmt.Printf("[WARN] %v\n", message)
}
