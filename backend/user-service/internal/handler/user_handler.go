func RegisterRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
    users := router.Group("/users")
    {
        users.GET("/me", userHandler.GetCurrentUser)
        users.GET("/:id", userHandler.GetUserByID)
    }
}