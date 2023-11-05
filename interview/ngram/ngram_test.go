package ngram

import (
	"testing"
)

func Test_TwoGram(t *testing.T) {
	result := TwoGram([]string{
		"the cow jumped over the moon",
		"the cow and the moon",
	})

	if len(result) != 7 {
		t.Errorf("Expected 7 results, got %d", len(result))
	}

	if result["the cow"] != 2 {
		t.Errorf("Expected 2 for \"the cow\", got %d", result["the cow"])
	}

	if result["cow jumped"] != 1 {
		t.Errorf("Expected 1 for \"cow jumped\", got %d", result["cow jumped"])
	}

	if result["jumped over"] != 1 {
		t.Errorf("Expected 1 for \"jumped over\", got %d", result["jumped over"])
	}

	if result["over the"] != 1 {
		t.Errorf("Expected 1 for \"over the\", got %d", result["over the"])
	}

	if result["the moon"] != 2 {
		t.Errorf("Expected 2 for \"the moon\", got %d", result["the moon"])
	}

	if result["cow and"] != 1 {
		t.Errorf("Expected 1 for \"cow and\", got %d", result["cow and"])
	}

	if result["and the"] != 1 {
		t.Errorf("Expected 1 for \"and the\", got %d", result["and the"])
	}
}
