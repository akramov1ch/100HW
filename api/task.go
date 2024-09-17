package api

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
    "100HW/db"
)

func CreateTask(ctx *gin.Context) {
    var req struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Status      string `json:"status"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task, err := db.CreateTask(ctx, db.CreateTaskParams{
        Title:       req.Title,
        Description: req.Description,
        Status:      req.Status,
    })

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, task)
}

func GetTask(ctx *gin.Context) {
    id := ctx.Param("id")

    task, err := db.GetTask(ctx, id)
    if err == sql.ErrNoRows {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    } else if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, task)
}

func ListTasks(ctx *gin.Context) {
    tasks, err := db.ListTasks(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, tasks)
}

func UpdateTask(ctx *gin.Context) {
    id := ctx.Param("id")
    var req struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Status      string `json:"status"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task, err := db.UpdateTask(ctx, db.UpdateTaskParams{
        ID:          id,
        Title:       req.Title,
        Description: req.Description,
        Status:      req.Status,
    })

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context) {
    id := ctx.Param("id")

    err := db.DeleteTask(ctx, id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusOK)
}
