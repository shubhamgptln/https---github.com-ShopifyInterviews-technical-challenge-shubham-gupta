package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestInteractiveCommandsUpdateRobotState(t *testing.T) {
	input := strings.NewReader("turn right\nmove 5\nturn left\nmove 2\nstatus\nexit\n")
	var output bytes.Buffer

	if err := run(nil, input, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	if !strings.Contains(output.String(), "Location: (5, 2)\nDirection: North") {
		t.Fatalf("expected final status at (5, 2) facing North, got:\n%s", output.String())
	}
}

func TestMoveRejectsNegativeDistance(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"move", "-1"}, strings.NewReader(""), &output); err == nil {
		t.Fatal("expected negative distance to be rejected")
	}
}

func TestTurnAndMoveInOneShotCommand(t *testing.T) {
	r := robot{direction: north}
	var output bytes.Buffer

	if err := execute([]string{"turn", "left"}, &r, &output); err != nil {
		t.Fatalf("turn returned error: %v", err)
	}
	if err := execute([]string{"move", "4"}, &r, &output); err != nil {
		t.Fatalf("move returned error: %v", err)
	}

	if r.x != -4 || r.y != 0 || r.direction != west {
		t.Fatalf("unexpected robot state: %+v", r)
	}
}
