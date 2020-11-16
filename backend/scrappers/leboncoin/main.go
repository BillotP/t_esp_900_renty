package main

import (
	"fmt"
	"log"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/cespare/xxhash"
	"github.com/BillotP/gorenty"
	mod "github.com/BillotP/gorenty/models"
	"github.com/BillotP/gorenty/v2/models"
	"github.com/BillotP/gorenty/v2/service"
)

// ScrappingInfos is meta info for each job id key
type ScrappingInfos struct {
	Assets       map[string][]models.Asset
	Titles       map[string]models.Label
	Descriptions map[string]models.Label
	Prices       map[string]models.Price
	Locations    map[string]models.Location
	Offerors     map[string]models.Offeror
	Surfaces     map[string]*models.Surface
}

// AddOfferorPhone ...
func AddOfferorPhone(jobs []mod.Job, infos map[string]models.Offeror) map[string]models.Offeror {
	for el := range jobs {
		var phone string
		var key = jobs[el].ID
		var err *mod.RequestError
		if phone, err = GetPhoneNumber(key, jobs[el].Type); err != nil {
			fmt.Printf("Error(AddOfferorPhone): Got bad resp: %s\n", err.Error())
			continue
		}
		offr := infos[key]
		offr.Phone = &models.Phone{
			Value:       phone,
			CountryCode: "+33",
		}
		infos[key] = offr
	}
	return infos
}

// ToRentOffers serialize jobs and scrappinginfos into rentoffer model
func ToRentOffers(jobs []mod.Job, datas ScrappingInfos) []models.RentOffer {
	var tosave []models.RentOffer
	for el := range jobs {
		key := jobs[el].ID
		url := fmt.Sprintf("%s/%s/%s.htm", Leboncoin, jobs[el].Type, jobs[el].ID)
		uniqueKey := xxhash.Sum64String(url)
		nro := models.RentOffer{
			Base: models.Base{
				DocumentMeta: driver.DocumentMeta{
					Key: fmt.Sprintf("%v", uniqueKey),
				},
			},
			Title:        []models.Label{datas.Titles[key]},
			Source:       models.Source{URL: url},
			Description:  []models.Label{datas.Descriptions[key]},
			Price:        []models.Price{datas.Prices[key]},
			Assets:       datas.Assets[key],
			CreatedAt:    time.Now(),
			ExpiredAt:    time.Now().Add((168 * 2) * time.Hour), // 2 weeks later
			Location:     datas.Locations[key],
			Offeror:      datas.Offerors[key],
			Surface:      *datas.Surfaces[key],
			Requirements: nil,
		}
		tosave = append(tosave, nro)
	}
	return tosave
}

// SaveResults from leboncoin scrapping
func SaveResults(data []models.RentOffer) (saved int, err error) {
	var dbservice *service.Repository
	if dbservice, err = service.New("renty-dev"); err != nil {
		return saved, err
	}
	var rentofferService = service.RentOfferItem{
		Repository: *dbservice,
	}
	for el := range data {
		if len(data[el].Assets) == 0 {
			data[el].Assets = []models.Asset{
				{
					URL:     "https://via.placeholder.com/150",
					Type:    "PIC",
					Storage: "EXT",
				},
			}
		}
		if _, err = rentofferService.Create(data[el]); err != nil {
			fmt.Printf("Got error saving %+v\n", data[el])
			continue
		}
		saved++
	}
	return saved, err
}

// DefaultScrappingRequestItem is the default scrapping request
var DefaultScrappingRequestItem = mod.ScrappingRequestItem{
	ScrappingRequest: mod.ScrappingRequest{
		TargetItems: []mod.TargetItem{
			{
				Target: mod.Target{
					Department: "33",
					City:       "Bordeaux",
					Pages:      1,
				},
			},
		},
		CategoryItems: []mod.CategoryItem{
			{
				Category: mod.Category{
					Label: "locations",
				},
			},
		},
	},
}

func main() {
	var err error
	var saved = 0
	jobs := GetAnnonces(DefaultScrappingRequestItem)
	if goscrappy.Debug {
		fmt.Printf("Info(leboncoin): Found %v offers\n", len(jobs))
	}
	datas := GetAnnoncesInfo(jobs)
	nrentoffers := ToRentOffers(jobs, datas)
	if saved, err = SaveResults(nrentoffers); err != nil {
		log.Fatal(err)
	}
	if goscrappy.Debug {
		fmt.Printf("Info(leboncoin): Having saved %v rent offfers\n", saved)
	}
}
