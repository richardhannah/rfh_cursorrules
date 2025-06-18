package db

import (
	"totmapi/internal/dto"
	"totmapi/internal/logger"
	"totmapi/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	dbContext IDbContext
}

func NewUserRepository(ctx IDbContext) *UserRepository {
	return &UserRepository{dbContext: ctx}
}

func (u UserRepository) SelectAll() []models.Users {
	var users []models.Users

	predicate := func(sb sq.SelectBuilder) sq.SelectBuilder {
		return sb.OrderBy("username")
	}

	if err := u.dbContext.Select(&users, predicate); err != nil {
		logger.Error("Error selecting all users", err)
	}
	return users
}

func (u UserRepository) SelectById(id string) []models.Users {
	var users []models.Users

	predicate := func(sb sq.SelectBuilder) sq.SelectBuilder {
		return sb.
			Where(sq.Eq{"id": id})
	}

	if err := u.dbContext.Select(&users, predicate); err != nil {
		logger.Error("Error selecting user by id", err,
			logger.String("user_id", id),
		)
	}
	return users
}

func (u UserRepository) SelectByUsername(username string) []models.Users {
	var users []models.Users

	predicate := func(sb sq.SelectBuilder) sq.SelectBuilder {
		return sb.
			Where(sq.Eq{"username": username})
	}

	if err := u.dbContext.Select(&users, predicate); err != nil {
		logger.Error("Error selecting user by username", err,
			logger.String("username", username),
		)
	}
	return users
}

func (u UserRepository) Insert(user dto.UserDTO) {
	predicate := func(sb sq.InsertBuilder) sq.InsertBuilder {
		return sb.
			Columns("id", "username", "password", "salt", "ipaddress", "enabled", "role").
			Values(user.ID, user.Username, user.Password, user.Salt, user.Ipaddress, user.Enabled, user.Role)
	}

	u.dbContext.Insert(&[]models.Users{}, predicate)
}

func (u UserRepository) Update(user dto.UserDTO) {
	predicate := func(ub sq.UpdateBuilder) sq.UpdateBuilder {
		return ub.
			Set("username", user.Username).
			Set("password", user.Password).
			Set("salt", user.Salt).
			Set("ipaddress", user.Ipaddress).
			Set("enabled", user.Enabled).
			Set("role", user.Role).
			Where(
				sq.Eq{"id": user.ID},
			)
	}

	u.dbContext.Update(&[]models.Users{}, predicate)
}

func (u UserRepository) Delete(user dto.UserDTO) {
	predicate := func(db sq.DeleteBuilder) sq.DeleteBuilder {
		return db.Where(
			sq.Eq{"id": user.ID},
		)
	}

	err := u.dbContext.Delete(&[]models.Users{}, predicate)
	if err != nil {
		logger.Error("Error deleting user", err,
			logger.String("user_id", user.ID),
			logger.String("username", user.Username),
		)
	}
}
