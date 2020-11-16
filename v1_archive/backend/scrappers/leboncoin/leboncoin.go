package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/BillotP/gorenty"
	mod "github.com/BillotP/gorenty/models"
	"github.com/BillotP/gorenty/v2/models"
)

// All Leboncoin categories
const (
	// Leboncoin is a french rent offer listing website
	Leboncoin = "https://www.leboncoin.fr"
	// LBCSearch ...
	LBCSearch = "/recherche"
	// LBCPage is the page result index
	LBCPage = "page"
	// LBCRealEstateCategoryKey indicate the kind off real estate contract (colocation, classic rents ...)
	LBCRealEstateCategoryKey = "category"
	// LBCLocationCodeKey is the french regions code label
	LBCLocationCodeKey = "locations"
	// LBCRealEstateTypeKey indicate whether the good is a flat, a house etc
	LBCRealEstateTypeKey = "real_estate_type"
	// LBCRealEstateSellerType is c to c vs b to c
	LBCRealEstateSellerType = "owner_type"
	// LBCRealEstateEquipmentType indicate wether the good is equiped (furnished) or not
	LBCRealEstateEquipmentType = "furnished"
	// LBCOwnerTypeKey ...
	LBCOwnerTypeKey                = "owner_type"
	AllRealEstateCategoryID        = "8"
	RentsCategoryID                = "10"
	ColocationCategoryID           = "11"
	TouristicrentCategoryID        = "12"
	BusinesslocalsCategoryID       = "13"
	AlsaceLocationCode             = "r_1"
	AquitaineLocationCode          = "r_2"
	AuvergneLocationCode           = "r_3"
	BasseNormandieLocationCode     = "r_4"
	BourgogneLocationCode          = "r_5"
	BretagneLocationCode           = "r_6"
	CentreLocationCode             = "r_7"
	ChampagneArdenneLocationCode   = "r_8"
	CorseLocationCode              = "r_9"
	FrancheCompteLocationCode      = "r_10"
	HauteNormandieLocationCode     = "r_11"
	IleDeFranceLocationCode        = "r_12"
	LanguedocRousillonLocationCode = "r_13"
	LimousinLocationCode           = "r_14"
	LorraineLocationCode           = "r_15"
	MidiPyreneesLocationCode       = "r_16"
	NordPasDeCalaisLocationCode    = "r_17"
	PaysDeLaLoireLocationCode      = "r_18"
	PicardieLocationCode           = "r_19"
	PoitouCharentesLocationCode    = "r_20"
	ProvenceAlpesLocationCode      = "r_21"
	RhoneAlpesLocationCode         = "r_22"
	GuadeloupeLocationCode         = "r_23"
	MartiniqueLocationCode         = "r_24"
	GuyaneLocationCode             = "r_25"
	ReunionLocationCode            = "r_26"
	HouseRealEstateType            = "1"
	FlatRealEstateType             = "2"
	LandRealEstateType             = "3"
	ParkingRealEstateType          = "4"
	OtherRealEstateType            = "5"
	IndividualOwnerType            = "private"
	ProfessionalOwnerType          = "pro"
	FurnishedEquipmentType         = "1"
	UnFurnishedEquipmentType       = "2"
)

// LBCPhoneAPIResponse ...
type LBCPhoneAPIResponse struct {
	Utils struct {
		Status  string `json:"status"`
		Annonce string `json:"phonenumber"`
	} `json:"utils"`
}

// LBCPhoneAPIURL ...
const LBCPhoneAPIURL = "https://api.leboncoin.fr/api/utils/phonenumber.json"

// SetDatas set the required params for Leboncoin phone number api call
func SetDatas(annonceID string) (data url.Values) {
	data = url.Values{}
	data.Set("app_id", "leboncoin_web_utils")
	data.Set("key", "54bb0281238b45a03f0ee695f73e704f")
	data.Set("list_id", annonceID)
	data.Set("text", "1")
	return data
}

