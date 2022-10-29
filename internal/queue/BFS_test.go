package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMazeSolverBFS(t *testing.T) {
	maze := map[string]interface{}{
		"forward": "tiger",
		"left": map[string]interface{}{
			"forward": map[string]interface{}{
				"upstairs": "exit",
			},
			"left": "exit",
		},
		"right": map[string]interface{}{
			"forward": "dead end",
		}}
	res := MazeSolverBFS(maze)
	expected := []string{"[left-> left->]", "[left-> forward-> upstairs->]"}
	assert.ElementsMatch(t, expected, res)
}
