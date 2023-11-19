package retry

import "fmt"

func Example_fail() {
	attempts, err := attempt(someOperationThatFails, 3)
	fmt.Printf("attempts: %v\n", attempts)
	fmt.Printf("error: %v\n", err != nil)

	// Output:
	// attempts: 3
	// error: true
}

func Example_success() {
	attempts, err := attempt(someOperationThatSucceeds, 3)
	fmt.Printf("attempts: %v\n", attempts)
	fmt.Printf("error: %v\n", err != nil)

	// Output:
	// attempts: 1
	// error: false
}

func Example_failOnceThenSucceed() {
	attempts, err := attempt(getFailOnceThenSucceed(2), 3)
	fmt.Printf("attempts: %v\n", attempts)
	fmt.Printf("error: %v\n", err != nil)

	// Output:
	// attempts: 2
	// error: false
}

func Example_succeedOnAttemptNumber() {
	attempts, err := attempt(getSucceedOnAttemptNumber(3), 3)
	fmt.Printf("attempts: %v\n", attempts)
	fmt.Printf("error: %v\n", err != nil)

	// Output:
	// attempts: 3
	// error: false
}