// GetSearchURL return the leboncoin search page url for a department and a search page
func GetSearchURL(target mod.Target, page int) string {
	return fmt.Sprintf(
		"%s%s/?%s=%s&%s=%s&%s=%s&page=%v",
		Leboncoin,
		LBCSearch,
		LBCRealEstateCategoryKey,
		RentsCategoryID,
		LBCOwnerTypeKey,
		IndividualOwnerType,
		LBCLocationCodeKey,
		// target.Department,
		target.City,
		page,
	)
}

// IsValidAnnonceID check if annonceID is valid
func IsValidAnnonceID(annonceID string) bool {
	annonceID = strings.ReplaceAll(annonceID, ".htm", "")
	if _, err := strconv.ParseInt(annonceID, 10, 64); err == nil {
		return true
	}
	return false
}

// GetPhoneNumber from leboncoin phone number api call
func GetPhoneNumber(annonceID string, annonceType string) (string, *mod.RequestError) {
	data := SetDatas(annonceID)
	req, err := http.NewRequest(
		http.MethodPost,
		LBCPhoneAPIURL,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", &mod.RequestError{
			Response: err.Error(),
		}
	}
	SetHeaders(req, annonceID, annonceType)
	resp, err := client.Do(req)
	if err != nil {
		return "", &mod.RequestError{
			Response: err.Error(),
		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", &mod.RequestError{
			Response: err.Error(),
		}
	}
	if resp.StatusCode != http.StatusOK {
		if goscrappy.Debug && resp.StatusCode == http.StatusForbidden {
			ip, _ := GetMyIP()
			fmt.Printf("Info(GetPhoneNumber): Request blocked from ip [%s]\n", ip)
		}
		client.CloseIdleConnections()
		return "", &mod.RequestError{
			Status:   resp.StatusCode,
			Response: string(body),
		}
	}
	client.CloseIdleConnections()
	var apiResponse LBCPhoneAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", &mod.RequestError{
			Response: err.Error(),
		}
	}
	return apiResponse.Utils.Annonce, nil
}

