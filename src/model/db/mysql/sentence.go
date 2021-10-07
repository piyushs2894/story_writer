package mysql

import (
	"context"
	"story_writer/src/constant"
	"story_writer/src/model"

	log "github.com/sirupsen/logrus"
)

type SentenceDB interface {
	InsertSentence(ctx context.Context, statement *model.Sentence) (int64, error)
	IncrementLength(ctx context.Context, statement *model.Sentence) (int64, error)
	GetSentenceInProgress(ctx context.Context) (*model.Sentence, error)
	GetSentenceParaIdsByStoryId(ctx context.Context, storyId int64) ([]model.SentenceParaId, error)
}

// Empty sentenceDB declared just to bind below methods with statementDB interface
type sentenceDB struct{}

func NewSentenceDB() SentenceDB {
	return &sentenceDB{}
}

func (sdb *sentenceDB) InsertSentence(ctx context.Context, sentence *model.Sentence) (int64, error) {
	res, err := statements.InsertSentence.Exec(sentence.ParagraphId, sentence.Length, sentence.Status)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (sdb *sentenceDB) IncrementLength(ctx context.Context, sentence *model.Sentence) (int64, error) {
	if sentence.Length == constant.MAX_WORDS_LENGTH {
		sentence.Status = model.Completed
	}

	_, err := statements.IncrementSentenceLength.Exec(sentence.Length, sentence.Status, sentence.ID)
	if err != nil {
		return 0, err
	}

	return sentence.ID, nil
}

func (sdb *sentenceDB) GetSentenceInProgress(ctx context.Context) (*model.Sentence, error) {
	var sentences []model.Sentence

	err := statements.GetSentenceInProgress.Select(&sentences, model.InProgress)
	if err != nil {
		log.Errorln("[GetSentenceInProgress]", err)
		return nil, err
	}
	if len(sentences) == 0 {
		return nil, nil
	}

	return &sentences[0], err
}

func (sbd *sentenceDB) GetSentenceParaIdsByStoryId(ctx context.Context, storyId int64) ([]model.SentenceParaId, error) {
	var sentenceParaIds []model.SentenceParaId

	err := statements.GetSentenceParaIdsByStoryId.Select(&sentenceParaIds, storyId)
	if err != nil {
		log.Errorln("[GetSentenceParaIdsByStoryId] Error: ", err)
		return nil, err
	}
	if len(sentenceParaIds) == 0 {
		return nil, nil
	}

	return sentenceParaIds, err
}
