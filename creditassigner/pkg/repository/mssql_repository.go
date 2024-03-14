package repository

import (
	"creditassigner/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

var dbConnectionString string

type MSSQL struct {
}

func NewMSSQLRepository() IRepository {

	return MSSQL{}
}

type StoreProcedureResult struct {
	Error string `json:"errors"`
	Ok    string `json:"ok"`
}

type Cmd80EventsMSSQL struct {
	Accion   string `json:"accion"`
	EventoId int    `json:"evento_id"`
}

func (mssql MSSQL) openDatabaseConnection() (*sql.DB, error) {
	LoadEnv()
	return sql.Open("mssql", dbConnectionString)
}

func (mssql MSSQL) closeDatabaseConnection(databaseConnection *sql.DB) error {
	return databaseConnection.Close()
}

func (mssql MSSQL) SaveStatistics(creditStatus models.CreditStatus) error {

	databaseConnection, err := mssql.openDatabaseConnection()
	if err != nil {
		return errors.New("No se pudo crear la conexion a MSSQL " + err.Error())
	}
	defer mssql.closeDatabaseConnection(databaseConnection)

	_, err = databaseConnection.Query("EXEC insert_credit_assigment_statistics @assigment_succesful = ?;",
		creditStatus.AssigmentSuccessful)

	if err != nil {
		return errors.New("No se pudo ejecutar el store procedure insert_credit_assigment_statistics: " + err.Error())

	}

	return nil
}

func (mssql MSSQL) GetStatistics() (models.CreditStatistics, error) {
	databaseConnection, err := mssql.openDatabaseConnection()
	if err != nil {
		return models.CreditStatistics{}, errors.New("No se pudo crear la conexion a MSSQL " + err.Error())
	}
	defer mssql.closeDatabaseConnection(databaseConnection)

	result, err := databaseConnection.Query("SELECT  assignments_made, successful_assignments, failed_assignments, average_successful_assignments, average_failed_assignments FROM credit_statistics;")

	if err != nil {
		return models.CreditStatistics{}, errors.New("No se pudo obtener las estadisticas : " + err.Error())
	}

	if !result.Next() {
		return models.CreditStatistics{}, errors.New("No se pudo obtener las estadisticas : " + err.Error())
	}

	var creditStatistics models.CreditStatistics
	result.Scan(&creditStatistics.AssigmentsMade, &creditStatistics.SuccessfulAssignments, &creditStatistics.FailedAssignments, &creditStatistics.AverageSuccessfulAssignments, &creditStatistics.AverageFailedAssignments)

	return creditStatistics, nil
}

func LoadEnv() {
	dbConnectionString = os.Getenv("DB_STRING")
	fmt.Println(dbConnectionString)
	if dbConnectionString == "" {
		panic("No hay cadena de conexion a base de datos")
	}
}
