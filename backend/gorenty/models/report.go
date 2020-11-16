// Package models contains all db schemes
package models

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"sort"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/BillotP/gorenty"
	"gorm.io/gorm"
)

///////////////////////
// SCRAPPY REPORT CONTENT

// DepartmentReport is a per departement report (included in global report)
type DepartmentReport struct {
	Department   string    `json:"department"`
	Annonces     int64     `json:"annonces"`
	SMSSent      int64     `json:"smsSent"`
	ErrorBlocked int64     `json:"errorBlocked"`
	ErrorNoPhone int64     `json:"errorNoPhone"`
	ReportItemID uuid.UUID `json:"reportItemId"`
}

// DepartmentReportItem is a db item
type DepartmentReportItem struct {
	Base `json:"-"`
	DepartmentReport
}

// ReportType enumeration
type ReportType int64

const (
	// CUSTOM is a custom date range
	CUSTOM ReportType = iota
	// DAILY is a daily report
	DAILY
	// MONTHLY is a monthly report
	MONTHLY
	// WEEKLY is a weekly report
	WEEKLY
	// YEARLY is a yearly report
	YEARLY
)

var types = []string{
	"custom",
	"daily",
	"monthly",
	"weekly",
	"yearly",
}

func (r ReportType) String() string {
	return types[r]
}

// FromString get the report type enum value from string
func FromString(el string) ReportType {
	val := sort.SearchStrings(types, el)
	if val == len(el) {
		return CUSTOM
	}
	rt := ReportType(val)
	if goscrappy.Debug {
		fmt.Printf("Info(FromString): Got %v report type from %s string\n",
			rt, el)
	}
	return rt
}

// Scan a custom report type value
func (r *ReportType) Scan(value interface{}) error { *r = ReportType(value.(int64)); return nil }

// Value return a custom report type value
func (r ReportType) Value() (driver.Value, error) { return int64(r), nil }

// Report is the report format
type Report struct {
	Type                   ReportType             `json:"-"`
	Period                 string                 `json:"period" gorm:"-"`
	Annonces               int64                  `json:"annonces"`
	SMSSent                int64                  `json:"smsSent"`
	ErrorBlocked           int64                  `json:"errorBlocked"`
	ErrorOther             int64                  `json:"errorOther"`
	ScrappingRequestItemID uuid.UUID              `json:"scrappingRequestId"`
	DepartmentReportItems  []DepartmentReportItem `json:"perDepartment"`
	FromDate               time.Time              `json:"fromDate" gorm:"index:report_index"`
	ToDate                 time.Time              `json:"toDate" gorm:"index:report_index"`
}

// ReportItem is a message in db
type ReportItem struct {
	Base
	Report
}

// Save a new report in database
func (m Report) Save(db *gorm.DB) (*ReportItem, error) {
	var err error
	reportitem := &ReportItem{
		Report: m,
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Save): Will try to save %+v\n", reportitem.Report)
	}
	if err = db.Create(reportitem).Error; err != nil {
		return nil, err
	}
	return reportitem, nil
}

// Update a report item in database
func (m Report) Update(db *gorm.DB, uuid uuid.UUID) (*ReportItem, error) {
	var err error
	reportitem := &ReportItem{
		Base: Base{
			UUID: uuid,
		},
		Report: m,
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Save): Will try to update %+v\n", reportitem.Report)
	}
	if err = db.Model(reportitem).Updates(reportitem).Error; err != nil { // Remember to update if exist
		return nil, err
	}
	return reportitem, nil
}

// Delete a report item in database
func (m *Report) Delete(db *gorm.DB, uuid uuid.UUID) (*ReportItem, error) {
	var err error
	reportitem := &ReportItem{
		Base: Base{
			UUID: uuid,
		},
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Delete): Will try to delete %+v\n", reportitem.ID)
	}
	if err = db.Delete(reportitem).Error; err != nil {
		return nil, err
	}
	return reportitem, nil
}

// GetAllReports return all reports in postgres db
func GetAllReports(db *gorm.DB) (reports []ReportItem, err error) {
	if err = db.Find(&reports).Error; err != nil {
		return nil, err
	}
	for el := range reports { // Convert reporttype to human readable string
		reports[el].Period = reports[el].Type.String()
		if err = db.Model(&reports[el]).
			Association("DepartmentReportItems").
			Find(&reports[el].DepartmentReportItems); err != nil {
			fmt.Printf("Error(GetAllReports): %s\n", err.Error())
			return nil, err
		}
	}
	return reports, nil
}

// GetReportByID return a reportitem by its id
func GetReportByID(db *gorm.DB, uuid uuid.UUID) (*ReportItem, error) {
	var report ReportItem
	if err := db.Where("id = ?", uuid).Find(&report).Error; err != nil {
		fmt.Printf("Error(GetReportByID): %s\n", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	if err := db.Model(&report).
		Association("DepartmentReportItems").
		Find(&report.DepartmentReportItems); err != nil {
		fmt.Printf("Error(DepartmentReportItems): %s\n", err.Error())
		return nil, err
	}
	return &report, nil
}

// GetScrappingReportTemplate return an html template string filled with the required variables
func GetScrappingReportTemplate(report Report) *string {
	var err error
	var t *template.Template

	// TODO: put this template in database (get this shit out from server code)
	if t, err = template.ParseFiles("./scrappy-report.html"); err != nil {
		fmt.Printf("Error(getTemplate1): %s\n", err.Error())
		return nil
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, report); err != nil {
		fmt.Printf("Error(getTemplate): %s\n", err.Error())
		return nil
	}

	result := tpl.String()
	return &result
}

// GetReport in database from its type
func GetReport(db *gorm.DB, period ReportType) (*ReportItem, error) {
	var item ReportItem
	if err := db.Where("type = ?", period).Find(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
