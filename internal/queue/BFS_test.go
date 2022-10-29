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
func TestMazeSolverBFS5Levels(t *testing.T) {
	maze := map[string]interface{}{
		"forward": "tiger",
		"left": map[string]interface{}{
			"forward": map[string]interface{}{
				"left": "tiger",
				"upstairs": map[string]interface{}{
					"right": map[string]interface{}{
						"forward": "dead end",
					},
					"forward": map[string]interface{}{"left": "exit"},
				},
			},
			"left": "tiger",
		},
		"right": map[string]interface{}{
			"forward": "dead end",
		}}
	res := MazeSolverBFS(maze)
	expected := []string{"[left-> forward-> upstairs-> forward-> left->]"}
	assert.ElementsMatch(t, expected, res)
}
func TestMazeSolverBFS7Levels(t *testing.T) {
	maze := map[string]interface{}{
		"forward": "tiger",
		"left": map[string]interface{}{
			"forward": map[string]interface{}{
				"left": "tiger",
				"upstairs": map[string]interface{}{
					"right": map[string]interface{}{
						"forward": "dead end",
					},
					"forward": map[string]interface{}{"maindoor": map[string]interface{}{
						"downstairs": "exit",
					}},
				},
			},
			"left": "tiger",
		},
		"right": map[string]interface{}{
			"forward": "dead end",
		}}
	res := MazeSolverBFS(maze)
	expected := []string{"[left-> forward-> upstairs-> forward-> maindoor-> downstairs->]"}
	assert.ElementsMatch(t, expected, res)
}
