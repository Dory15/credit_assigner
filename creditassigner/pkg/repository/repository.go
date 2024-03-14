package repository

import (
	"creditassigner/pkg/models"
)

type IRepository interface {
	SaveStatistics(models.CreditStatus) error
	GetStatistics() (models.CreditStatistics, error)
}

func NewRepository() IRepository {
	return NewMSSQLRepository()
}
