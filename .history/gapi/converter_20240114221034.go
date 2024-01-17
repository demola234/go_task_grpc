package gapi

import (
	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
)

func convertTasks(tasks []db.Tasks) []*pb.Task {
	var task []*pb.Task
	for _, t := range tasks {
		task = append(task, &pb.Task{
			TaskName:        t.TaskName,
			TaskDescription: t.TaskName,
			UserId:          t.UserID,
			CategoryId:      t.CategoryID,
			TaskDeadline:    t.DueDate.Format("2006-01-02 15:04:05"),
			TaskStatus:      t.TaskStatus,
			TaskId:          t.ID,
		})
	}
	return task
}
