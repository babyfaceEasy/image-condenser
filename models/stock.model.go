package models

import "database/sql"

type Stock struct {
	ID            int
	PicturePath   sql.NullString
	ThumbnailPath sql.NullString
}
