// Package models contains all db schemes
package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/BillotP/gorenty"
	"gorm.io/gorm"
)

var (
	// ErrInvalidMonth is returned when param month is invalid
	ErrInvalidMonth = errors.New("Invalid month param (should be an int from 1 to 12)")
)

// DoneJob is a done job model in redis queue (and in postgres report)
type DoneJob struct {
	AnnonceID   string    `json:"annonceId" gorm:"uniqueIndex"`
	Departement string    `json:"department"`
	SMSID       string    `json:"smsId"`
	SendedAt    time.Time `json:"sendedAt"`
}

// DoneJobItem is a contacted Leboncoin user
type DoneJobItem struct {
	gorm.Model `json:"-"`
	DoneJob
}

// IsTest check smsid if it contains the test variable
func (d DoneJob) IsTest() bool {
	return strings.Contains(d.SMSID, "test")
}

// Save a new done job in database
func (d DoneJob) Save(db *gorm.DB) (*DoneJobItem, error) {
	donejobItem := &DoneJobItem{
		DoneJob: d,
	}
	if err := db.Create(donejobItem).Error; err != nil {
		return nil, err
	}
	return donejobItem, nil
}

// DropAllDoneJobs clear all donejob items in db (!USE WITH CAUTION)
func DropAllDoneJobs(db *gorm.DB) error {
	if err := db.Session(&gorm.Session{
		AllowGlobalUpdate: true,
	}).Delete(&DoneJobItem{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAll done job(s), optionaly filtered by month and/or department
func GetAll(db *gorm.DB, month *string, department *string) (doneJobs []DoneJobItem, err error) {
	var mval int
	switch {
	case month != nil && department == nil:
		if mval, err = strconv.Atoi(*month); err != nil {
			fmt.Printf("Error(GetAllContacted): Invalid month param : %s\n", err.Error())
			return nil, ErrInvalidMonth
		}
		if mval < 1 || mval > 12 {
			return nil, ErrInvalidMonth
		}
		if goscrappy.Debug {
			fmt.Printf("Info(GetAllContacted): Will looking for month %v jobs\n", mval)
		}
		if err = db.Where("EXTRACT(MONTH FROM sended_at) = ?", mval).Find(&doneJobs).Error; err != nil {
			fmt.Printf("Error(GetAll): %s\n", err.Error())
			return nil, err
		}
		break
	case department != nil && month == nil:
		if err = db.Where("departement = ?", *department).Find(&doneJobs).Error; err != nil {
			fmt.Printf("Error(GetAll): %s\n", err.Error())
			return nil, err
		}
		break
	case department != nil && month != nil:
		if mval, err = strconv.Atoi(*month); err != nil {
			fmt.Printf("Error(GetAllContacted): Invalid month param : %s\n", err.Error())
			return nil, ErrInvalidMonth
		}
		if mval < 1 || mval > 12 {
			return nil, ErrInvalidMonth
		}
		if goscrappy.Debug {
			fmt.Printf("Info(GetAllContacted): Will looking for month %v and department %s job(s)\n", mval, *department)
		}
		if err = db.Where("departement = ?", *department).Where("EXTRACT(MONTH FROM sended_at) = ?", mval).Find(&doneJobs).Error; err != nil {
			fmt.Printf("Error(GetAll): %s\n", err.Error())
			return nil, err
		}
		break
	default:
		if err = db.Find(&doneJobs).Error; err != nil {
			fmt.Printf("Error(GetAll): %s\n", err.Error())
			return nil, err
		}
	}
	if goscrappy.Debug {
		fmt.Printf("Info(GetAllContacted): Got results [%+v]\n", doneJobs)
	}
	return doneJobs, err
}
