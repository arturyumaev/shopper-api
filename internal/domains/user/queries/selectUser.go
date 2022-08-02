package queries

import (
	"fmt"
)

var selectUserQuery = `select id, username, password, email from %s.%s where id=$1`

var SelectUser = fmt.Sprintf(selectUserQuery, SCHEMA, TABLE_USERS_NAME)
