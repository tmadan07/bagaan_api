package types

import (
	"database/sql"
	"encoding/json"
)

// First create a type alias
type NullString sql.NullString

func (u NullString) MarshalJSON() ([]byte, error) {
	if u.Valid {
		return json.Marshal(u.String)
	}
	return nil, nil
}
