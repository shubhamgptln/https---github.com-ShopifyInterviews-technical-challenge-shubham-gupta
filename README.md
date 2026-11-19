# Robot CLI

A 2D robot simulator. The robot starts at `(0, 0)` facing north. North
increases `y`, east increases `x`, and the robot turns in 90-degree increments.
Its state is held in memory for the lifetime of the process.

## Interactive use

Run the CLI without arguments to start an interactive session:

```sh
go run .
```

Enter commands in that session to maintain the robot's position and direction:

```text
turn right
move 5
status
```

Use `exit` or `quit` to end the session.

## One-shot commands

Commands can also be run directly:

```sh
go run . turn left
go run . move 3
go run . status
```

Each one-shot invocation starts a new robot at `(0, 0)` facing north. Use
interactive mode when state must be retained between commands.

## Commands

- `turn left` / `turn right` — rotate 90 degrees.
- `move <distance>` — move forward a non-negative integer distance.
- `status` — print the current `(x, y)` coordinates and facing direction.
