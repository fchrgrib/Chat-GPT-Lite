package utils

import "time"

func GetJktTimeZone() string {
	localTime := time.Now()

	jakartaTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	jakartaNow := localTime.In(jakartaTime)

	return jakartaNow.String()
}
