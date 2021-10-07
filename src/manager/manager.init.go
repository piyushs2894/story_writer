package manager

import (
	"context"

	"story_writer/src/common/config"
	"story_writer/src/model"
	"story_writer/src/model/usecase"
)

type Manager interface {
	AddWord(ctx context.Context, word string) (*model.AddWordResponse, error)
	GetStories(ctx context.Context, params model.Params) (model.StoriesResponse, error)
	GetStoryById(ctx context.Context, wordId int64) (*model.StoryResponse, error)
}

type Module struct {
	cfg      *config.Config
	useCases *usecase.UseCasesGroup
}

func New() *Module {
	return &Module{useCases: usecase.GetUseCases(), cfg: config.GetConfig()}
}
