package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

func (d direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type robot struct {
	x, y      int64
	direction direction
}

func (r *robot) turnLeft() {
	r.direction = (r.direction + 3) % 4
}

func (r *robot) turnRight() {
	r.direction = (r.direction + 1) % 4
}

func (r *robot) move(distance int64) error {
	switch r.direction {
	case north:
		if r.y > math.MaxInt64-distance {
			return errors.New("movement would exceed coordinate limits")
		}
		r.y += distance
	case east:
		if r.x > math.MaxInt64-distance {
			return errors.New("movement would exceed coordinate limits")
		}
		r.x += distance
	case south:
		if r.y < math.MinInt64+distance {
			return errors.New("movement would exceed coordinate limits")
		}
		r.y -= distance
	case west:
		if r.x < math.MinInt64+distance {
			return errors.New("movement would exceed coordinate limits")
		}
		r.x -= distance
	}
	return nil
}

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(args []string, input io.Reader, output io.Writer) error {
	r := robot{direction: north}
	if len(args) > 0 {
		return execute(args, &r, output)
	}

	fmt.Fprintln(output, "Robot CLI. Type help for commands, or exit to quit.")
	scanner := bufio.NewScanner(input)
	for {
		fmt.Fprint(output, "robot> ")
		if !scanner.Scan() {
			break
		}

		command := strings.Fields(scanner.Text())
		if len(command) == 1 && (command[0] == "exit" || command[0] == "quit") {
			break
		}
		if err := execute(command, &r, output); err != nil {
			fmt.Fprintln(output, "Error:", err)
		}
	}
	return scanner.Err()
}

func execute(args []string, r *robot, output io.Writer) error {
	if len(args) == 0 || (len(args) == 1 && (args[0] == "help" || args[0] == "-h" || args[0] == "--help")) {
		fmt.Fprint(output, `Robot CLI

Commands:
  turn left       Turn 90 degrees counterclockwise
  turn right      Turn 90 degrees clockwise
  move <distance> Move a non-negative integer distance forward
  status          Show the robot's coordinates and direction
  help            Show this help message
  exit            Exit interactive mode
`)
		return nil
	}

	switch args[0] {
	case "turn":
		if len(args) != 2 {
			return errors.New("usage: turn left|right")
		}
		switch args[1] {
		case "left":
			r.turnLeft()
		case "right":
			r.turnRight()
		default:
			return errors.New("turn direction must be left or right")
		}
		fmt.Fprintf(output, "Direction: %s\n", r.direction)
	case "move":
		if len(args) != 2 {
			return errors.New("usage: move <distance>")
		}
		distance, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil || distance < 0 {
			return errors.New("distance must be a non-negative integer")
		}
		if err := r.move(distance); err != nil {
			return err
		}
		fmt.Fprintf(output, "Location: (%d, %d)\n", r.x, r.y)
	case "status":
		if len(args) != 1 {
			return errors.New("usage: status")
		}
		fmt.Fprintf(output, "Location: (%d, %d)\nDirection: %s\n", r.x, r.y, r.direction)
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
	return nil
}
