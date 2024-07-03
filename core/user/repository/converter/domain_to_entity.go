package converter

import (
	"github.com/GustavoZeglan/SaveHash/core/user/domain"
	"github.com/GustavoZeglan/SaveHash/core/user/entity"
)

func ConvertDomainToEntity(di domain.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		ID:       di.GetId(),
		Username: di.GetUsername(),
		Email:    di.GetEmail(),
		Password: di.GetPassword(),
	}
}
