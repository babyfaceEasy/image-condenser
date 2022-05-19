package models

import "database/sql"

type Stock struct {
	ID            int
	PicturePath   string
	ThumbnailPath sql.NullString
}
