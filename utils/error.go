package utils

import (
	"fmt"
	"net/http"
)

type Error struct { }

func (e Error) NotFoundError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "404! Not found")
}
