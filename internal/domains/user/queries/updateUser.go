package queries

import (
	"fmt"
)

var updateUserQuery = `
update %s.%s set id=$2, username=$3, email=$4, password=$5 where id=$1
`

var UpdateUser = fmt.Sprintf(updateUserQuery, SCHEMA, TABLE_USERS_NAME)
