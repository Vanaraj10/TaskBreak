package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Vanaraj10/taskmorph-backend/config"
	"github.com/Vanaraj10/taskmorph-backend/models"
	"github.com/Vanaraj10/taskmorph-backend/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BreakdownTask handles AI task breakdown requests
func BreakdownTask(c *gin.Context) {
	var req struct {
		Task string `json:"task" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task description is required"})
		return
	}

	steps, err := services.AskGemini(req.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate task breakdown"})
		return
	}

	c.JSON(http.StatusOK, steps)
}

// CreateTask creates a new task with AI-generated steps
func CreateTask(c *gin.Context) {
	var req struct {
		Title    string `json:"title" binding:"required"`
		Deadline string `json:"deadline"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse deadline
	var deadline time.Time
	var err error
	if req.Deadline != "" {
		deadline, err = time.Parse("2006-01-02", req.Deadline)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deadline format. Use YYYY-MM-DD"})
			return
		}
	} else {
		deadline = time.Now().AddDate(0, 0, 7) // Default: 7 days from now
	}

	// Generate steps using AI
	steps, err := services.AskGemini(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate steps"})
		return
	}

	// Assign unique IDs to each step
	for i := range steps {
		steps[i].ID = primitive.NewObjectID()
	}

	// Get user ID from email
	userCollection := config.GetCollection("users")
	var user models.User
	err = userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	task := models.Task{
		ID:       primitive.NewObjectID(),
		Title:    req.Title,
		Deadline: deadline,
		Steps:    steps,
		UserID:   user.ID.Hex(),
	}

	collection := config.GetCollection("tasks")
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": task})
}

// GetTasks retrieves all tasks for the authenticated user
func GetTasks(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get user ID from email
	userCollection := config.GetCollection("users")
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	collection := config.GetCollection("tasks")
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": user.ID.Hex()})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode tasks"})
		return
	}

	// Calculate progress for each task
	tasksWithProgress := make([]gin.H, len(tasks))
	for i, task := range tasks {
		completedSteps := 0
		for _, step := range task.Steps {
			if step.IsCompleted {
				completedSteps++
			}
		}

		progress := 0
		if len(task.Steps) > 0 {
			progress = (completedSteps * 100) / len(task.Steps)
		}

		tasksWithProgress[i] = gin.H{
			"id":       task.ID,
			"title":    task.Title,
			"deadline": task.Deadline,
			"progress": progress,
			"steps":    task.Steps,
		}
	}

	c.JSON(http.StatusOK, tasksWithProgress)
}

// GetTask retrieves a specific task by ID
func GetTask(c *gin.Context) {
	taskID := c.Param("id")
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get user ID from email
	userCollection := config.GetCollection("users")
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	collection := config.GetCollection("tasks")
	var task models.Task
	err = collection.FindOne(context.TODO(), bson.M{
		"_id":     objectID,
		"user_id": user.ID.Hex(),
	}).Decode(&task)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Calculate progress
	completedSteps := 0
	for _, step := range task.Steps {
		if step.IsCompleted {
			completedSteps++
		}
	}

	progress := 0
	if len(task.Steps) > 0 {
		progress = (completedSteps * 100) / len(task.Steps)
	}

	response := gin.H{
		"id":       task.ID,
		"title":    task.Title,
		"deadline": task.Deadline,
		"progress": progress,
		"steps":    task.Steps,
	}

	c.JSON(http.StatusOK, response)
}

// CompleteStep marks a step as completed or uncompleted
func CompleteStep(c *gin.Context) {
	taskID := c.Param("taskID")
	stepID := c.Param("stepID")
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get user ID from email
	userCollection := config.GetCollection("users")
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	stepObjectID, err := primitive.ObjectIDFromHex(stepID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid step ID"})
		return
	}

	collection := config.GetCollection("tasks")

	// Update the specific step's completion status
	filter := bson.M{
		"_id":       taskObjectID,
		"user_id":   user.ID.Hex(),
		"steps._id": stepObjectID,
	}
	update := bson.M{
		"$set": bson.M{
			"steps.$.is_completed": true,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update step"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task or step not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Step completed successfully"})
}

// DeleteTask deletes a task
func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get user ID from email
	userCollection := config.GetCollection("users")
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	collection := config.GetCollection("tasks")
	result, err := collection.DeleteOne(context.TODO(), bson.M{
		"_id":     objectID,
		"user_id": user.ID.Hex(),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
