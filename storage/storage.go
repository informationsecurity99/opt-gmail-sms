package storage

import (
	"context"
	"time"

	"test/api/models"
)


type IStorage interface {
	User() IUserStorage
	OTP() IOTPStorage
	Role() IRoleStorage
	Sysuser() ISysuserStorage
	Redis() IRedisStorage
	Close()
}


type IUserStorage interface {
	Create(ctx context.Context, req models.CreateUser) (string, error)
	GetForLoginByEmail(ctx context.Context, email string) (models.LoginUser, error)
}


type IOTPStorage interface {
	Create(ctx context.Context, email string, code string, expiresAt time.Time) (string, error)
	GetUnconfirmedByID(ctx context.Context, id string) (email string, code string, expiresAt time.Time, err error)
	UpdateStatusToConfirmed(ctx context.Context, id string) error
	GetByIDAndEmail(ctx context.Context, id string, email string) (bool, error)
}


type IRoleStorage interface {
	Create(ctx context.Context, name string, createdBy string) (string, error)
	Update(ctx context.Context, id, name string) error
	GetAll(ctx context.Context) ([]models.Role, error)
	Exists(ctx context.Context, id string) (bool, error)
}


type ISysuserStorage interface {
	GetByPhone(ctx context.Context, phone string) (id, hashedPassword string, status string, err error)
	Create(ctx context.Context, name, phone, hashedPassword, createdBy string) (string, error)
	AssignRoles(ctx context.Context, sysuserID string, roleIDs []string) error
}


type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) (string, error) 
}
