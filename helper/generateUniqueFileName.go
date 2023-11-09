package helper

import "time"

func GenerateUniqueFileName(filename string) string {
	// use a timestamp, a random string, or any other method to make the filename unique
	// for simplicity, let's append a timestamp to the original filename
	timestamp := time.Now().Format("20060102150405")
	return timestamp + ".jpg"
}
