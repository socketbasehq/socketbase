package socketbase

import (
	"github.com/socketbasehq/socketbase/pkg/db"
	"github.com/socketbasehq/socketbase/pkg/models"
)

func GetApp(id string) (*models.App, error) {
	var app models.App
	if err := db.DB.Where("id = ?", id).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}
