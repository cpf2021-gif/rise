package repo

import (
	"context"
	"strconv"

	"github.com/cpf2021-gif/rise/common/database"
	"github.com/cpf2021-gif/rise/internal/conf"

	"github.com/cpf2021-gif/rise/internal/repo/db"
)

type Repo struct {
	db *db.Queries
}

func NewRepo(conf *conf.Conf) *Repo {
	return &Repo{
		db: db.New(database.NewDB(conf.DB)),
	}
}

func (r *Repo) GetUserByID(ctx context.Context, uid int64) (*db.GetUserByIDRow, error) {
	return r.db.GetUserByID(ctx, int32(uid))
}

func (r *Repo) CreateUser(ctx context.Context, userName string) (int64, error) {
	return r.db.InsertUser(ctx, userName)
}

func (r *Repo) GetUserByIDs(ctx context.Context, uids []string) ([]*db.GetUserByIDsRow, error) {
	uidsInt := make([]int32, 0, len(uids))
	for _, uid := range uids {
		id, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			return nil, err
		}
		uid := int32(id)
		uidsInt = append(uidsInt, uid)
	}

	return r.db.GetUserByIDs(ctx, uidsInt)
}
