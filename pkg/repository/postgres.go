package repository

import (
	"context"
	"fmt"
	"gRPC_cutter/pkg/model"
	"gRPC_cutter/pkg/utils"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type database struct {
	conn *pgxpool.Pool
}

func NewDatabaseRep(conn *pgxpool.Pool) (*database, error) {
	return &database{conn: conn}, nil
}

func (r *database) AddModel(ctx context.Context, url string) (*model.Model, error) {
	q := `
	INSERT INTO url
	(longurl, shorturl)
	values
	($1, $2)
	returning id
`

	m, err := model.NewModel(0, url, utils.Encode())
	if err != nil {
		return nil, err
	}

	if err := r.conn.QueryRow(ctx, q, m.Longurl, m.Shorturl).Scan(&m.ID); err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: " +
				pgErr.Message +
				", Detail: " +
				pgErr.Detail +
				", Where: " +
				pgErr.Where +
				", SQLState: " +
				pgErr.SQLState()))
			return nil, newErr
		} else {
			return nil, err
		}
	}
	return m, nil
}
func (r *database) GetModel(ctx context.Context, shortURL string) (string, error) {

	var res string

	q := `
	SELECT longurl
	FROM url
	WHERE shorturl = $1
`
	if err := r.conn.QueryRow(ctx, q, shortURL).Scan(&res); err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: " +
				pgErr.Message +
				", Detail: " +
				pgErr.Detail +
				", Where: " +
				pgErr.Where +
				", SQLState: " +
				pgErr.SQLState()))
			return "", newErr
		} else {
			return "", err
		}
	}
	return res, nil
}
