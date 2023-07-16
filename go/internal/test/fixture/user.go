package fixture

import (
	"fmt"

	"github.com/nakaaaa/go-clean-architecture/go/internal/domain/model"
	"github.com/samber/lo"
)

func NewUser(user *model.User) *model.User {
	return &model.User{
		UserID: user.UserID,
		Name:   fmt.Sprintf("TestUser%d", user.UserID),
		Age:    lo.Ternary(user.Age != 0, user.Age, 20),
		Gender: lo.Ternary(user.Gender == model.UserGenderMan, user.Gender, model.UserGenderWoman),
	}
}
