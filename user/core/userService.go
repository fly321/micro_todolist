package core

import (
	ct "context"
	"errors"
	"gorm.io/gorm"
	"user/model"
	promos "user/services"
)

func (*UserService) BuildUser(item model.User) *promos.UserModel {
	userModel := &promos.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.Username,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return userModel
}

func (us *UserService) UserLogin(ctx ct.Context, req *promos.UserRequest, resp *promos.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.Db.Where("username = ?", req.UserName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if !user.CheckPassword(req.Password) {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = us.BuildUser(user)
	return nil
}

func (us *UserService) UserRegister(ctx ct.Context, req *promos.UserRequest, resp *promos.UserDetailResponse) error {
	if req.PasswordConfirm != req.Password {
		err := errors.New("两次密码不一致")
		return err
	}
	var user model.User
	resp.Code = 200
	var count int64 = 0
	if err := model.Db.Model(&model.User{}).Where("username=?", req.UserName).Count(&count).Error; err != nil {
		resp.Code = 500
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		if err != nil {
			return err
		}
	}
	user.Username = req.UserName
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.Db.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = us.BuildUser(user)
	return nil
}
