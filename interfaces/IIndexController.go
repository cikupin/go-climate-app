package interfaces

import (
	"net/http"
)

// IIndexController - interface for IndexController
type IIndexController interface {
	Index(w http.ResponseWriter, r *http.Request)
}
