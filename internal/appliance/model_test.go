package appliance

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
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

	testCases := []struct {
		name     string
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "Success",
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{"id", "serial_number", "brand", "model", "status", "date_bought"}).
					AddRow(1, "SN1234", "Whirlpool", "Refrigerator", "Active", "2023-07-23")

				mock.ExpectQuery("^SELECT (.+) FROM `appliances` WHERE `appliances`.`serial_number` = SN1234 ORDER BY `appliances`.`id` ASC LIMIT 1").WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure",
			mockFunc: func() {
				mock.ExpectQuery("^SELECT (.+) FROM `appliances` WHERE `appliances`.`deleted_at` IS NULL ORDER BY `appliances`.`id` ASC LIMIT 1").WillReturnError(errors.New("error fetching data"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			repo := NewApplianceModel(db, &Appliance{SerialNumber: "SN1234"})
			_, err := repo.Get()
			if tc.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCreate(t *testing.T) {
	db, mock := setupDBMock()
	defer db.Close()

	testCases := []struct {
		name     string
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "Success",
			mockFunc: func() {
				mock.ExpectBegin()
				mock.ExpectExec("^INSERT INTO `appliances` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Failure",
			mockFunc: func() {
				mock.ExpectBegin()
				mock.ExpectExec("^INSERT INTO `appliances` (.+)$").WillReturnError(errors.New("error inserting data"))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			repo := NewApplianceModel(db, &Appliance{
				SerialNumber: "SN1234",
				Brand:        "Whirlpool",
				Model:        "Refrigerator",
				Status:       "Active",
				DateBought:   "2023-07-23",
			})
			_, err := repo.Create()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdate(t *testing.T) {
	db, mock := setupDBMock()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `appliances` SET (.+) WHERE (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewApplianceModel(db, &Appliance{
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
