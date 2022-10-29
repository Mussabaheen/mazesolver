package maze

import (
	"github.com/Mussabeheen/mazeSolver/internal/queue"
	"github.com/Mussabeheen/mazeSolver/internal/tmpl"
)

type service struct{}

// NewService creates and returns the new maze service
func Newservice() *service {
	return &service{}
}

// SolveMaze solves the JSON maze and returns the list of exits
func (m *service) SolveMaze(request PostMazeDto) []string {
	return queue.MazeSolverBFS(request.Maze)
}

// GetMaze takes max-child as input and creates the maze
func (m *service) GetMaze(request GetMazeDto) (map[string]interface{}, error) {
	return tmpl.GenerateMaze(request.MaxChildren)
}
