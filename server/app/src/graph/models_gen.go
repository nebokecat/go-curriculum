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

type SearchTask struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
}

type SearchTaskInput struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type SearchTasksOutput struct {
	SearchTasks []*SearchTask `json:"SearchTasks,omitempty"`
}

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateTaskInput struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type UpdateTaskOutput struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
