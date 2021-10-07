package usecase

import (
	"context"
	"story_writer/src/model"
	"story_writer/src/model/db/mysql"

	log "github.com/sirupsen/logrus"
)

type ParagraphUseCase interface {
	InsertParagraph(ctx context.Context, paragraph *model.Paragraph) (int64, error)
	IncrementLength(ctx context.Context, paragraph *model.Paragraph) (int64, error)
	GetParagraphInProgress(ctx context.Context) (*model.Paragraph, error)
	GetParagraphsByStoryId(ctx context.Context, storyId int64) ([]model.Paragraph, error)
}

type paragraphUseCase struct {
	db mysql.ParagraphDB
}

func NewParagraphUseCase(db mysql.ParagraphDB) ParagraphUseCase {
	return &paragraphUseCase{
		db: db,
	}
}

func (puc *paragraphUseCase) InsertParagraph(ctx context.Context, paragraph *model.Paragraph) (int64, error) {
	//TODO: Implement jaegar open-tracing
	id, err := puc.db.InsertParagraph(ctx, paragraph)
	if err != nil {
		log.Errorln("[paragraphUseCase][InsertParagraph][Error] : ", err)
		return 0, err
	}

	return id, nil
}

func (puc *paragraphUseCase) IncrementLength(ctx context.Context, paragraph *model.Paragraph) (int64, error) {
	id, err := puc.db.IncrementLength(ctx, paragraph)
	if err != nil {
		log.Errorln("[paragraphUseCase][IncrementLength][Error] : ", err)
		return 0, err
	}

	return id, nil
}

func (puc *paragraphUseCase) GetParagraphInProgress(ctx context.Context) (*model.Paragraph, error) {
	paraDetail, err := puc.db.GetParagraphInProgress(ctx)
	if err != nil {
		log.Errorln("paragraphUseCase][GetParagraphInProgress][Error] : ", err)
		return nil, err
	}

	return paraDetail, err
}

func (puc *paragraphUseCase) GetParagraphsByStoryId(ctx context.Context, storyId int64) ([]model.Paragraph, error) {
	paraResp, err := puc.db.GetParagraphsByStoryId(ctx, storyId)
	if err != nil {
		log.Errorln("paragraphUseCase][GetParagraphsByStoryId][Error] : ", err)
		return nil, err
	}

	return paraResp, err
}
