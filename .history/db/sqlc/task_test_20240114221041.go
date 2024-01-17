package db

import (
	"context"
	"testing"
	"time"

	"github.com/demola234/todogrpc/utils"
	"github.com/stretchr/testify/require"
)

func createTask(t *testing.T) Tasks {


	arg := CreateTasksParams{
		TaskName:   utils.RandomOwner(),
		UserID:     1,
	
		DueDate:    time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC),
		CreatedAt:  time.Now(),
		TaskStatus: utils.PENDING,
	}

	task, err := testQueries.CreateTasks(context.Background(), arg)
	require.NoError(t, err)
	return task
}

func TestCreateTasks(t *testing.T) {
	task := createTask(t)

	require.NotEmpty(t, task)
	require.Equal(t, task.TaskName, task.TaskName)
	require.Equal(t, task.DueDate, task.DueDate)
	require.Equal(t, task.CreatedAt, task.CreatedAt)
}

func TestGetTasks(t *testing.T) {
	task1 := createTask(t)
	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
	require.Equal(t, task1.TaskName, task2.TaskName)

	require.Equal(t, task1.DueDate, task2.DueDate)
	require.Equal(t, task1.CreatedAt, task2.CreatedAt)
	require.Equal(t, task1.TaskStatus, task2.TaskStatus)
}

func TestDeleteTasks(t *testing.T) {
	task1 := createTask(t)
	err := testQueries.DeleteTasks(context.Background(), task1.ID)
	require.NoError(t, err)
	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.Error(t, err)
	require.Empty(t, task2)
}

func TestUpdateTasks(t *testing.T) {
	task1 := createTask(t)
	arg := UpdateTasksParams{
		ID:         task1.ID,
		TaskName:   "test",
		DueDate:    time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC),
		TaskStatus: utils.PENDING,
	}

	err := testQueries.UpdateTasks(context.Background(), arg)
	require.NoError(t, err)
}

func TestListTasks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTask := createTask(t)

		arg := ListTasksParams{
			UserID: createTask.UserID,
			Limit:  5,
			Offset: 5,
		}

		tasks, err := testQueries.ListTasks(context.Background(), arg)
		require.NoError(t, err)

		for _, task := range tasks {
			require.NotEmpty(t, task)
			require.Equal(t, createTask.UserID, task.UserID)

		}

	}

}

func TestListTasksByCategory(t *testing.T) {

	for i := 0; i < 10; i++ {
		createTask(t)
	}

	arg := ListTasksByCategoryParams{
		UserID: 1,
		Limit:  5,
		Offset: 5,
	}

	tasks, err := testQueries.ListTasksByCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)

	for _, task := range tasks {
		require.NotEmpty(t, task)
	}
}
