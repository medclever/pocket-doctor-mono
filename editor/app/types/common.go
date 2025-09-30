package types

import "time"

type TimeNow interface {
	Now() time.Time
}

// Ids
type ArticleId string
type ItemId string

// Enums
type ItemType string

const (
	ItemTypeMessage ItemType = "message"
	ItemTypeGallery ItemType = "gallery"
	ItemTypeLink    ItemType = "link"
)

type PlaceType string

const (
	PlaceTypeNecessery PlaceType = "necessary"
	PlaceTypePosible   PlaceType = "possible"
	PlaceTypeMustNot   PlaceType = "must_not"
	PlaceTypeImportant PlaceType = "important"
	PlaceTypeText      PlaceType = "text"
)

var PlaceTypeAvailable = []PlaceType{
	PlaceTypeNecessery,
	PlaceTypePosible,
	PlaceTypeMustNot,
	PlaceTypeImportant,
	PlaceTypeText,
}
