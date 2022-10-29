package queue

import (
	"fmt"
)

func MazeSolverBFS(maze map[string]interface{}) []string {
	resExits := []string{}
	queue := Queue{}
	for key, value := range maze {
		queue.Enqueue(Node{
			Key:        key,
			Val:        value,
			ParentNode: nil,
		})
	}
	for queue.Size != 0 {
		element := queue.Dequeue()
		elementMap, ok := element.Val.(map[string]interface{})
		if !ok {
			exitPath := []string{}
			if element.Val == "exit" {
				parentNode := element.ParentNode
				exitPath = append([]string{element.Key + "->"}, exitPath...)
				for parentNode != nil {
					exitPath = append([]string{parentNode.Key + "->"}, exitPath...)
					parentNode = parentNode.ParentNode
				}
			} else {
				continue
			}
			resExits = append(resExits, fmt.Sprint(exitPath))
		}
		for key, val := range elementMap {
			queue.Enqueue(
				Node{
					Key:        key,
					Val:        val,
					ParentNode: &element,
				},
			)
		}

	}
	return resExits
}
