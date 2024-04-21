package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDailyInterestRate(t *testing.T) {
	viper.Set("APR", 12)

	apr := GetDailyInterestRate()
	assert.Equal(t, "0.00033", fmt.Sprintf("%.5f", apr))
}
