package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"yofio/domain/model"
)

type CreditAssigner interface {
	Assign (investment int32 ) ( int32 , int32 , int32 , error )
}

var InvestmentList = []int32{700, 500, 300}

type Assign struct {

	ID            int   `json:"id"`
	Investment    int32 `json:"investment"`
	Success       bool  `json:"success"`
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`

}

func CreateCredit(investment int32, c *gin.Context) ( int32 , int32 , int32 , error ){

	//db := c.MustGet("db").(*gorm.DB)
	var creditAsigneer CreditAssigner
	var creditEntity model.InvestmentEntity

	creditAsigneer = Assign{}

	credit300, credit500, credit700, err := creditAsigneer.Assign(investment)

	creditEntity.CreditType300 = credit300
	creditEntity.CreditType500 = credit500
	creditEntity.CreditType700 = credit700
	creditEntity.Investment = investment

	if err != nil {
		creditEntity.Success = false
	//	db.Create(&creditEntity)
		return 0,0,0, errors.New("insufficient investment")
	} else {

		creditEntity.Success = true

	}

//	db.Create(&creditEntity)

	return credit300, credit500, credit700, nil

}

func (inv Assign) Assign (investment int32 ) ( int32 , int32 , int32 , error )  {


	if investment%100 != 0 {
		return inv.CreditType300, inv.CreditType500, inv.CreditType700,errors.New("inaccurate or insufficient investment")
	}

	resultDiv := investment / 1500

	resultOp := investment - (resultDiv * 1500)

	inv.CreditType300 = resultDiv
	inv.CreditType500 = resultDiv
	inv.CreditType700 = resultDiv

	if resultOp == 100 || resultOp == 200 || resultOp == 400 {
		return 0,0,0,errors.New("insufficient investment")
	}

	countCredit300, countCredit500, countCredit700 := InvestmentModCero(resultOp, inv.CreditType300, inv.CreditType500, inv.CreditType700)


	return countCredit300, countCredit500, countCredit700,nil

}


func InvestmentModCero(inversionRestante int32, val300 int32, val500 int32, val700 int32) (int32, int32, int32) {

	for _, v := range InvestmentList {

		if inversionRestante%v == 0 {

			switch v {
			case 300:
				val300 = val300 + (inversionRestante / v)
			case 500:
				val500 = val500 + (inversionRestante / v)
			case 700:
				val700 = val700 + (inversionRestante / v)
			}

			return val300, val500, val700
		}

	}


	return validateInvestment(inversionRestante, val300, val500, val700)
}

func validateInvestment(inversionRestante int32, val300 int32, val500 int32, val700 int32) (int32, int32, int32) {
	switch inversionRestante {
	case 800:
		val500 = val500 + 1
		val300 = val300 + 1
	case 1100:
		val500 = val500 + 1
		val300 = val300 + 2
	case 1200:
		val700 = val700 + 1
		val500 = val500 + 1
	case 1300:
		val700 = val700 + 1
		val300 = val300 + 2
	}

	return val300, val500, val700
}
