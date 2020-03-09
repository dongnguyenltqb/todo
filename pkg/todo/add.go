package todo

import (
	"errors"
	"todo/pkg/infra"
)

func Add(list string, value string) (err error) {
	if value == "" || list == "" {
		err = errors.New("Not empty")
		return
	}
	err = infra.GetRedis().RPush(list, value).Err()
	return
}
