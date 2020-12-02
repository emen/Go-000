// Dao Layer
// Sample MySQL

package mysql

import (
	"context"
	"database/sql"

	"github.com/emen/Go-000/Week02/model"
	"github.com/pkg/errors"
)

type MySQLClient struct {
	conn sql.Conn
}

//
// Example of handling errors
//
func (c *MySQLClient) Get(id int) (*model.User, error) {
	_, err := c.conn.QueryContext(context.Background(), `select id, first_name, last_name from users`)

	if err == sql.ErrNoRows {
		return nil, errors.Wrap(model.ErrNoRecord, err.Error())
	}

	return &model.User{}, nil
}

func (c *MySQLClient) Create(user *model.User) (int, error) {
	return 0, nil
}
