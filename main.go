package main

import (
    "net/http"

    "gee"
)

func main() {
    engine := gee.New()
    engine.Get("/", func(context *gee.Context) {
        context.String(http.StatusOK, "welcome")
    })
    engine.Get("/url", func(context *gee.Context) {
        context.JSON(http.StatusOK, gee.H{
            "username": context.Query("username"),
        })
    })
    engine.Run(":8080")
}
