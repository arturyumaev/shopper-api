package queries

import (
	"fmt"
)

var selectUsersQuery = `select * from %s.%s`

var SelectUsers = fmt.Sprintf(selectUsersQuery, SCHEMA, TABLE_USERS_NAME)
