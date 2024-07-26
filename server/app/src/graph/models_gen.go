// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type CreateTaskInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type CreateTaskOutput struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type GetTaskInput struct {
	ID int `json:"id"`
}

type GetTaskOutput struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
