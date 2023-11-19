package retry

// operation is a function that succeeds on the (one-based) attempt.
type operation func(succeedAttempt int) error

func attempt(op operation, attempts int) (int, error) {
	var err error
	for i := 1; i <= attempts; i++ {
		err = op(i)
		if err == nil {
			return i, nil
		}
	}
	return attempts, err
}
