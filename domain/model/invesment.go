package model

type InvestmentEntity struct {
	ID            int
	Investment    int32
	Success       bool
	CreditType300 int32
	CreditType500 int32
	CreditType700 int32
}