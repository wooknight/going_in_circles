package list

import (
	"encoding/json"
	"os"
	"time"
)

type Todo struct {
	Title       string
	Done        bool
	CreatedAt   string
	CompletedAt string
}

type List []Todo

func (l *List) String() string {
	bytes, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(bytes)
}
func (l *List) Add(t Todo) {
	*l = append(*l, t)
}
func (l *List) DeleteByIdx(idx int) {
	*l = append((*l)[:idx], (*l)[idx+1:]...)
}

func (l *List) DeleteByTitle(cmp string) {
	for idx, todo := range *l {

		if todo.Title == cmp {
			if idx == len(*l)-1 {
				*l = (*l)[:idx]
			} else {
				*l = append((*l)[:idx], (*l)[idx+1:]...)
			}
		}
	}
}
func (l *List) CompleteByIdx(idx int) {
	(*l)[idx].Done = true
	(*l)[idx].CompletedAt = time.Now().Format(time.RFC3339)
}

func (l *List) CompleteByTitle(cmp string) {
	for idx, todo := range *l {
		if todo.Title == cmp {
			(*l)[idx].Done = true
			(*l)[idx].CompletedAt = time.Now().Format(time.RFC3339)
		}
	}

}

func (l *List) Save(filename string) error {
	bytes, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bytes, 0644)

}
func (l *List) Load(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, l)
	if err != nil {
		return err
	}
	return nil
}
func (l *List) List() string {
	bytes, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(bytes)
}
