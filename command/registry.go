package command

import (
	"fmt"
	"sort"
	"strings"
)

type Registry map[string]Command

var registry = make(Registry)

func List() {
	keys := make([]string, 0, len(registry))
	for k := range registry {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Available commands:", strings.Join(keys, ", "))
}

func Register(name string, command Command) {
	registry[name] = command
}
