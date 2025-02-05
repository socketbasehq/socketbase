package migrate

import (
	"github.com/socketbase/socketbase/internal/config"
	"github.com/socketbase/socketbase/internal/db"
	"github.com/socketbase/socketbase/internal/models"
)

func init() {
	config.LoadEnvs()
	db.ConnectDB()

}

func Migrate() {
	db.DB.AutoMigrate(&models.User{})
}
