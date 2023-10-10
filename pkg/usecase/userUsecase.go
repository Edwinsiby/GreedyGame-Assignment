package usecase

import (
	"greedy/pkg/domain"
	"greedy/pkg/repository"
)

var (
	err error
)

type Usecase struct {
	repo repository.RepositoryMethods
}

type UsecaseMethods interface {
	Set(string, domain.KeyValue) error
	Get(string) (*domain.KeyValue, error)
	QPush(string, ...string) error
	QPop(string) (string, error)
	BQPop(string, int) (string, error)
}

func NewUsecaseLayer(config *domain.Config, repo repository.RepositoryMethods) UsecaseMethods {
	return Usecase{
		repo: repo,
	}
}

func (u Usecase) Set(key string, keyValue domain.KeyValue) error {
	if keyValue.Expiry != "" {
		keyValue.Expiry = "EX " + keyValue.Expiry
	}
	keyExist := u.repo.FindKey(key)
	if keyValue.Condition == "NX" && keyExist == false {
		if err = u.repo.Set(key, keyValue); err != nil {
			return err
		}
		return nil
	} else if keyValue.Condition == "XX" && keyExist == true {
		if err = u.repo.Set(key, keyValue); err != nil {
			return err
		}
		return nil
	}
	if err = u.repo.Set(key, keyValue); err != nil {
		return err
	}
	return nil
}

func (u Usecase) Get(key string) (*domain.KeyValue, error) {
	value, err := u.repo.Get(key)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (u Usecase) QPush(key string, values ...string) error {
	if err = u.repo.QPush(key, values...); err != nil {
		return err
	}
	return nil
}

func (u Usecase) QPop(key string) (string, error) {
	value, err := u.repo.QPop(key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (u Usecase) BQPop(key string, timeout int) (string, error) {
	value, err := u.repo.BQPop(key, timeout)
	if err != nil {
		return "", err
	}
	return value, nil
}
