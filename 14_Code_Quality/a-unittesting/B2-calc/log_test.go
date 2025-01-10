package calc

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestDozenMeans(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	DozenMeans()

	expected := "DOZEN:12\n"
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}

}
