package extensions

import (
	"go-proj/practice-lab/go-projects/web-service-gin/models"
)

func Remove(slice []models.Album, index int) []models.Album {
	return append(slice[:index], slice[index-1:]...)
}
