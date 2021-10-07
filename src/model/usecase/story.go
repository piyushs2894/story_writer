package usecase

import (
	"context"
	"story_writer/src/model"
	"story_writer/src/model/db/mysql"

	log "github.com/sirupsen/logrus"
)

type StoryUseCase interface {
	InsertStory(ctx context.Context, story *model.Story) (int64, error)
	GetStoryInProgress(ctx context.Context) (*model.Story, error)
	UpdateTitle(ctx context.Context, story *model.Story) (int64, error)
	IncrementLength(ctx context.Context, story *model.Story) (int64, error)
	GetStories(ctx context.Context, params model.Params) ([]model.Story, error)
	GetStoryById(ctx context.Context, wordId int64) (*model.StoryResponse, error)
}

type storyUseCase struct {
	db mysql.StoryDB
}

func NewStoryUseCase(db mysql.StoryDB) StoryUseCase {
	return &storyUseCase{
		db: db,
	}
}

func (suc *storyUseCase) InsertStory(ctx context.Context, story *model.Story) (int64, error) {
	id, err := suc.db.InsertStory(ctx, story)
	if err != nil {
		log.Errorln("[storyUseCase][InsertStory][Error] : ", err)
		return 0, err
	}

	return id, nil
}

func (suc *storyUseCase) GetStoryInProgress(ctx context.Context) (*model.Story, error) {
	storyDetail, err := suc.db.GetStoryInProgress(ctx)
	if err != nil {
		log.Errorln("storyUseCase][GetStoryInProgress][Error] : ", err)
		return nil, err
	}

	return storyDetail, err
}

func (suc *storyUseCase) UpdateTitle(ctx context.Context, story *model.Story) (int64, error) {
	id, err := suc.db.UpdateTitle(ctx, story)
	if err != nil {
		log.Errorln("storyUseCase][UpdateTitle][Error] : ", err)
		return 0, err
	}

	return id, err
}

func (suc *storyUseCase) IncrementLength(ctx context.Context, story *model.Story) (int64, error) {
	id, err := suc.db.IncrementLength(ctx, story)
	if err != nil {
		log.Errorln("[storyUseCase][IncrementLength][Error] : ", err)
		return 0, err
	}

	return id, nil
}

func (suc *storyUseCase) GetStories(ctx context.Context, params model.Params) ([]model.Story, error) {
	stories, err := suc.db.GetStories(ctx, params)
	if err != nil {
		log.Errorln("storyUseCase][GetStories][Error] : ", err)
		return nil, err
	}

	return stories, err
}

func (suc *storyUseCase) GetStoryById(ctx context.Context, storyId int64) (*model.StoryResponse, error) {
	storyResp, err := suc.db.GetStoryById(ctx, storyId)
	if err != nil {
		log.Errorln("storyUseCase][GetStoryById][Error] : ", err)
		return nil, err
	}

	return storyResp, err
}
