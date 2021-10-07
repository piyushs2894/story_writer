package usecase

import (
	"context"
	"story_writer/src/model"
	"story_writer/src/model/db/mysql"

	log "github.com/sirupsen/logrus"
)

type SentenceUseCase interface {
	InsertSentence(ctx context.Context, sentence *model.Sentence) (int64, error)
	IncrementLength(ctx context.Context, sentence *model.Sentence) (int64, error)
	GetSentenceInProgress(ctx context.Context) (*model.Sentence, error)
	GetSentenceParaIdsByStoryId(ctx context.Context, storyId int64) ([]model.SentenceParaId, error)
}

type sentenceUseCase struct {
	db mysql.SentenceDB
}

func NewSentenceUseCase(db mysql.SentenceDB) SentenceUseCase {
	return &sentenceUseCase{
		db: db,
	}
}

func (suc *sentenceUseCase) InsertSentence(ctx context.Context, sentence *model.Sentence) (int64, error) {
	//TODO: Implement jaegar open-tracing
	id, err := suc.db.InsertSentence(ctx, sentence)
	if err != nil {
		log.Errorln("[sentenceUseCase][InsertSentence][Error] : ", err)
		return 0, err
	}

	return id, nil
}

func (suc *sentenceUseCase) IncrementLength(ctx context.Context, sentence *model.Sentence) (int64, error) {
	id, err := suc.db.IncrementLength(ctx, sentence)
	if err != nil {
		log.Errorln("[sentenceUseCase][IncrementLength][Error] : ", err)
		return 0, err
	}

	return id, nil
}

func (suc *sentenceUseCase) GetSentenceInProgress(ctx context.Context) (*model.Sentence, error) {
	sentenceDetail, err := suc.db.GetSentenceInProgress(ctx)
	if err != nil {
		log.Errorln("sentenceUseCase][GetSentenceInProgress][Error] : ", err)
		return nil, err
	}

	return sentenceDetail, err
}

func (suc *sentenceUseCase) GetSentenceParaIdsByStoryId(ctx context.Context, storyId int64) ([]model.SentenceParaId, error) {
	sentenceParaIds, err := suc.db.GetSentenceParaIdsByStoryId(ctx, storyId)
	if err != nil {
		log.Errorln("sentenceUseCase][GetSentenceParaIdsByStoryId][Error] : ", err)
		return nil, err
	}

	return sentenceParaIds, err
}
