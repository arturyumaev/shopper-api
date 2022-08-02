package queries

import (
	"fmt"
)

var deleteUserQuery = `delete from %s.%s where id=$1 returning id`

var DeleteUser = fmt.Sprintf(deleteUserQuery, SCHEMA, TABLE_USERS_NAME)
