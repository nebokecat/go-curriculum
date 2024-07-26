package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"bst-tech/program/graph"
	"bst-tech/program/infra/boiler"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input graph.UpdateTaskInput) (*graph.UpdateTaskOutput, error) {
	db, err := sql.Open("postgres", "host=postgres user=tech password=secret dbname=tech sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Find a pilot and update his name
	task, err := boiler.FindTask(ctx, db, int64(input.ID))
	if err != nil {
		return nil, err
	}

	task.Name = *input.Name

	rowsAff, err := task.Update(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	task, err = boiler.FindTask(ctx, db, int64(rowsAff))
	if err != nil {
		return nil, err
	}

	return &graph.UpdateTaskOutput{
		ID:          int(task.ID),
		Name:        &task.Name,
		Description: task.Description.Ptr(),
	}, nil
}