// GetAnnoncesInfo return annonce informations from annonce URLs
func GetAnnoncesInfo(jobs []mod.Job) ScrappingInfos {
	imgValues := map[string][]models.Asset{}
	titleLabelValues := map[string]models.Label{}
	descLabelValues := map[string]models.Label{}
	priceValues := map[string]models.Price{}
	geolocValues := map[string]models.Location{}
	surfaceValues := map[string]*models.Surface{}
	offerorValues := map[string]models.Offeror{}

	c := colly.NewCollector(
		colly.UserAgent(userAgents[0]),
		colly.AllowURLRevisit(),
	)
	// if goscrappy.Debug {
	// 	fmt.Printf("Info(GetAnnonces): connecting to proxy [%s]\n", proxyURL.Hostname())
	// }
	// c.SetProxy(proxyString)
	c.SetRequestTimeout(20 * time.Second)
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	for _, job := range jobs {
		url := fmt.Sprintf("%s/%s/%s.htm", Leboncoin, job.Type, job.ID)
		imgValues[job.ID] = []models.Asset{}
		c.OnHTML("img", func(e *colly.HTMLElement) {
			exist := false
			offerImageLink := e.Attr("src")
			if goscrappy.Debug {
				fmt.Printf("Info(GetAnnoncesInfo): Got img url %s for job id [%s]\n", offerImageLink, job.ID)
			}
			for el := range imgValues[job.ID] {
				if imgValues[job.ID][el].URL == offerImageLink {
					exist = true
				}
			}
			if !exist {
				imgValues[job.ID] = append(imgValues[job.ID], models.Asset{
					URL:     offerImageLink,
					Type:    "PIC",
					Storage: "EXT",
				})
			}
		})
		c.OnHTML("h1[data-qa-id]", func(e *colly.HTMLElement) {
			if e.Attr("data-qa-id") == "adview_title" {
				offerTitle := e.Text
				if goscrappy.Debug {
					fmt.Printf("Info(GetAnnoncesInfo): Got title [%s] for job id [%s]\n", offerTitle, job.ID)
				}
				titleLabelValues[job.ID] = models.Label{
					Value:  offerTitle,
					Locale: "fr",
				}
			}
		})
		c.OnHTML("div[data-qa-id]", func(e *colly.HTMLElement) {
			offerAttr := e.Attr("data-qa-id")
			switch offerAttr {
			case "adview_description_container":
				offerDescription := e.Text
				if goscrappy.Debug {
					fmt.Printf("Info(GetAnnoncesInfo): Got description for job id [%s]\n", job.ID)
				}
				descLabelValues[job.ID] = models.Label{
					Value:  offerDescription,
					Locale: "fr",
				}
				break
			case "criteria_item_rooms":
				offerRooms := e.Text
				elems := strings.Split(offerRooms, "Pièces")
				numval, _ := strconv.ParseInt(elems[1], 10, 64)
				if surfaceValues[job.ID] != nil {
					surf := surfaceValues[job.ID]
					surf.Rooms = numval
					surfaceValues[job.ID] = surf
				} else {
					surfaceValues[job.ID] = &models.Surface{
						Rooms: numval,
					}
				}
				break
			case "criteria_item_square":
				offerSurface := e.Text
				offerSurface = strings.ReplaceAll(offerSurface, "Surface", "")
				elems := strings.Split(offerSurface, " ")
				numval, _ := strconv.ParseFloat(elems[0], 64)
				if goscrappy.Debug {
					fmt.Printf("Info(GetAnnoncesInfo): Got surface %v %s for job id [%s]\n", numval, elems[1], job.ID)
				}
				if surfaceValues[job.ID] != nil {
					surf := surfaceValues[job.ID]
					surf.Value = numval
					surf.Unit = elems[1]
					surfaceValues[job.ID] = surf
				} else {
					surfaceValues[job.ID] = &models.Surface{
						Value: numval,
						Unit:  elems[1],
					}
				}
				break
			case "adview_price":
				offerPriceText := e.Text
				clean := strings.Split(offerPriceText, "€")
				trimed := strings.ReplaceAll(clean[0], " ", "")
				if goscrappy.Debug {
					fmt.Printf("Info(GetAnnoncesInfo): Got price [%v] for job id [%s]\n", trimed, job.ID)
				}
				if val, err := strconv.ParseFloat(trimed, 64); err == nil && val > 0 {
					priceValues[job.ID] = models.Price{
						Value:     val,
						Currency:  "EUR",
						Condition: clean[1],
					}
				}
				break
			case "adview_contact_container":
				offerOfferorText := e.Text
				clean := strings.ReplaceAll(offerOfferorText, "avatar", "")
				clean = strings.ReplaceAll(clean, "Envoyer un message", "")
				clean = strings.ReplaceAll(clean, "Voir le numéro", "")
				if goscrappy.Debug {
					fmt.Printf("Info(GetAnnoncesInfo): Got offeror name [%s] for job id [%s]\n", clean, job.ID)
				}
				offerorValues[job.ID] = models.Offeror{
					Name: clean,
					Type: "individual",
				}
			default:
				break
			}
		})
		// Need to improve this!!
		if geoloc, err := DoSearch(job.Summary + " Bordeaux"); err == nil && geoloc != nil && len(geoloc.Features) > 0 {
			geolocValues[job.ID] = models.Location{
				GeoJSON: models.GeoJSON{
					Type: geoloc.Features[0].Type,
					Properties: models.Properties{
						Name:    geoloc.Features[0].Properties.Label,
						Country: "France",
					},
					Geometry: models.Geometry{
						Type:        geoloc.Features[0].Geometry.Type,
						Coordinates: geoloc.Features[0].Geometry.Coordinates,
					},
				},
			}
		} else {
			geolocValues[job.ID] = models.Location{
				GeoJSON: models.GeoJSON{
					Type: "Feature",
					Properties: models.Properties{
						Name:    "Bordeaux",
						Country: "France",
					},
					Geometry: models.Geometry{
						Type:        "Point",
						Coordinates: []float64{-0.587876, 44.853383},
					},
				},
			}
		}
		c.Visit(url)
	}
	if goscrappy.Debug {
		fmt.Printf(
			"Descriptions (%v)\nTitles (%v)\nPrices (%v)\nOfferors (%v)\nAssets (%v)\nSurfaces (%v)",
			len(descLabelValues),
			len(titleLabelValues),
			len(priceValues),
			len(offerorValues),
			len(imgValues),
			len(surfaceValues),
		)
	}

	return ScrappingInfos{
		Assets:       imgValues,
		Titles:       titleLabelValues,
		Descriptions: descLabelValues,
		Prices:       priceValues,
		Locations:    geolocValues,
		Offerors:     offerorValues,
		Surfaces:     surfaceValues,
	}
}

