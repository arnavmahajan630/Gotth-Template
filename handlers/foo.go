package handler

import (
	"net/http"

	foo "github.com/arnavmahajan630/GOTTH-TEMPLATE/views"
)

func Foo(w http.ResponseWriter, r * http.Request) error {
	return Render(w, r, foo.Index())
}