package maze

import (
	"net/http"

	"github.com/Mussabeheen/mazeSolver/httputil"
	"github.com/gorilla/mux"
)

var errExitNotFound = "sorry"

// Service provides an interface for Solving and Creating the maze
type Service interface {
	SolveMaze(request PostMazeDto) []string
	GetMaze(request GetMazeDto) error
}

// Controller handles requests for the maze endpoints
type Controller struct {
	Service Service
}

// NewController creates a Controller
func NewController(service Service) *Controller {
	return &Controller{
		Service: service,
	}
}

// RegisterRoutes registers the endpoints for maze
func (m *Controller) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/maze/solve", m.PostMazeSolution).Methods("POST")
	r.HandleFunc("/maze/generate", m.GetMaze).Methods("GET")
}

// PostMazeSolution gets maze in JSON form and returns the solved list of exits
func (m *Controller) PostMazeSolution(w http.ResponseWriter, r *http.Request) {
	var request PostMazeDto
	if err := httputil.ParseBody(r, &request); err != nil {
		httputil.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	res := m.Service.SolveMaze(request)
	if len(res) == 0 {
		httputil.RespondWithError(w, http.StatusNotFound, errExitNotFound)
		return
	}
	httputil.RespondWithJSON(w, 200, PostMazeResponseDto{Exits: res})
}

// GetMaze gets number of max-child and returns the maze in JSON
func (m *Controller) GetMaze(w http.ResponseWriter, r *http.Request) {
	var request GetMazeDto
	if err := httputil.ParseQuery(r, &request); err != nil {
		httputil.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	m.Service.GetMaze(request)
}
