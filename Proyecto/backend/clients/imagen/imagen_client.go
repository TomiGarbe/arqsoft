package clients

import (
	"backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"errors"
	e "backend/utils/errors"
)

var Db *gorm.DB

func GetImagenById(id int) model.Imagen {
	var imagen model.Imagen

	Db.Where("id = ?", id).Preload("Hotel").First(&imagen)
	log.Debug("Image: ", imagen)

	return imagen
}

func InsertImageByHotelId(imagen model.Imagen) model.Imagen {
	result := Db.Create(&imagen)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Image Created: ", imagen.ID)
	return imagen
}

func GetImagenesByHotelId(hotelID int) model.Imagenes {
	var imagenes model.Imagenes

	Db.Where("hotel_id = ?", hotelID).Preload("Hotel").Find(&imagenes)
	log.Debug("Imagenes: ", imagenes)

	return imagenes
}

func DeleteImagenById(imagenID int) e.ApiError {

	err := Db.Delete(&model.Imagen{}, imagenID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Image not found")
		}
		return e.NewBadRequestApiError("Failed to delete Image")
	}

	return nil
}