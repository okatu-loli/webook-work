package repository

import (
	"context"
	"github.com/okatu-loli/webook/internal/domain"
	"github.com/okatu-loli/webook/internal/repository/dao"
	"time"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// SELECT * FROM `users` WHERE `email`=?
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		Bio:      u.Bio,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindById(int64) {
	// 先从 cache 里面找
	// 再从 dao 里面找
	// 找到了回写 cache
}

func (r *UserRepository) Update(ctx context.Context, u domain.User) error {
	// Calling the DAO layer's corresponding method to update the database
	return r.dao.Update(ctx, dao.User{
		Id:       u.Id,
		Email:    u.Email,
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		Bio:      u.Bio,
		Password: u.Password,
	})
}

func (r *UserRepository) GetProfile(ctx context.Context, userID int64) (domain.User, error) {
	user, err := r.dao.GetProfile(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	nickname := user.Nickname
	birthday := user.Birthday
	bio := user.Bio

	return domain.User{
		Id:       user.Id,
		Email:    user.Email,
		Nickname: nickname,
		Birthday: birthday,
		Bio:      bio,
		Password: user.Password,
		Ctime:    time.Unix(user.Ctime, 0),
	}, nil
}
