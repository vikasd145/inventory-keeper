package sql

import (
	"github.com/vikasd145/inventory-keeper/internal/appliance"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setupDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		panic(err)
	}

	return gdb, mock
}

func TestGet(t *testing.T) {
	db, mock := setupDBMock()
	defer db.Close()

	mock.ExpectQuery("^SELECT (.+) FROM `appliances` WHERE `appliances`.`deleted_at` IS NULL ORDER BY `appliances`.`id` ASC LIMIT 1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "serial_number", "brand", "model", "status", "date_bought"}).
			AddRow(1, "SN1234", "Whirlpool", "Refrigerator", "Active", "2023-07-23"))

	repo := NewApplianceModel(db, &appliance.Appliance{})

	result, err := repo.Get()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "SN1234", result.SerialNumber)
	assert.Equal(t, "Whirlpool", result.Brand)
	assert.Equal(t, "Refrigerator", result.Model)
	assert.Equal(t, "Active", result.Status)
	assert.Equal(t, "2023-07-23", result.DateBought)
}

func TestSearch(t *testing.T) {
	db, mock := setupDBMock()
	defer db.Close()

	mock.ExpectQuery("^SELECT (.+) FROM `appliances` WHERE (.+)$").
		WillReturnRows(sqlmock.NewRows([]string{"id", "serial_number", "brand", "model", "status", "date_bought"}).
			AddRow(1, "SN1234", "Whirlpool", "Refrigerator", "Active", "2023-07-23"))

	repo := appliance.NewApplianceModel(db, &appliance.Appliance{
		SerialNumber: "SN1234",
	})

	result, err := repo.Search()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, "SN1234", result[0].SerialNumber)
}

func TestCreate(t *testing.T) {
	db, mock := setupDBMock()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `appliances` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := appliance.NewApplianceModel(db, &appliance.Appliance{
		SerialNumber: "SN1234",
		Brand:        "Whirlpool",
		Model:        "Refrigerator",
		Status:       "Active",
		DateBought:   "2023-07-23",
	})

	result, err := repo.Create()
	assert.Nil(t, err)
	assert.Equal(t, "SN1234", result.SerialNumber)
	assert.Equal(t, "Whirlpool", result.Brand)
}

func TestUpdate(t *testing.T) {
	db, mock := setupDBMock()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `appliances` SET (.+) WHERE (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := appliance.NewApplianceModel(db, &appliance.Appliance{
		ID:           1,
		SerialNumber: "SN1234",
		Brand:        "Whirlpool",
		Model:        "Refrigerator",
		Status:       "Active",
		DateBought:   "2023-07-23",
	})

	err := repo.Update(1)
	assert.Nil(t, err)
}
