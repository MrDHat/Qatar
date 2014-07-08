package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

func main() {
    m := martini.Classic()

    m.Use(render.Renderer(render.Options{
      Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
      IndentJSON: true, // Output human readable JSON
    }))

    m.Get("/", func(r render.Render) {
        r.JSON(200, map[string]interface{}{"code": "200","message": "status ok"})
    })
    m.Run()
}