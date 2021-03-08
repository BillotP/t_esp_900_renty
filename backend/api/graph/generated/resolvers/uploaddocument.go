package resolvers

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"io"
	"os"
	"strconv"
)

func (r *MutationResolver) UploadDocument(ctx context.Context, file graphql.Upload, title string) (*bool, error) {
	var (
		tenant models.Tenant

		success bool
		err     error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)
	tenant = models.Tenant{User: &models.User{Username: username}}
	if err = r.DB.Joins("User").Where("username = ?", tenant.User.Username).First(&tenant).Error; err != nil {
		lib.LogError("mutation/UpdateTenantProfile", err.Error())
		return nil, err
	}
	path, err := os.Getwd()
	if err != nil {
		success = false
		return &success, err
	}
	url := "/documents/" + strconv.FormatInt(*(tenant.ID), 10) + "/" + title
	if _, err := os.Stat(path + "/documents/" + strconv.FormatInt(*(tenant.ID), 10)); os.IsNotExist(err) {
		if err = os.MkdirAll(path+"/documents/"+strconv.FormatInt(*(tenant.ID), 10), os.ModePerm); err != nil {
			success = false
			return &success, err
		}
	}
	fo, err := os.Create(path + url)
	if err != nil {
		success = false
		return &success, err
	}

	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := file.File.Read(buf)
		if err != nil && err != io.EOF {
			success = false
			return &success, err
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			success = false
			return &success, err
		}
	}
	if err = r.DB.Model(&tenant).Association("Documents").Append(&models.Asset{
		URL:  url,
		Type: file.ContentType,
	}); err != nil {
		success = false
		return &success, err
	}
	success = true
	return &success, nil
}
