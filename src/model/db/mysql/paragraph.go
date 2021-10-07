package mysql

import (
	"context"
	"log"
	"story_writer/src/constant"
	"story_writer/src/model"
)

type ParagraphDB interface {
	InsertParagraph(ctx context.Context, paragraph *model.Paragraph) (int64, error)
	IncrementLength(ctx context.Context, paragraph *model.Paragraph) (int64, error)
	GetParagraphInProgress(ctx context.Context) (*model.Paragraph, error)
	GetParagraphsByStoryId(ctx context.Context, storyId int64) ([]model.Paragraph, error)
}

// Empty paragraphDB declared just to bind below methods with paragraphDB interface
type paragraphDB struct{}

func NewParagraphDB() ParagraphDB {
	return &paragraphDB{}
}

func (pdb *paragraphDB) InsertParagraph(ctx context.Context, para *model.Paragraph) (int64, error) {
	res, err := statements.InsertParagraph.Exec(para.StoryId, para.Length, para.Status)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (pdb *paragraphDB) IncrementLength(ctx context.Context, paragraph *model.Paragraph) (int64, error) {
	if paragraph.Length == constant.MAX_SENTENCE_LENGTH {
		paragraph.Status = model.Completed
	}

	_, err := statements.IncrementParagraphLength.Exec(paragraph.Length, paragraph.Status, paragraph.ID)
	if err != nil {
		return 0, err
	}

	return paragraph.ID, nil
}

func (pdb *paragraphDB) GetParagraphInProgress(ctx context.Context) (*model.Paragraph, error) {
	var paragraphs []model.Paragraph

	err := statements.GetParagraphInProgress.Select(&paragraphs, model.InProgress)
	if err != nil {
		log.Println("[GetParagraphInProgress]", err)
		return nil, err
	}

	if len(paragraphs) == 0 {
		return nil, nil
	}

	return &paragraphs[0], err
}

func (pdb *paragraphDB) GetParagraphsByStoryId(ctx context.Context, storyId int64) ([]model.Paragraph, error) {
	var paragraphs []model.Paragraph

	err := statements.GetParagraphsByStoryId.Select(&paragraphs, model.InProgress)
	if err != nil {
		log.Println("[GetParagraphsByStoryId] Error: ", err)
		return nil, err
	}

	if len(paragraphs) == 0 {
		return nil, nil
	}

	return paragraphs, err
}
