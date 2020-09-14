package dbconnection

import (
	"context"
	"time"

	"github.com/federicoalonso/socialnetfa/app/models"
	"github.com/federicoalonso/socialnetfa/config"
)

// InsertoRelacion crea la rel
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoCN.Database("socialNetFA")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
