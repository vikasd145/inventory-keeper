package appliance

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ApplianceRepository interface {
	Get() (Appliance, error)
	Search() ([]Appliance, error)
	Create() (Appliance, error)
	Update(id int64) error
	GetAll() ([]Appliance, error)
}

type ApplianceModel struct {
	DB    *gorm.DB
	Model *Appliance
}

func NewApplianceModel(db *gorm.DB, model *Appliance) ApplianceRepository {
	return &ApplianceModel{
		DB:    db,
		Model: model,
	}
}

type Appliance struct {
	ID           int64  `json:"id" gorm:"primary_key"`
	SerialNumber string `json:"serial_number" gorm:"unique_index:idx_serial_number_brand_model"`
	Brand        string `json:"brand" gorm:"unique_index:idx_serial_number_brand_model"`
	Model        string `json:"model" gorm:"unique_index:idx_serial_number_brand_model"`
	Status       string `json:"status"`
	DateBought   string `json:"date_bought"`
}

func (c *ApplianceModel) Get() (Appliance, error) {
	err := c.DB.Model(c.Model).First(c.Model).Error
	return *c.Model, err
}

func (c *ApplianceModel) Search() ([]Appliance, error) {
	var cs []Appliance
	whereStr := ""
	if c.Model.SerialNumber != "" {
		whereStr += fmt.Sprintf("%s=\"%s\" and ", "serial_number", c.Model.SerialNumber)
	}
	if c.Model.Brand != "" {
		whereStr += fmt.Sprintf("%s=\"%s\" and ", "brand", c.Model.Brand)
	}
	if c.Model.Model != "" {
		whereStr += fmt.Sprintf("%s=\"%s\" and ", "model", c.Model.Model)
	}
	if c.Model.Status != "" {
		whereStr += fmt.Sprintf("%s=\"%s\" and ", "status", c.Model.Status)
	}
	if c.Model.DateBought != "" {
		whereStr += fmt.Sprintf("%s=\"%s\" and ", "date_bought", c.Model.DateBought)
	}
	if len(whereStr) >= 4 && whereStr[len(whereStr)-4:] == "and " {
		whereStr = whereStr[:len(whereStr)-4]
	}
	err := c.DB.Model(c.Model).Where(whereStr).Find(&cs).Error
	return cs, err
}

func (c *ApplianceModel) GetAll() ([]Appliance, error) {
	var cs []Appliance
	err := c.DB.Model(c.Model).Find(&cs).Error
	return cs, err
}

func (c *ApplianceModel) Create() (Appliance, error) {
	err := c.DB.Model(c.Model).Create(c.Model).Error
	return *c.Model, err
}

func (c *ApplianceModel) Update(id int64) (err error) {
	err = c.DB.Model(c.Model).Where("id=?", id).Update(c.Model).Error
	return
}
