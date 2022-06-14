package utils

import "database/sql"

// StringToNullString ...
func StringToNullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}
