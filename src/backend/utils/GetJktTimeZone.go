package utils

import "time"

func GetJktTimeZone() string {
	localTime := time.Now()
	gmtOffset := 7 * 60 * 60

	jakartaTime := time.FixedZone("GMT+7", gmtOffset)

	jakartaNow := localTime.In(jakartaTime)

	return jakartaNow.String()
}
