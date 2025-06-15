package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Step struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	IsCompleted bool               `json:"is_completed" bson:"is_completed"` // Completed indicates if the step is done
}

type Task struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Deadline time.Time             `json:"deadline" bson:"deadline"`
	Steps    []Step             `json:"steps" bson:"steps"` // Steps is an array of Step objects
	UserID    string 			`json:"user_id" bson:"user_id"` // Owner is the ID of the user who created the task
}
