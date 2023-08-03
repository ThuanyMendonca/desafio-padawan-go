package dependency

import (
	"github.com/ThuanyMendonca/desafio-padawan-go/config/env"
	db "github.com/ThuanyMendonca/desafio-padawan-go/gormdb"
	dbInit "github.com/ThuanyMendonca/desafio-padawan-go/internal/infra"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/domain"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/repository"
)

var (
	ConverterRepository domain.IConverterRepository
)

func Load() error {
	// Db
	projectDb, err := dbInit.InitMySqlDb(env.DbConnection.Host, env.DbConnection.Port, env.DbConnection.User, env.DbConnection.Name, env.DbConnection.Password, env.DbConnection.TimeZone)
	if err != nil {
		return err
	}

	db.Load(projectDb)

	// Repositories
	ConverterRepository = repository.NewConverterRepository(projectDb)

	return nil
}
