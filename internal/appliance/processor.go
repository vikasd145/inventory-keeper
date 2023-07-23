package appliance

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vikasd145/inventory-keeper/pkg/logger"
	"net/http"
	"strconv"
)

type ApplianceProcessorInterface interface {
	ViewAppliance(g *gin.Context)
	SearchAppliance(g *gin.Context)
	CreateAppliance(g *gin.Context)
	UpdateAppliance(g *gin.Context)
	ViewAllAppliance(g *gin.Context)
}

type ApplianceProcessor struct {
	Repo ApplianceRepository
	DB   *gorm.DB
}

func NewApplianceProcessor(repo ApplianceRepository, db *gorm.DB) ApplianceProcessorInterface {
	return &ApplianceProcessor{
		Repo: repo,
		DB:   db,
	}
}

func (p *ApplianceProcessor) ViewAppliance(g *gin.Context) {
	var id int
	var err error
	ids := g.Param("id")
	id, err = strconv.Atoi(ids)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error converting id to integer",
		})
		return
	}
	app := &Appliance{
		ID: int64(id),
	}
	p.Repo = &ApplianceModel{
		Model: app,
		DB:    p.DB,
	}
	respApp, err := p.Repo.Get()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error finding user from id",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     respApp,
		"debugMsg": "",
	})
	return
}

func (p *ApplianceProcessor) SearchAppliance(g *gin.Context) {
	appReq := &ApplianceReq{}
	err := g.ShouldBindJSON(appReq)
	if err != nil {
		logger.ErrorLogger.Printf("error in CreateAppliance ShouldBindJSON, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "Invalid json provided",
		})
		return
	}
	logger.InfoLogger.Printf("CreateAppliance, req:%v", appReq)
	app := &Appliance{
		SerialNumber: appReq.SerialNumber,
		Brand:        appReq.Brand,
		Model:        appReq.Model,
		Status:       appReq.Status,
		DateBought:   appReq.DateBought,
	}
	p.Repo = &ApplianceModel{
		Model: app,
		DB:    p.DB,
	}
	appList, err := p.Repo.Search()
	if err != nil {
		logger.ErrorLogger.Printf("error in creating appliance, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in creating appliance",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     appList,
		"debugMsg": "",
	})
	return
}

func (p *ApplianceProcessor) CreateAppliance(g *gin.Context) {
	appReq := &ApplianceReq{}
	err := g.ShouldBindJSON(appReq)
	if err != nil {
		logger.ErrorLogger.Printf("error in CreateAppliance ShouldBindJSON, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "Invalid json provided",
		})
		return
	}
	logger.InfoLogger.Printf("CreateAppliance, req:%v", appReq)
	app := &Appliance{
		SerialNumber: appReq.SerialNumber,
		Brand:        appReq.Brand,
		Model:        appReq.Model,
		Status:       appReq.Status,
		DateBought:   appReq.DateBought,
	}
	p.Repo = &ApplianceModel{
		Model: app,
		DB:    p.DB,
	}

	_, err = p.Repo.Create()
	if err != nil {
		logger.ErrorLogger.Printf("error in creating appliance, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in creating appliance",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     "",
		"debugMsg": "",
	})
	return
}

type ApplianceReq struct {
	ID           string `json:"id"`
	SerialNumber string `json:"serial_number"`
	Brand        string `json:"brand"`
	Model        string `json:"Model"`
	Status       string `json:"status"`
	DateBought   string `json:"date_bought"`
}

func (p *ApplianceProcessor) UpdateAppliance(g *gin.Context) {
	appReq := &ApplianceReq{}
	err := g.ShouldBindJSON(appReq)
	if err != nil {
		logger.ErrorLogger.Printf("error in UpdateAppliance ShouldBindJSON, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "Invalid json provided",
		})
		return
	}
	logger.InfoLogger.Printf("UpdateAppliance, req:%v", appReq)
	id, err := strconv.Atoi(appReq.ID)
	if err != nil {
		logger.ErrorLogger.Printf("error in UpdateAppliance Atoi, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in atoi",
		})
	}

	app := &Appliance{
		SerialNumber: appReq.SerialNumber,
		Brand:        appReq.Brand,
		Model:        appReq.Model,
		Status:       appReq.Status,
		DateBought:   appReq.DateBought,
	}

	p.Repo = &ApplianceModel{
		Model: app,
		DB:    p.DB,
	}

	err = p.Repo.Update(int64(id))
	if err != nil {
		logger.ErrorLogger.Printf("error in updating appliance, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in updating appliance",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     app,
		"debugMsg": "",
	})
	return
}

func (p *ApplianceProcessor) ViewAllAppliance(g *gin.Context) {
	var err error
	app := &Appliance{}
	err = g.ShouldBindJSON(app)
	if err != nil {
		logger.ErrorLogger.Printf("error in ViewAllAppliance ShouldBindJSON, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "Invalid json provided",
		})
		return
	}
	p.Repo = &ApplianceModel{
		Model: app,
		DB:    p.DB,
	}
	appList, err := p.Repo.GetAll()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error finding user from id",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     appList,
		"debugMsg": "",
	})
	logger.InfoLogger.Printf("req:%v, resp:%v", app, appList)
	return
}
