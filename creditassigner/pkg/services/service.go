package services

import (
	"creditassigner/pkg/core"
	"creditassigner/pkg/logger"
	"creditassigner/pkg/models"
	repositoryLib "creditassigner/pkg/repository"
)

var fileName string = "service"
var serviceLogger logger.ILogger

type Service struct {
	repository repositoryLib.IRepository
}

func init() {
	serviceLogger = logger.NewLoggerInstace(fileName)
}

func NewService() Service {
	return Service{repository: repositoryLib.NewRepository()}
}

func (service Service) CreditAssignerService(investment models.Investment) (models.Credits, error) {
	creditAssignerCore := core.CreditAssigner{}
	threeHundred, fiveHundred, sevenHundred, err := creditAssignerCore.Assign(investment.Investment)

	if err != nil {
		serviceLogger.Error(err.Error())
		go func() {
			service.repository.SaveStatistics(models.CreditStatus{AssigmentSuccessful: false})
		}()
		return models.Credits{}, err
	}

	go func() {
		service.repository.SaveStatistics(models.CreditStatus{AssigmentSuccessful: true})
	}()

	return models.Credits{
		ThreeHundred: threeHundred,
		FiveHundred:  fiveHundred,
		SevenHundred: sevenHundred,
	}, nil
}

func (service Service) GetStatisticsService() (models.CreditStatistics, error) {
	creditStatistics, err := service.repository.GetStatistics()
	if err != nil {
		serviceLogger.Error(err.Error())
		return models.CreditStatistics{}, err
	}
	return creditStatistics, nil
}
