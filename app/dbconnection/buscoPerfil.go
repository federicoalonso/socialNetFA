package dbconnection

import (
	"context"
	"fmt"
	"time"

	"github.com/federicoalonso/socialnetfa/app/models"
	"github.com/federicoalonso/socialnetfa/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BuscoPerfil busca un perfil en la bd
func BuscoPerfil(ID string) (models.Usuario, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoCN.Database("socialNetFA")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no Encontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
