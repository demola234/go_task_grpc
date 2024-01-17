// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
)

type Querier interface {
	CreateTasks(ctx context.Context, arg CreateTasksParams) (Tasks, error)
	DeleteTasks(ctx context.Context, id int64) error
	GetTask(ctx context.Context, id int64) (Tasks, error)
	ListTasks(ctx context.Context) ([]Tasks, error)
	ListTasksByCategory(ctx context.Context, arg ListTasksByCategoryParams) ([]Tasks, error)
	UpdateTasks(ctx context.Context, arg UpdateTasksParams) error
}

var _ Querier = (*Queries)(nil)