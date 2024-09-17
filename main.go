package main

import (
    "log"
    "100HW/api"
)

func main() {
    router := api.SetupRouter()
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to start server", err)
    }
}
