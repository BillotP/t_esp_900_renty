package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/v2/models"
)

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var data io.Reader
	var results models.User
	bodyByte, err := json.Marshal(ByID{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	if blob := bytes.NewBuffer(bodyByte); blob != nil {
		data = blob
	}
	res, err := http.Post(GatewayBaseURL+"/getuserbyid", "application/json", data)
	if err != nil {
		fmt.Printf("Error(User): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(User): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *queryResolver) Users(ctx context.Context, query *model.SearchQuery) ([]*models.User, error) {
	var data io.Reader
	var results []*models.User
	if query != nil {
		bodyByte, err := json.Marshal(*query)
		if err != nil {
			return nil, err
		}
		if blob := bytes.NewBuffer(bodyByte); blob != nil {
			data = blob
		}
	}
	res, err := http.Post(GatewayBaseURL+"/getusers", "application/json", data)
	if err != nil {
		fmt.Printf("Error(Users): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(Users): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *queryResolver) Rentoffer(ctx context.Context, id string) (*models.RentOffer, error) {
	var data io.Reader
	var results models.RentOffer
	bodyByte, err := json.Marshal(ByID{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	if blob := bytes.NewBuffer(bodyByte); blob != nil {
		data = blob
	}
	res, err := http.Post(GatewayBaseURL+"/getrentofferbyid", "application/json", data)
	if err != nil {
		fmt.Printf("Error(Rentoffer): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(Rentoffer): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *queryResolver) Rentoffers(ctx context.Context, query *model.SearchQuery) ([]*models.RentOffer, error) {
	var data io.Reader
	var results []*models.RentOffer
	if query != nil {
		bodyByte, err := json.Marshal(*query)
		if err != nil {
			return nil, err
		}
		if blob := bytes.NewBuffer(bodyByte); blob != nil {
			data = blob
		}
	}
	url := GatewayBaseURL + "/getrentoffers"
	if goscrappy.Debug {
		fmt.Printf("Info(Rentoffers): Will do request to [%s]\n", url)
	}
	res, err := http.Post(url, "application/json", data)
	if err != nil {
		fmt.Printf("Error(Rentoffers): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(Rentoffers): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *queryResolver) Offeror(ctx context.Context, id string) (*models.Offeror, error) {
	var data io.Reader
	var results models.Offeror
	bodyByte, err := json.Marshal(ByID{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	if blob := bytes.NewBuffer(bodyByte); blob != nil {
		data = blob
	}
	res, err := http.Post(GatewayBaseURL+"/getofferorbyid", "application/json", data)
	if err != nil {
		fmt.Printf("Error(Offerror): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(Offerror): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *queryResolver) Offerors(ctx context.Context, query *model.SearchQuery) ([]*models.Offeror, error) {
	var data io.Reader
	var results []*models.Offeror
	if query != nil {
		bodyByte, err := json.Marshal(*query)
		if err != nil {
			return nil, err
		}
		if blob := bytes.NewBuffer(bodyByte); blob != nil {
			data = blob
		}
	}
	res, err := http.Post(GatewayBaseURL+"/getofferors", "application/json", data)
	if err != nil {
		fmt.Printf("Error(Offerors): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(Offerrors): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