// GetAnnonces from leboncoin website for each department in SEARCHDPMTS
func GetAnnonces(scrappingrequest mod.ScrappingRequestItem) []mod.Job {
	var dpt string
	var page int
	var res []mod.Job
	count := len(scrappingrequest.TargetItems)
	if count == 0 {
		fmt.Printf("Error(GetAnnonces): No such scrapping targets\n")
		return nil
	} else if goscrappy.Debug {
		logs := fmt.Sprintf("Info(GetAnnonces): Got %v scrapping targets to do\n",
			count)
		fmt.Printf(logs)
		// botClient.SendMessage(logs)
	}
	// proxyString := proxyURL.String()
	c := colly.NewCollector(
		colly.UserAgent(userAgents[0]),
		colly.AllowURLRevisit(),
	)
	// if goscrappy.Debug {
	// 	fmt.Printf("Info(GetAnnonces): connecting to proxy [%s]\n", proxyURL.Hostname())
	// }
	// c.SetProxy(proxyString)
	c.SetRequestTimeout(20 * time.Second)
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		el := e.Attr("href")
		path := strings.Split(el, "/")
		if len(path) == 4 && IsValidAnnonceID(path[2]) {
			annonceID := strings.ReplaceAll(path[2], ".htm", "")
			if goscrappy.Debug {
				fmt.Printf(
					"Info(run): found ID [%s] for department %s on page %v\n",
					annonceID,
					dpt,
					page,
				)
			}
			// Checking for published date
			if strings.Contains(e.Text, "Aujourd'hui,") &&
				scrappingrequest.HasCategory(path[1]) {
				res = append(res, mod.Job{
					ID:                 annonceID,
					Type:               path[1],
					Departement:        dpt,
					Summary:            e.Text,
					ScrappingRequestID: scrappingrequest.UUID,
					CollectedAt:        time.Now(),
				})
			} else if goscrappy.Debug {
				fmt.Printf("Info(run): Annonce ID [%s] is not from today or not in scrapping request category (%s / %s)\n",
					annonceID,
					e.Text,
					path[1])
			}
		}

	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Accept-Language", "fr-FR,fr;q=0.5")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		fmt.Printf("Info(GetAnnonces): visiting [%s]\n", r.URL)

	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf(
			"Error(GetAnnonces): request URL %s failed with response [%s] got error %s\n",
			r.Request.URL,
			string(r.Body),
			err.Error(),
		)
		// fmt.Printf("Error(GetAnnonces): Got error, sleeping %v s before retry...\n", WaitBeforeRetry)
		// time.Sleep(WaitBeforeRetry)
		// panic(err)
	})
	// Visit each targetted department
	for d, target := range scrappingrequest.TargetItems {
		dpt = scrappingrequest.TargetItems[d].Department
		for i := 1; i <= scrappingrequest.TargetItems[d].Pages; i++ {
			page = i
			url := GetSearchURL(target.Target, page)
			if err := c.Visit(url); err != nil {
				// fmt.Printf("Error(GetAnnonces): Got error, sleeping %vs before retry...\n", WaitBeforeRetry)
				// time.Sleep(WaitBeforeRetry)
				c.Visit(url)
			}
		}
	}
	return res
}
