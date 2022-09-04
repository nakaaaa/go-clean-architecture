package presenter

import "github.com/nakaaaa/go-clean-architecture/go/internal/usecase"

func FromUserGetListOutputPort(output *usecase.UserGetListOutputPort) (interface{}, error) {
	return output.List, nil
}
