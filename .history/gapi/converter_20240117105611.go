package gapi

import (
	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
)

func convertTasks(tasks []db.Tasks) []*pb.Task {
	var task []*pb.Task
	for _, t := range tasks {
		task = append(task, &pb.Task{
			TaskId:          t.,
			TaskName:        t.TaskName,
			TaskDescription: t.TaskDescription,
			TaskStatus:      t.TaskStatus,
			TaskDeadline:    t.DueDate.String(),
		})
	}
	return task
}
