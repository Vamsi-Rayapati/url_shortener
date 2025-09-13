package redirect

import (
	"errors"
	"net/http"
	"time"

	"github.com/url_shortener/database"
	"github.com/url_shortener/pkg/dbclient"
	apiError "github.com/url_shortener/pkg/errors"
	"gorm.io/gorm"
)

type RedirectService struct {
}

func (rs *RedirectService) GetActualUrl(shortKey string) (*string, *apiError.ApiError) {
	db := dbclient.GetCient()
	var mapping database.UrlMapping
	result := db.Where("short_key = ?", shortKey).First(&mapping)

	db.Model(&mapping).Update("count", mapping.Count+1)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apiError.NotFoundError("Url Not found")
		}

		return nil, apiError.InternalServerError("Failed to retrive url")
	}
	now := time.Now()

	if now.After(mapping.ExpiresAt) {

		return nil, &apiError.ApiError{
			Code:    http.StatusGone,
			Message: "Link expired",
		}
	}

	return &mapping.LongURL, nil
}
