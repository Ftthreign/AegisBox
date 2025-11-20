package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Ftthreign/aegisbox/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	DataSourceName string
	LogLevel       logger.LogLevel
}

func BuildDataSourceName(dbCfg config.DatabaseConfig) string {
	if dbCfg.DatabaseURL != "" {
		return dbCfg.DatabaseURL
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbCfg.DBHost,
		dbCfg.DBPort,
		dbCfg.DBPassword,
		dbCfg.DBUser,
		dbCfg.DBName,
	)
}

func CheckLogLevel(cfg *config.Config) PostgresConfig {
	logLevel := logger.Info
	if cfg.ServerConfig.Env == "production" {
		logLevel = logger.Error
	}

	return PostgresConfig{
		DataSourceName: BuildDataSourceName(cfg.DatabaseConfig),
		LogLevel:       logLevel,
	}
}

func Connect(pg PostgresConfig) (*gorm.DB, error) {
	var (
		db *gorm.DB
		err error
		maxRetries = 5
	)

	gormLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: pg.LogLevel,
			Colorful: true,
		},
	)

	for i := 1; i <= maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(pg.DataSourceName), &gorm.Config{
			Logger: gormLogger,
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		})

		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database after %d attempts: %w", maxRetries, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}