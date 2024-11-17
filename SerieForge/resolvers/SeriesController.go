package resolvers

import (
	"SerieForge/connector"
	"SerieForge/entities"
	"SerieForge/graph/model"
	"context"
)

func CreateSeries(ctx context.Context, input model.SeriesInput) (*entities.Series, error) {

	series := &entities.Series{
		Title:        input.Title,
		Rating:       input.Rating,
		SeasonCount:  input.SeasonCount,
		Description:  input.Description,
		EpisodeCount: input.EpisodeCount,
		Status:       input.Status,
		Creators:     input.Creators,
		Cast:         input.Cast,
	}
	if err := connector.Db.Create(&series).Error; err != nil {
		panic(err)
	}
	return series, nil
}

func GetAllSeries() []*entities.Series {
	var series []*entities.Series
	connector.Db.Find(&series)

	return series
}
