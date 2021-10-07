package usecase

import (
	"story_writer/src/common/config"
	"story_writer/src/common/database"
	"story_writer/src/model/db/mysql"
)

type UseCasesGroup struct {
	StoryUseCase     StoryUseCase
	ParagraphUseCase ParagraphUseCase
	SentenceUseCase  SentenceUseCase
	WordUseCase      WordUseCase
}

var UseCases UseCasesGroup

func Init(dbConnMap map[string]*database.MasterSlave, cfg *config.Config) {
	mysql.Init(dbConnMap)

	SetWordUseCase()
	SetSentenceUseCase()
	SetParagraphsUseCase()
	SetStoryUseCase()
}

func GetUseCases() *UseCasesGroup {
	return &UseCases
}

func SetWordUseCase() {
	UseCases.WordUseCase = NewWordUseCase(mysql.NewWordDB())
}

func SetParagraphsUseCase() {
	UseCases.ParagraphUseCase = NewParagraphUseCase(mysql.NewParagraphDB())
}

func SetSentenceUseCase() {
	UseCases.SentenceUseCase = NewSentenceUseCase(mysql.NewSentenceDB())
}

func SetStoryUseCase() {
	UseCases.StoryUseCase = NewStoryUseCase(mysql.NewStoryDB())
}
