package shorten

import (
	"errors"
	"log"
	"time"

	"github.com/jcoene/go-base62"
	"github.com/url_shortener/database"
	"github.com/url_shortener/pkg/dbclient"
	apiError "github.com/url_shortener/pkg/errors"
	"gorm.io/gorm"
)

type ShortenService struct {
}

func (sc *ShortenService) CreateShortURL(req ShortenRequest) (*ShortenResponse, *apiError.ApiError) {
	log.Println("CreateShortURL service called with request:", req.CustomAlias == "")

	db := dbclient.GetCient()
	now := time.Now()
	expireAt := now.Add(time.Duration(req.Expiry) * time.Minute)

	mapping := database.UrlMapping{
		LongURL:   req.LongURL,
		ExpiresAt: expireAt,
	}

	if req.CustomAlias != "" {
		mapping.ShortKey = "_" + req.CustomAlias
	}

	tx := db.Begin()
	result := tx.Create(&mapping)

	if result.Error != nil {
		tx.Rollback()

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, apiError.ConfilctError("Specified Alias already exist")
		}
		return nil, apiError.InternalServerError("Failed to create short url")
	}

	if req.CustomAlias == "" {
		base62Code := base62.Encode(int64(mapping.ID))
		result := tx.Model(&mapping).Update("short_key", base62Code)

		if result.Error != nil {
			tx.Rollback()
			return nil, apiError.InternalServerError("Failed to create short url")
		}

	}

	tx.Commit()

	return &ShortenResponse{
		ShortURL: mapping.ShortKey,
		LongURL:  mapping.LongURL,
		Expiry:   mapping.ExpiresAt.Format(time.RFC3339),
	}, nil
}
