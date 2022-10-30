# Maze Solver

MazeSolver can solve and generate maze.

## API Reference

#### POST mazesolve

```http
  POST /maze/solve
```

| Parameter | Type   | Description                       |
| :-------- | :----- | :-------------------------------- |
| `maze`    | `json` | **Required**, maze in json format |

#### Get maze

```http
  GET /maze/generate
```

| Parameter     | Type  | Description                               |
| :------------ | :---- | :---------------------------------------- |
| `maxchildren` | `int` | **Required**. number of children per node |

## Installation

Install mazesolver with go tidy

```bash
  cd my-project
  make run
```

## Running Tests

To run tests, run the following command

```bash
  make test
```
