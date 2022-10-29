package maze

// PostMazeDto Represents the request to solve maze
type PostMazeDto struct {
	Maze map[string]interface{} `json:"maze" validate:"required"`
}

// PostMazeResponseDto represents the response of /maze/solve endpoint
type PostMazeResponseDto struct {
	Exits []string `json:"exits"`
}

// GetMazeDto represents the request to generate maze
type GetMazeDto struct {
	MaxChildren int `json:"maxchildren" validate:"required,min=1"`
}

// GetMazeResponseDto represents the response of /maze/generate endpoint
type GetMazeResponseDto struct {
	Maze map[string]interface{} `json:"maze"`
}
