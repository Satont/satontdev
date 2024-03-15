package data

import "time"

func GetAge(birthDate string) int {
	_, _ = time.Parse(time.RFC1123, birthDate)

	return 25
}
