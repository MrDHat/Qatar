package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/martini-contrib/cors"
)

func main() {
    m := martini.Classic()

    m.Use(render.Renderer(render.Options{
      Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
      IndentJSON: true, // Output human readable JSON
    }))
    m.Use(cors.Allow(&cors.Options{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"GET", "POST", "PUT", "PATCH"},
        ExposeHeaders: []string{"Content-Length"},
    }))

    m.Get("/", func(r render.Render) {
        r.JSON(200, map[string]interface{}{"code": "200","message": "status ok"})
    })

    m.Run()
}