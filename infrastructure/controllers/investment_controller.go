package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"yofio/domain/service"
	"yofio/infrastructure/adapters/modelRequest"
)

type assignResponse struct {

	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`

}

type Stats struct {
	TotalAssigRealizadas    int `json:"total_asignments_done"`
	TotalAssigSuccess int `json:"total_asignments_success"`
	TotalAssigFail    int `json:"total_asignments_fail"`
	AvgInvestmentSuccess float64 `json:"avg_investment_success"`
	AvgInvestmentFail       float64 `json:"avg_investment_fail"`
}

func  Create(c *gin.Context) {

	var assigner modelRequest.Assigner
	var response assignResponse
	if err:= c.ShouldBindJSON(&assigner); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
		return
	}

	Credit300,Credit500,Credit700,err := service.CreateCredit(assigner.Investment, c)

	if err != nil {
		c.JSON(http.StatusBadRequest,   "bad request"  )
		return
	} else {

		response.CreditType300 = Credit300
		response.CreditType500 = Credit500
		response.CreditType700 = Credit700
	}

	c.JSON(http.StatusCreated, response)
}




func  GetStatf(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var stats Stats

	db.Raw("select count(*) from investment_entities" ).Row().Scan(&stats.TotalAssigRealizadas)
	db.Raw("select count(*) from investment_entities where success = true" ).Row().Scan(&stats.TotalAssigSuccess)
	db.Raw("select count(*) from investment_entities where success = false" ).Row().Scan(&stats.TotalAssigFail)
	db.Raw("select avg(investment)::numeric(12, 1) from investment_entities where success = true" ).Row().Scan(&stats.AvgInvestmentSuccess)
	db.Raw("select avg(investment)::numeric(12, 1) from investment_entities where success = false" ).Row().Scan(&stats.AvgInvestmentFail)

	c.JSON(http.StatusCreated,   stats  )
}
