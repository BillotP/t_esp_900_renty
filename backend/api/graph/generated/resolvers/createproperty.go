package resolvers

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"io"
	"os"
	"strconv"
)

func saveFile(url string, file *graphql.Upload) (string, error) {
	var (
		path string
		fo   *os.File

		err error
	)

	if path, err = os.Getwd(); err != nil {
		return "", err
	}
	path = path + "/data/"

	fullUrl := url + "/" + file.Filename
	if _, err = os.Stat(path + url); os.IsNotExist(err) {
		if err = os.MkdirAll(path+url, os.ModePerm); err != nil {
			return "", err
		}
	}

	if fo, err = os.Create(path + fullUrl); err != nil {
		return "", err
	}

	buf := make([]byte, file.Size)
	for {
		// read a chunk
		n, err := file.File.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			return "", err
		}
	}

	return fullUrl, nil
}

func (r *MutationResolver) CreateProperty(ctx context.Context, input *models.PropertyInput) (*models.Property, error) {
	var (
		estateAgent models.EstateAgent

		property *models.Property
		url      string
		err      error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)
	estateAgent = models.EstateAgent{User: &models.User{Username: username}}

	if err = r.DB.Joins("User").Joins("Company").Where("username = ?", estateAgent.User.Username).First(&estateAgent).Error; err != nil {
		lib.LogError("mutation/CreateProperty", err.Error())
		return nil, err
	}

	property = &models.Property{
		Area:             input.Area,
		Country:          input.Country,
		CityName:         input.CityName,
		Address:          input.Address,
		PostalCode:       input.PostalCode,
		Type:             input.Type,
		Company:          estateAgent.Company,
		CompanyID:        estateAgent.CompanyID,
		Description:      input.Description,
		Rooms:            input.Rooms,
		Bedrooms:         input.Bedrooms,
		Furnished:        input.Furnished,
		ConstructionDate: input.ConstructionDate,
		EnergyRating:     input.EnergyRating,
		RentAmount:       input.RentAmount,
		ChargesAmount:    input.ChargesAmount,
	}
	if err = r.DB.Where("address = ?", property.Address).First(&property).Error; err == nil {
		return nil, fmt.Errorf("there is already a property at this address")
	}
	if err = r.DB.Create(&property).Error; err != nil {
		lib.LogError("mutation/Register/Property", err.Error())
		return nil, err
	}

	for _, badge := range input.Badges {
		lib.LogError("mutation/CreateProperty", string(*badge))
		if err = r.DB.Model(&property).Association("Badges").Append(&models.Badge{
			Type: badge,
		}); err != nil {
			lib.LogError("mutation/CreateProperty", err.Error())
			return nil, err
		}
	}

	for _, photo := range input.Photos {
		if url, err = saveFile("/photos/"+strconv.FormatInt(*property.ID, 10), photo); err != nil {
			lib.LogError("mutation/CreateProperty", err.Error())
			return nil, err
		}

		if err = r.DB.Model(&property).Association("Photos").Append(&models.Asset{
			URL:  url,
			Type: photo.ContentType,
		}); err != nil {
			lib.LogError("mutation/CreateProperty", err.Error())
			return nil, err
		}
	}

	if input.Model != nil {
		if url, err = saveFile("/models3d/"+strconv.FormatInt(*property.ID, 10), input.Model); err != nil {
			lib.LogError("mutation/CreateProperty", err.Error())
			return nil, err
		}
		if err = r.DB.Model(&property).Update("Model", &models.Asset{
			URL:  url,
			Type: input.Model.ContentType,
		}).Error; err != nil {
			lib.LogError("mutation/CreateProperty", err.Error())
			return nil, err
		}
	}

	return property, nil
}
