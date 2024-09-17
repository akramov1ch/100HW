package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.POST("/tasks", CreateTask)
    router.GET("/tasks", ListTasks)
    router.GET("/tasks/:id", GetTask)
    router.PUT("/tasks/:id", UpdateTask)
    router.DELETE("/tasks/:id", DeleteTask)

    router.POST("/authors", CreateAuthor)
    router.GET("/authors", ListAuthors)
    router.POST("/tasks/:task_id/authors/:author_id", AssignAuthorToTask)

    return router
}
