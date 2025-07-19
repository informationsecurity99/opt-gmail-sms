package service

import (
	"test/pkg/logger"
	"test/pkg/mailer"
	"test/storage"
)

type IServiceManager interface {
	User() UserService
	Otp() OtpService
	Role() RoleService
	SysUser() SysUserService
	Mailer() MailerService 
}

type service struct {
	userService    UserService
	otpService     OtpService
	roleService    RoleService
	sysUserService SysUserService
	mailer         MailerService 
}

func New(storage storage.IStorage, log logger.ILogger, mailerCore *mailer.Mailer, redis storage.IRedisStorage) IServiceManager {
	return &service{
		userService:    NewUserService(storage, log),
		otpService:     NewOtpService(storage, log, mailerCore, redis), 
		roleService:    NewRoleService(storage, log),
		sysUserService: NewSysUserService(storage, log),
		mailer:         NewMailerService(mailerCore),
	}
}



func (s *service) User() UserService {
	return s.userService
}

func (s *service) Otp() OtpService {
	return s.otpService
}

func (s *service) Role() RoleService {
	return s.roleService
}

func (s *service) SysUser() SysUserService {
	return s.sysUserService
}

func (s *service) Mailer() MailerService {
	return s.mailer
}
