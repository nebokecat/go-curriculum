package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"bst-tech/program/graph"
	"bst-tech/program/infra/boiler"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// GetTask is the resolver for the getTask field.
func (r *queryResolver) GetTask(ctx context.Context, taskID int) (*graph.GetTaskOutput, error) {
	db, err := sql.Open("postgres", "host=postgres user=tech password=secret dbname=tech sslmode=disable")
	if err != nil {
		return nil, err
	}

	task, err := boiler.Tasks(boiler.TaskWhere.ID.EQ(int64(taskID))).One(ctx, db)
	if err != nil {
		return nil, err
	}

	return &graph.GetTaskOutput{
		ID:          int(task.ID),
		Name:        task.Name,
		Description: task.Description.Ptr(),
		Priority:    task.Priority.Ptr(),
	}, nil
}