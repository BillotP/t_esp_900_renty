package main

import (
	"context"
	"fmt"
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/renty/backend/lib/v2/models"
)

func seedDatas(ctx context.Context, db driver.Database) error {
	var seeds = map[string]interface{}{}
	seeds["labels"] = []models.Label{
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "title1",
				},
			},
			Value:  "Awesome T2",
			Locale: "en",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "title2",
				},
			},
			Value:  "Incredible T3",
			Locale: "en",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "desc1",
				},
			},
			Value:  "Awesome T2 in the center of nowhere, close to all commodities",
			Locale: "en",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "desc2",
				},
			},
			Value:  "Incredible T3 in the center of nowhere, close to all commodities",
			Locale: "en",
		},
	}
	seeds["sources"] = []models.Source{
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "source1",
				},
			},
			URL: "https://rentofferwebsite.com/rents/123456789",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "source2",
				},
			},
			URL: "https://rentofferwebsite.com/rents/1011121314",
		},
	}
	seeds["prices"] = []models.Price{
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "price1",
				},
			},
			Value:    1200000.000,
			Currency: "EUR",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "price2",
				},
			},
			Value:    600.000,
			Currency: "EUR",
		},
	}
	seeds["locations"] = []models.Location{
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "loc1",
				},
			},
			GeoJSON: models.GeoJSON{
				Type: "Feature",
				Geometry: models.Geometry{
					Type:        "Point",
					Coordinates: []float64{2.3522219, 48.856614},
				},
				Properties: models.Properties{
					Name:    "Dinagat Islands",
					Country: "Philippines",
				},
			},
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "loc2",
				},
			},
			GeoJSON: models.GeoJSON{
				Type: "Feature",
				Geometry: models.Geometry{
					Type:        "Point",
					Coordinates: []float64{2.3488, 48.8534},
				},
				Properties: models.Properties{
					Name:    "Paris city",
					Country: "France",
				},
			},
		},
	}
	seeds["assets"] = []models.Asset{
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "asset1",
				},
			},
			URL:     "https://cdn.rentofferwebsite.com/img/123456789-main.jpg",
			Type:    "PICT",
			Storage: "EXT",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "asset2",
				},
			},
			URL:     "https://cdn.rentofferwebsite.com/img/1011121314-main.jpg",
			Type:    "PICT",
			Storage: "EXT",
		},
	}
	seeds["offerors"] = []models.Offeror{
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "offeror1",
				},
			},
			Name: "M. Smith",
			Email: &models.Email{
				Value: "smith@yopmail.com",
			},
			Phone: &models.Phone{
				CountryCode: "+33",
				Value:       "656453463",
			},
			Type: "PERS",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "offeror2",
				},
			},
			Name: "M. Doe",
			Email: &models.Email{
				Value: "doe@yopmail.com",
			},
			Phone: &models.Phone{
				CountryCode: "+33",
				Value:       "6785643785",
			},
			Type: "PERS",
		},
	}
	// Create documents
	seeds["rentoffers"] = []models.RentOfferItem{
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "ro1",
				},
			},
			Title:       []string{"title1"},
			Source:      "source1",
			Description: []string{"desc1"},
			Price:       []string{"price1"},
			Assets:      []string{"asset1"},
			Location:    "loc1",
			Offeror:     "offeror1",
		},
		{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: "ro2",
				},
			},
			Title:       []string{"title2"},
			Source:      "source2",
			Description: []string{"desc2"},
			Price:       []string{"price2"},
			Assets:      []string{"asset2"},
			Location:    "loc2",
			Offeror:     "offeror2",
		},
	}
	for itm := range seeds {
		col, err := db.Collection(ctx, itm)
		if err != nil {
			return err
		}
		meta, _, err := col.CreateDocuments(ctx, seeds[itm])
		if err != nil {
			if !driver.IsConflict(err) {
				log.Fatal(err)
			} else {
				cvrt := seeds[itm].(models.Base)
				docKey := string(cvrt.Key)
				meta, err := col.UpdateDocument(ctx, docKey, seeds[itm])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Updated document in collection '%s' in database '%s'\n", col.Name(), db.Name())
				fmt.Printf("Got metas : %+v\n", meta)
			}
		} else {
			fmt.Printf("Created document in collection '%s' in database '%s'\n", col.Name(), db.Name())
			fmt.Printf("Got metas : %+v\n", meta)
		}
	}
	return nil
}

func seedEdges(ctx context.Context, db driver.Database) error {
	var err error
	seeds := map[string]interface{}{
		"rentofferTitles": []Edge{
			{
				From: "labels/title1",
				To:   "rentoffers/ro1",
			},
			{
				From: "labels/title2",
				To:   "rentoffers/ro2",
			},
		},
		"rentofferSources": []Edge{
			{
				From: "sources/source1",
				To:   "rentoffers/ro1",
			},
			{
				From: "sources/source2",
				To:   "rentoffers/ro2",
			},
		},
		"rentofferDescriptions": []Edge{
			{
				From: "labels/desc1",
				To:   "rentoffers/ro1",
			},
			{
				From: "labels/desc2",
				To:   "rentoffers/ro2",
			},
		},
		"rentofferPrices": []Edge{
			{
				From: "prices/price1",
				To:   "rentoffers/ro1",
			},
			{
				From: "prices/price2",
				To:   "rentoffers/ro2",
			},
		},
		"rentofferAssets": []Edge{
			{
				From: "assets/asset1",
				To:   "rentoffers/ro1",
			},
			{
				From: "assets/asset2",
				To:   "rentoffers/ro2",
			},
		},
		"rentofferLocations": []Edge{
			{
				From: "locations/loc1",
				To:   "rentoffers/ro1",
			},
			{
				From: "locations/loc2",
				To:   "rentoffers/ro2",
			},
		},
		"rentofferOfferors": []Edge{
			{
				From: "offerors/offeror1",
				To:   "rentoffers/ro1",
			},
			{
				From: "offerors/offeror2",
				To:   "rentoffers/ro2",
			},
		},
	}
	for el := range seeds {
		var col driver.Collection
		var colType = driver.CollectionTypeEdge
		if col, err = db.CreateCollection(ctx, el, &driver.CreateCollectionOptions{
			Type: colType,
		}); err != nil {
			return err
		}
		meta, _, err := col.CreateDocuments(ctx, seeds[el])
		if err != nil {
			return err
		}
		fmt.Printf("Created document in collection '%s' in database '%s'\n", col.Name(), db.Name())
		fmt.Printf("Got metas : %+v\n", meta)
	}
	return nil
}
