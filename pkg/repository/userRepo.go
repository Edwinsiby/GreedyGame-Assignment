package repository

import (
	"errors"
	"fmt"
	"greedy/pkg/domain"
	"log"
	"time"
)

var (
	err error
)

type Repository struct {
	data map[string]domain.KeyValue
	que  domain.Que
}

type RepositoryMethods interface {
	FindKey(string) bool
	Set(string, domain.KeyValue) error
	Get(string) (*domain.KeyValue, error)
	QPush(string, ...string) error
	QPop(string) (string, error)
	BQPop(string, int) (string, error)
}

func NewRepoLayer(config *domain.Config) RepositoryMethods {
	return &Repository{
		data: make(map[string]domain.KeyValue),
		que: domain.Que{
			Data: make(map[string][]string),
		},
	}
}

func (r *Repository) FindKey(key string) bool {
	_, exists := r.data[key]
	return exists
}

func (r *Repository) Set(key string, value domain.KeyValue) error {
	r.data[key] = value
	log.Println(r.data)
	return nil
}

func (r *Repository) Get(key string) (*domain.KeyValue, error) {
	value, exists := r.data[key]
	if exists != true {
		return nil, errors.New("Key not found")
	}
	return &value, nil
}

func (r *Repository) QPush(key string, values ...string) error {
	r.que.M.Lock()
	defer r.que.M.Unlock()
	if _, ok := r.que.Data[key]; !ok {
		r.que.Data[key] = []string{}
	}
	r.que.Data[key] = append(r.que.Data[key], values...)
	fmt.Println(r.que.Data)
	return nil
}

func (r *Repository) QPop(key string) (string, error) {
	r.que.M.Lock()
	defer r.que.M.Unlock()

	queue, ok := r.que.Data[key]
	if !ok || len(queue) == 0 {
		return "", errors.New("queue is empty")
	}

	value := queue[len(queue)-1]
	r.que.Data[key] = queue[:len(queue)-1]

	return value, nil
}
func (r *Repository) BQPop(key string, timeout int) (string, error) {
	queue, ok := r.que.Data[key]
	if !ok {
		return "", errors.New("queue is empty")
	}

	if len(queue) == 0 {
		select {
		case <-time.After(time.Duration(timeout) * time.Second):
			r.que.M.Lock()
			defer r.que.M.Unlock()
			if len(queue) == 0 {
				return "", errors.New("TimeOut")
			}
		}
	}

	resultCh := make(chan string, 1)
	go func() {
		r.que.M.Lock()
		defer r.que.M.Unlock()
		if len(queue) == 0 {
			resultCh <- "null"
		} else {
			value := queue[len(queue)-1]
			r.que.Data[key] = queue[:len(queue)-1]
			resultCh <- value
		}
	}()

	select {
	case value := <-resultCh:
		return value, nil
	case <-time.After(time.Duration(timeout) * time.Second):
		return "null", errors.New("TimeOut")
	}
}
