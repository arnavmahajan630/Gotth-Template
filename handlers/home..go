package handler

import (
	"net/http"

	"github.com/arnavmahajan630/GOTTH-TEMPLATE/views/home"
)

func Foo(w http.ResponseWriter, r * http.Request) error {
	return Render(w, r, home.Index())
}