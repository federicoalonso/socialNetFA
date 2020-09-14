package models

// Tweet el tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
