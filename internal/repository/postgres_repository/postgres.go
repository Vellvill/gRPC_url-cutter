package postgres_repository

import (
	"context"
	"fmt"
	"gRPC_cutter/internal/model"
	"gRPC_cutter/internal/utils"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type database struct {
	conn *pgxpool.Pool
}

func NewDatabaseRep(conn *pgxpool.Pool) (*database, error) {
	return &database{conn: conn}, nil
}

//AddModel создаёт новую запись в таблице. Если код ошибки 23505 (создание записи с уже существующим LongUrl),
//то функция возвращает модель с уже имеющимся ShortURL из базы.
func (r *database) AddModel(ctx context.Context, url string) (*model.Model, error) {
	q := `
	INSERT INTO url
	(longurl, shorturl)
	VALUES
	($1, $2)
	ON CONFLICT (longurl) DO UPDATE SET
	longurl = $1
	RETURNING shorturl
		`

	m, err := model.NewModel(0, url, utils.Encode())
	if err != nil {
		return nil, err
	}

	if err = r.conn.QueryRow(ctx, q, m.Longurl, m.Shorturl).Scan(&m.Shorturl); err != nil {
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
