package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "100HW/db/sqlc"
)

func CreateAuthor(ctx *gin.Context) {
    var req struct {
        Name string `json:"name"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    author, err := db.CreateAuthor(ctx, req.Name)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, author)
}

func ListAuthors(ctx *gin.Context) {
    authors, err := db.ListAuthors(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, authors)
}

func AssignAuthorToTask(ctx *gin.Context) {
    taskID := ctx.Param("task_id")
    authorID := ctx.Param("author_id")

    err := db.AssignAuthorToTask(ctx, sqlc.AssignAuthorToTaskParams{
        TaskID:  taskID,
        AuthorID: authorID,
    })
    
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusOK)
}
