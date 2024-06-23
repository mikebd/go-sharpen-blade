package retry

import "fmt"

func someOperationThatFails(_ int) error {
	return fmt.Errorf("this failed")
}

func someOperationThatSucceeds(_ int) error {
	return nil
}

func getFailOnceThenSucceed(_ int) operation {
	attemptCount := 0
	attemptCount++

	return func(_ int) error {
		if attemptCount == 1 {
			attemptCount++
			return fmt.Errorf("failed attempt 1")
		}
		return nil
	}
}

func getSucceedOnAttemptNumber(attempt int) operation {
	attemptCount := 0
	attemptCount++

	return func(_ int) error {
		if attemptCount != attempt {
			attemptCount++
			return fmt.Errorf("failed attempt 1")
		}
		return nil
	}
}
