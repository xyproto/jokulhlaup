package main

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/unrolled/render"
	"github.com/xyproto/fizz"
)

const (
	Version = "0.1"
)

func main() {
	fmt.Println("jøkulhlaup ", Version)

	r := render.New(render.Options{})

	m := martini.Classic()

	fizz := fizz.New()

	// Dashboard
	m.Get("/", func(w http.ResponseWriter, req *http.Request) {
		data := map[string]string{
			"title": "Jøkulhlaup",
			"imgsrc": "img/jøkulhlaup.png",
			"width": "1440",
			"height": "900",
		}

		// !! Reload template !!
		//r = render.New(render.Options{})

		// Render the specified templates/.tmpl file as HTML and return
		r.HTML(w, http.StatusOK, "black", data)
	})

	// Activate the permission middleware
	m.Use(fizz.All())

	// Share the files in static
	m.Use(martini.Static("static"))

	m.Run() // port 3000 by default
}
