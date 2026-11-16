package livelychart

import (
	"strings"
	"testing"
)

func TestRenderEscapesLabel(t *testing.T) {
	svg, err := Render([]Bar{{Label: `<x>`, Value: 3}}, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(svg, `<x>`) || !strings.Contains(svg, `&lt;x&gt;`) {
		t.Fatal("label was not escaped")
	}
}

func TestRenderRejectsNegative(t *testing.T) {
	if _, err := Render([]Bar{{Label: "x", Value: -1}}, Options{}); err == nil {
		t.Fatal("expected error")
	}
}
