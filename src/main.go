package main

import (
	"github.com/bilalkocoglu/file-service/pkg/config"
	"github.com/bilalkocoglu/file-service/pkg/database"
	"github.com/bilalkocoglu/file-service/pkg/minio"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.ApplicationConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Load config failed")
	}

	router := config.PrepareServer(cfg)
	database.DB, err = gorm.Open(mysql.Open(database.DbURL(database.BuildDBConfig())), &gorm.Config{})
	database.Migration()
	minio.StartMinioConnection(minio.BuildMinioConfig())

	log.Info().Str("addr", cfg.Addr).Msg("starting http listener")
	err = router.Run(cfg.Addr)
	log.Fatal().Err(err).Msg("Server failed")
}
