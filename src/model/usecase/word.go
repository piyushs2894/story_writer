package usecase

import (
	"context"
	"log"

	"story_writer/src/model"
	"story_writer/src/model/db/mysql"
)

type WordUseCase interface {
	InsertWord(ctx context.Context, word *model.Word) (int64, error)
	GetWordById(ctx context.Context, wordId int64) (*model.Word, error)
	GetWordsBySentenceId(ctx context.Context, sentenceId int64) ([]model.Word, error)
}

type wordUseCase struct {
	db mysql.WordDB
}

func NewWordUseCase(db mysql.WordDB) WordUseCase {
	return &wordUseCase{
		db: db,
	}
}

func (wuc *wordUseCase) InsertWord(ctx context.Context, word *model.Word) (int64, error) {
	//TODO: Implement jaegar open-tracing
	id, err := wuc.db.InsertWord(ctx, word)
	if err != nil {
		log.Printf("[wordUseCase][InsertWord][Error] : %+v ", err)
		return 0, err
	}

	return id, nil
}

func (wuc *wordUseCase) GetWordById(ctx context.Context, wordId int64) (*model.Word, error) {
	//TODO: Implement jaegar open-tracing
	wordDetail, err := wuc.db.GetWordById(ctx, wordId)
	if err != nil {
		log.Printf("[wordUseCase][GetWordById][Error] : %+v ", err)
		return nil, err
	}

	return wordDetail, err
}

func (wuc *wordUseCase) GetWordsBySentenceId(ctx context.Context, sentenceId int64) ([]model.Word, error) {
	words, err := wuc.db.GetWordsBySentenceId(ctx, sentenceId)
	if err != nil {
		log.Printf("[wordUseCase][GetWordsBySentenceId][Error] : %+v ", err)
		return nil, err
	}

	return words, err
}
