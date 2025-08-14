package giftExchange

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo/mutable"
)

type person struct {
	name  string
	email string
}

type senderReceiver struct {
	sender   person
	receiver person
}

func giftExchange(file *os.File) ([]senderReceiver, error) {
	people, err := parse(file)
	if err != nil {
		return nil, err
	}

	if len(people) < 2 {
		return nil, fmt.Errorf("not enough people, require at least 2, got %d", len(people))
	}

	result := make([]senderReceiver, 0, len(people))
	mutable.Shuffle(people)
	for i := 0; i < len(people)-1; i += 1 {
		result = append(result, senderReceiver{
			sender:   people[i],
			receiver: people[i+1],
		})
	}
	result = append(result, senderReceiver{
		sender:   people[len(people)-1],
		receiver: people[0],
	})

	return result, nil
}

// parse parses the CSV file, discards the header and returns a list of person.
// It also closes the file.
func parse(file *os.File) ([]person, error) {
	scanner := bufio.NewScanner(file)
	defer file.Close()

	firstLine := true
	var result []person

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if firstLine {
			firstLine = false
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) == 2 {
			result = append(result, person{fields[0], fields[1]})
		} else {
			fmt.Fprintf(os.Stderr, "invalid line: %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return result, nil
}
