package nithian

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTranslateQuickBrownFox(t *testing.T) {
	input := "The quick brown fox jumps over the lazy dog."
	output := TranslateText(input)

	jsonOutput, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		t.Fatalf("Failed to serialize ideogram output: %v", err)
	}

	fmt.Println(string(jsonOutput))

	if len(output) == 0 {
		t.Errorf("Expected non-empty ideogram output for input: %s", input)
	}

	// Optional: assert specific ideograms are present if known
	found := false
	for _, ideogram := range output {
		if ideogram.Ideogram == "sun" { // for example
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected ideogram 'sun' to appear in translation")
	}
}
