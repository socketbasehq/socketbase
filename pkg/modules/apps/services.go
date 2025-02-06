package apps

import (
	"github.com/socketbase/socketbase/pkg/db"
	"github.com/socketbase/socketbase/pkg/models"
)

func CreateApp(name string, userID string) (*models.App, error) {
	id, appKey, appSecret, err := generateAppCredentials()

	if err != nil {
		return nil, err
	}

	app := models.App{
		ID:        id,
		Name:      name,
		AppKey:    appKey,
		AppSecret: appSecret,
		UserID:    userID,
	}

	if err := db.DB.Create(&app).Error; err != nil {
		return nil, err
	}

	return &app, nil
}

func GetAppByID(id string) (*models.App, error) {
	var app models.App
	if err := db.DB.Where("id = ?", id).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func ListApps(userID string) ([]models.App, error) {
	var apps []models.App
	if err := db.DB.Where("user_id = ?", userID).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}
