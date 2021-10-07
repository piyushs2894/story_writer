package mysql

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"story_writer/src/model"
)

type WordDB interface {
	InsertWord(ctx context.Context, word *model.Word) (int64, error)
	GetWordById(ctx context.Context, wordID int64) (*model.Word, error)
	GetWordsBySentenceId(ctx context.Context, sentenceId int64) ([]model.Word, error)
}

// Empty wordDB declared just to bind below methods with WordDB interface
type wordDB struct{}

func NewWordDB() WordDB {
	return &wordDB{}
}

func (wdb *wordDB) InsertWord(ctx context.Context, word *model.Word) (int64, error) {
	res, err := statements.InsertWord.Exec(word.Word, word.SentenceId)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (wdb *wordDB) GetWordById(ctx context.Context, wordID int64) (*model.Word, error) {
	var word model.Word

	err := statements.GetWordById.Get(&word, wordID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorln("[GetWordById] No rows found. Error: ", err)
			return nil, nil
		}
		log.Errorln("[GetWordById] Error: ", err)
		return nil, err
	}
	return &word, err
}

func (wdb *wordDB) GetWordsBySentenceId(ctx context.Context, sentenceId int64) ([]model.Word, error) {
	var words []model.Word

	err := statements.GetWordsBySentenceId.Select(&words, sentenceId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorln("[GetWordsBySentenceId]", err)
			return nil, nil
		}
		log.Errorln("[GetWordsBySentenceId]", err)
		return nil, err
	}
	return words, err
}
