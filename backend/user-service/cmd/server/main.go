package main

import (
    "github.com/gin-gonic/gin"

    "github.com/CapThunder19/Sentinel/backend/user-service/internal/routes"
)

func main() {
    router := gin.Default()

    routes.RegisterRoutes(router)

    router.Run(":8082")
}