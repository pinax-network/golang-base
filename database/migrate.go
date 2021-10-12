package database

import (
	"fmt"
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func Migrate(clusterConfig *ClusterConfig, config *MigrationConfig) error {

	connection := &MysqlConnectionOptions{
		User:     clusterConfig.User,
		Password: clusterConfig.Password,
		Database: clusterConfig.Database,
		Host:     clusterConfig.Connections[0].Host,
		Port:     clusterConfig.Connections[0].Port,
	}

	dir := fmt.Sprintf("file://%s", config.MigrationDir)
	dsn := fmt.Sprintf("mysql://%s", GetMysqlDsn(connection, true))

	migration, err := migrate.New(dir, dsn)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	defer func() {
		sErr, dErr := migration.Close()
		log.CriticalIfError("failed to close source connection after migration", sErr, zap.String("source_dir", config.MigrationDir))
		log.CriticalIfError("failed to close database connection after migration", dErr, zap.Any("db_connection", connection))
	}()

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("an error occurred while syncing the database: %v", err)
	}

	return nil
}
