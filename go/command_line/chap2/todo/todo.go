package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var outOfBoundsError = errors.New("out of bounds")

type item struct {
	Task        string
	Completed   bool
	CompletedOn time.Time
	Created     time.Time
	Updated     time.Time
}

type List []item

func NewList() *List {
	return &List{}
}

func (l *List) Count() int {
	return len(*l)
}

func (l *List) Add(task string) {
	itm := item{
		Task:      task,
		Completed: false,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	*l = append(*l, itm)
}

func (l *List) Remove(idx int) error {
	if idx < 0 || idx >= len(*l) {
		return outOfBoundsError
	}
	*l = append((*l)[:idx], (*l)[idx+1:]...)
	return nil
}

func (l *List) Get(idx int) (*item, error) {
	if idx < 0 || idx >= len(*l) {
		return nil, outOfBoundsError
	}

	return &(*l)[idx], nil
}

func (l *List) Complete(idx int) error {
	if idx < 0 || idx >= len(*l) {
		return outOfBoundsError
	}
	(*l)[idx].Completed = true
	(*l)[idx].CompletedOn = time.Now()
	return nil
}

func (l *List) Save(filename string) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (l *List) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Completed {
			prefix = "X "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
