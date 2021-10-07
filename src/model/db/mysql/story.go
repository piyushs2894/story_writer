package mysql

import (
	"context"
	"fmt"
	"log"
	"story_writer/src/constant"
	"story_writer/src/model"
)

type StoryDB interface {
	InsertStory(ctx context.Context, story *model.Story) (int64, error)
	UpdateTitle(ctx context.Context, story *model.Story) (int64, error)
	GetStoryInProgress(ctx context.Context) (*model.Story, error)
	IncrementLength(ctx context.Context, story *model.Story) (int64, error)
	GetStories(ctx context.Context, params model.Params) ([]model.Story, error)
	GetStoryById(ctx context.Context, storyId int64) (*model.StoryResponse, error)
}

// Empty storyDB declared just to bind below methods with storyDB interface
type storyDB struct{}

func NewStoryDB() StoryDB {
	return &storyDB{}
}

func (sdb *storyDB) InsertStory(ctx context.Context, story *model.Story) (int64, error) {
	res, err := statements.InsertStory.Exec(story.Title, story.TitleLength, story.Length, model.TitleInProgress)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (sdb *storyDB) UpdateTitle(ctx context.Context, story *model.Story) (int64, error) {
	_, err := statements.UpdateTitle.Exec(story.Title, story.TitleLength, story.Status, story.ID)
	if err != nil {
		return 0, err
	}

	return story.ID, nil
}

func (sdb *storyDB) IncrementLength(ctx context.Context, story *model.Story) (int64, error) {
	if story.Length == constant.MAX_PARAGRAPHS_LENGTH {
		story.Status = model.Completed
	}

	_, err := statements.IncrementStoryLength.Exec(story.Length, story.Status, story.ID)
	if err != nil {
		return 0, err
	}

	return story.ID, nil
}

func (sdb *storyDB) GetStoryInProgress(ctx context.Context) (*model.Story, error) {
	var stories []model.Story
	err := statements.GetStoryInProgress.Select(&stories, model.InProgress)
	if err != nil {
		log.Println("[GetStoryInProgress] Error: ", err)
		return nil, err
	}
	if len(stories) == 0 {
		return nil, nil
	}

	return &(stories[0]), err
}

func (sdb *storyDB) GetStories(ctx context.Context, params model.Params) ([]model.Story, error) {
	var stories []model.Story

	query := "Select id, title, created_at, updated_at from stories"

	whereClause := ""
	if len(params.SortBy) > 0 {
		whereClause = fmt.Sprintf("%s order by %s %s", whereClause, params.SortBy, params.Order)
	}

	whereClause = fmt.Sprintf("%s limit %d offset %d", whereClause, params.Limit, params.Offset)
	query = fmt.Sprintf("%s %s", query, whereClause)

	execStmt := db.Slave.Preparex(query)
	defer execStmt.Close()

	if err := execStmt.Select(&stories); err != nil {
		log.Println("GetStories SQLX Error:", err)
		return stories, err
	}

	return stories, nil
}

func (sdb *storyDB) GetStoryById(ctx context.Context, storyId int64) (*model.StoryResponse, error) {
	var stories []model.Story
	var storyResp model.StoryResponse

	err := statements.GetStoryById.Select(&stories, storyId)
	if err != nil {
		log.Println("[GetStoryById] Error: ", err)
		return nil, err
	}
	if len(stories) == 0 {
		return nil, nil
	}

	storyResp = model.StoryResponse{Id: stories[0].ID, Title: stories[0].Title, CreatedAt: stories[0].CreatedAt, UpdatedAt: stories[0].UpdatedAt}
	return &(storyResp), err
}
