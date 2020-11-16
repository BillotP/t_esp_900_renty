package models

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/BillotP/gorenty"
	"gorm.io/gorm"
)

// Target is a scrapping target for leboncoin real estate offers
type Target struct {
	ScrappingRequestItemID uuid.UUID `json:"scrappingRequestId" gorm:"index:unique_target"`
	Country                string    `json:"country" gorm:"index:unique_target"`
	Department             string    `json:"department" gorm:"index:unique_target"`
	City                   string    `json:"city"`
	Place                  string    `json:"place"`
	Budget                 []float64 `json:"budget"`
	Coordinates            []float64 `json:"coordinates"`
	Pages                  int       `json:"pages"`
}

// TargetItem is a db Item
type TargetItem struct {
	Base
	Target
}

// Category is a scrapping category for leboncoin real estate offers
type Category struct {
	ScrappingRequestItemID uuid.UUID `json:"scrappingRequestId" gorm:"index:unique_category"`
	Label                  string    `json:"label" gorm:"index:unique_category"`
	RealEstateType         string    `json:"realEstateType" gorm:"index:unique_category"`
}

// CategoryItem is a db Item
type CategoryItem struct {
	Base
	Category
}

// ScrappingRequest is a scrapping request for a user
type ScrappingRequest struct {
	BaseURL       string         `json:"baseUrl" gorm:"index:scrappingrequest_index"`
	RetryCount    int            `json:"retryCount"`
	MessageItem   MessageItem    `json:"messageItem,omitempty"`
	TargetItems   []TargetItem   `json:"targets"`
	CategoryItems []CategoryItem `json:"categories"`
	UserItemID    uuid.UUID      `json:"userId" gorm:"index:scrappingrequest_index"`
}

// ScrappingRequestItem is a db item
type ScrappingRequestItem struct {
	Base
	ScrappingRequest
}

// Save a new scrapping request in database
func (s ScrappingRequest) Save(db *gorm.DB) (*ScrappingRequestItem, error) {
	item := &ScrappingRequestItem{
		ScrappingRequest: s,
	}
	if err := db.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

// Update an existing scrapping request in database
func (s ScrappingRequest) Update(db *gorm.DB, uuid uuid.UUID) (*ScrappingRequestItem, error) {
	item := ScrappingRequestItem{
		Base: Base{
			UUID: uuid,
		},
		ScrappingRequest: s,
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Update): Will  update scrapping requests [%+v]\n", s)
	}
	if err := db.Model(&item).Updates(&item).Error; err != nil {
		fmt.Printf("Error(Update): %s\n", err.Error())
		return nil, err
	}
	return &item, nil
}

// HasCategory check if a scrapping request have a certain scrapping request category in his list
func (s ScrappingRequest) HasCategory(categorylabel string) bool {
	if s.CategoryItems == nil {
		return false
	}
	for el := range s.CategoryItems {
		if s.CategoryItems[el].Label == categorylabel {
			return true
		}
	}
	return false
}

// Delete a scrapping request in database
func (s *ScrappingRequest) Delete(db *gorm.DB, uuid uuid.UUID) (*ScrappingRequestItem, error) {
	item := &ScrappingRequestItem{
		Base: Base{
			UUID: uuid,
		},
	}
	if err := db.Delete(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

// GetScrappingRequestByID return a scrapping request by its id
func GetScrappingRequestByID(db *gorm.DB, uuid uuid.UUID) (*ScrappingRequestItem, error) {
	var scrappingrequest ScrappingRequestItem
	if err := db.Where("id = ?", uuid).Preload("MessageItem").
		Find(&scrappingrequest).Error; err != nil {
		fmt.Printf("Error(GetUserByID): %s\n", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	if goscrappy.Debug {
		fmt.Printf("Info(GetScrappingRequestByID): Found scrapping request : [%+v]\n", scrappingrequest.ScrappingRequest)
	}
	scrappingrequest.TargetItems = []TargetItem{}
	if err := db.Model(&scrappingrequest).
		Association("TargetItems").
		Find(&scrappingrequest.TargetItems); err != nil {
		fmt.Printf("Error(GetScrappingRequestByID): %s\n", err.Error())
		return nil, err
	}
	scrappingrequest.CategoryItems = []CategoryItem{}
	if err := db.Model(&scrappingrequest).
		Association("CategoryItems").
		Find(&scrappingrequest.CategoryItems); err != nil {
		fmt.Printf("Error(GetScrappingRequestByID): %s\n", err.Error())
		return nil, err
	}
	return &scrappingrequest, nil
}

// GetScrappingRequests return all scrapping request in db
func GetScrappingRequests(db *gorm.DB, limit *string, targets, categories, message *bool) (*[]ScrappingRequestItem, error) {
	var err error
	var scrappingrequests []ScrappingRequestItem
	if limit != nil {
		var limitval int
		if limitval, err = strconv.Atoi(*limit); err != nil {
			fmt.Printf("Error(GetScrappingRequests): Invalid limit param : %s\n", err.Error())
			return nil, ErrInvalidLimit
		} else if limitval < 1 || limitval > maxLimit {
			return nil, ErrInvalidLimit
		}
		if goscrappy.Debug {
			fmt.Printf("Info(GetScrappingRequests): Will looking for %v first scrapping request\n", limitval)
		}
		if err = db.Limit(limitval).Find(&scrappingrequests).Error; err != nil {
			fmt.Printf("Error(GetScrappingRequests): %s\n", err.Error())
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}

	} else if err = db.Find(&scrappingrequests).Error; err != nil {
		fmt.Printf("Error(GetScrappingRequests): %s\n", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	for el := range scrappingrequests {
		if targets != nil && *targets == true {
			scrappingrequests[el].TargetItems = []TargetItem{}
			if err := db.Model(&scrappingrequests[el]).
				Association("TargetItems").
				Find(&scrappingrequests[el].TargetItems); err != nil {
				fmt.Printf("Error(GetScrappingRequests): %s\n", err.Error())
				return nil, err
			}
		}
		if categories != nil && *categories == true {
			scrappingrequests[el].CategoryItems = []CategoryItem{}
			if err := db.Model(&scrappingrequests[el]).
				Association("CategoryItems").
				Find(&scrappingrequests[el].CategoryItems); err != nil {
				fmt.Printf("Error(GetScrappingRequests): %s\n", err.Error())
				return nil, err
			}
		}
		if message != nil && *message == true {
			if err := db.Model(&scrappingrequests[el]).
				Association("MessageItem").
				Find(&scrappingrequests[el].MessageItem); err != nil {
				fmt.Printf("Error(GetScrappingRequests): %s\n", err.Error())
				return nil, err
			}
		}
	}
	return &scrappingrequests, nil
}
