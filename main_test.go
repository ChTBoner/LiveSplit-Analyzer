package main

import (
	"fmt"
	"os"
	"testing"
)


func TestDecodeRun(t *testing.T) {
	filenames := []string{"sm", "alttp", "portal"}

	for _, name := range filenames {
		filepath := fmt.Sprintf("test_data/%s.lss", name)
		data, err := os.ReadFile(filepath)
		if err != nil {
			t.Fatalf("Could not open file %s.", filepath)
		}
		run := DecodeRun(data)
		t.Logf(run.GameName)
		t.Logf("%T", run.AutoSplitterSettings)
	}
}