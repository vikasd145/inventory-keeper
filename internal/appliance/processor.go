package appliance

import (
	"github.com/gin-gonic/gin"
	"github.com/vikasd145/inventory-keeper/pkg/logger"
	"net/http"
	"strconv"
)

func ViewAppliance(g *gin.Context) {
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
	err = app.Get()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error finding user from id",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     app,
		"debugMsg": "",
	})
	return
}

func SearchAppliance(g *gin.Context) {
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
	appList, err := app.Search()
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

func CreateAppliance(g *gin.Context) {
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
	err = app.Create()
	if err != nil {
		logger.ErrorLogger.Printf("error in creating appliance, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in creating appliance",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     app,
		"debugMsg": "",
	})
	return
}

type ApplianceReq struct {
	ID           string `json:"id"`
	SerialNumber string `json:"serial_number"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Status       string `json:"status"`
	DateBought   string `json:"date_bought"`
}

func UpdateAppliance(g *gin.Context) {
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

	err = app.Update(int64(id))
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

func ViewAllAppliance(g *gin.Context) {
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
	appList, err := app.GetAll()
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
