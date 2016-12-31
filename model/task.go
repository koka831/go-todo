package model

import (
	"fmt"
)

// deadline : the day left
// input: int -> save: int->Date
type Task struct {
	id	int
	task	string
	deadline	int
}

// 次のtodo作成時にid++を使うために保持
// TODO: desc in english
type TaskManager struct {
	tasks []*Task
	id	int
}

func Add(task string, deadline int) (*Task, err) {
	if task == "" {
		return nil, fmt.Errorf("cannot create task with empty title")
	}
	return &Task{0, task, deadline}, nil
}

func (manager *TaskManager) Save(task *Task) error {
	if task.id == 0 {
		manager.id ++
		task.id = manager.id
		manager.tasks = append(manager.tasks, task)
	}
}

