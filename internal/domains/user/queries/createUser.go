package queries

import (
	"fmt"
)

var createUserQuery = `
insert into %s.%s (id, username, email, password)
	values ($1, $2, $3, $4)
`

var CreateUser = fmt.Sprintf(createUserQuery, SCHEMA, TABLE_USERS_NAME)
