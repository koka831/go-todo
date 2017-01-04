package model

import (
	"fmt"
)

// Task[]todo <-> csv <-> file

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

func NewTask(task string, deadline int) (*Task, error) {
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
		return nil
	}

	for i, t := range manager.tasks {
		if t.id == task.id {
			manager.tasks[i] = clone(task)
			return nil
		}
	}
	return fmt.Errorf("may conflict happens in configfile")
	
}

func (manager *TaskManager) Delete(id int) error {
	if len(manager.tasks) < id {
		return fmt.Errorf("cannot find the task")
	}
	// TODO: write
	return nil
}

func (manager *TaskManager) List() []*Task {
	return manager.tasks
}

// reconstruct slices at save
func clone(task *Task) *Task {
	c := *task
	return &c
}



