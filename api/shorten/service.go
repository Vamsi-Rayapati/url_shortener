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
		mapping.ShortKey = req.CustomAlias
	}
	result := db.Create(&mapping)

	if result.Error != nil {
		log.Println("Failed", result.Error)
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, apiError.ConfilctError("Specified Alias already exist")
		}
		return nil, apiError.InternalServerError("Failed to create short url")
	}

	base62Code := base62.Encode(int64(mapping.ID))

	if req.CustomAlias == "" {
		db.Model(&mapping).Update("short_key", base62Code)
	}

	return &ShortenResponse{
		ShortURL: mapping.ShortKey,
		LongURL:  mapping.LongURL,
		Expiry:   mapping.ExpiresAt.Format(time.RFC3339),
	}, nil
}
