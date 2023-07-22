package appliance

import (
	"fmt"
	"github.com/vikasd145/inventory-keeper/pkg/sql"
)

type Appliance struct {
	ID           int64  `json:"id" gorm:"primary_key"`
	SerialNumber string `json:"serial_number" gorm:"unique_index:idx_serial_number_brand_model"`
	Brand        string `json:"brand" gorm:"unique_index:idx_serial_number_brand_model"`
	Model        string `json:"model" gorm:"unique_index:idx_serial_number_brand_model"`
	Status       string `json:"status"`
	DateBought   string `json:"date_bought"`
}

func (c *Appliance) Get() error {
	err := sql.DB.Model(c).First(c).Error
	return err
}

func (c *Appliance) Search() ([]Appliance, error) {
	var cs []Appliance
	whereStr := ""
	if c.SerialNumber != "" {
		whereStr += fmt.Sprintf("%s=\"%s\"", "serial_number", c.SerialNumber)
	}
	if c.Brand != "" {
		whereStr += fmt.Sprintf("%s=\"%s\"", "brand", c.Brand)
	}
	if c.Model != "" {
		whereStr += fmt.Sprintf("%s=\"%s\"", "model", c.Model)
	}
	if c.Status != "" {
		whereStr += fmt.Sprintf("%s=\"%s\"", "status", c.Status)
	}
	if c.DateBought != "" {
		whereStr += fmt.Sprintf("%s=\"%s\"", "date_bought", c.DateBought)
	}
	err := sql.DB.Model(c).Where(whereStr).Find(&cs).Error
	return cs, err
}

func (c *Appliance) GetAll() ([]Appliance, error) {
	var cs []Appliance
	err := sql.DB.Model(c).Find(&cs).Error
	return cs, err
}

func (c *Appliance) Create() (err error) {
	err = sql.DB.Model(c).Create(c).Error
	return
}

func (c *Appliance) Update(id int64) (err error) {
	err = sql.DB.Model(c).Where("id=?", id).Update(c).Error
	return
}
