// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package login

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeLoginService() *LoginService {
	db := NewDbHandler2()
	loginLoginDal := NewloginDal(db)
	loginLoginMgr := NewloginMgr(loginLoginDal)
	loginService := NewLoginService(loginLoginMgr)
	return loginService
}
