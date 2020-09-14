package dbconnection

import (
	"context"
	"time"

	"github.com/federicoalonso/socialnetfa/app/models"
	"github.com/federicoalonso/socialnetfa/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro insertar registros en usuario
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoCN.Database("socialNetFA")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	u.ID = primitive.NewObjectID()

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
