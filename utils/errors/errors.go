package internal_errors

import "errors"

var (
	InvalidIdParameter = errors.New("invalid id parameter")

	InsufficientBalance = errors.New("insufficient balance")

	DailyRateError = errors.New("error in daily rate calculation")

	JobAlreadyCompleted = errors.New("job is already completed")
	JobCreationFailed   = errors.New("job creation failed")
	JobProcessingFailed = errors.New("job processing failed")
)
