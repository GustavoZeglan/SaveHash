package converter

import (
	"github.com/GustavoZeglan/SaveHash/core/user/domain"
	"github.com/GustavoZeglan/SaveHash/core/user/entity"
)

func ConvertEntityToDomain(userEntity *entity.UserEntity) domain.UserDomainInterface {
	d := domain.NewUserDomain(userEntity.ID, userEntity.Username, userEntity.Email, userEntity.Password)
	return d
}
