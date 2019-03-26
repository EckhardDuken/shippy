package go_micro_srv_user

import (
	uuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	//uuid "github.com/satori/go.uuid"
)

// BeforeCreate method
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	luuid := uuid.New()
	//Sluuid, err := uuid.NewV4()
	// if err != nil {
	// 	return err
	// }
	return scope.SetColumn("Id", luuid.String())
}
