package utils

import (
	"github.com/spf13/viper"
	"time"
)

func GetDailyInterestRate() float64 {
	apr := viper.GetFloat64("APR")

	return apr / 100 / float64(daysInYear())
}

func daysInYear() int {
	// Extract the year from the current time
	year := time.Now().Year()

	// Determine the number of days based on leap year status
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return 366
	}

	return 365

}
