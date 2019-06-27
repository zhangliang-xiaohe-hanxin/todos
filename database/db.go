package db 

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"context"
)

func GetSession(c context.Context) (*sql.DB, error) {
	session, ok := c.Value("session").(*sql.DB)
	if !ok {
		return nil, errors.New("cannot get session")
	}
	return session, nil
}
