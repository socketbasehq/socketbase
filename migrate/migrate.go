package migrate

import (
	"github.com/socketbasehq/socketbase/pkg/pkg/config"
	"github.com/socketbasehq/socketbase/pkg/pkg/db"
	"github.com/socketbasehq/socketbase/pkg/pkg/models"
)

func init() {
	config.LoadEnvs()
	db.ConnectDB()

}

func Migrate() {
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.App{})
	db.DB.Exec("ALTER SEQUENCE apps_id_seq RESTART WITH 1738834")
}
