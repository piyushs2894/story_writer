package mysql

import (
	"log"
	"story_writer/src/common/database"
	"story_writer/src/constant"

	"github.com/jmoiron/sqlx"
)

var db database.MasterSlave

type PreparedStatements struct {
	InsertWord                  *sqlx.Stmt
	GetWordById                 *sqlx.Stmt
	GetStoryById                *sqlx.Stmt
	InsertParagraph             *sqlx.Stmt
	InsertStory                 *sqlx.Stmt
	UpdateTitle                 *sqlx.Stmt
	InsertSentence              *sqlx.Stmt
	GetStoryInProgress          *sqlx.Stmt
	GetParagraphInProgress      *sqlx.Stmt
	GetSentenceInProgress       *sqlx.Stmt
	GetWordsBySentenceId        *sqlx.Stmt
	IncrementSentenceLength     *sqlx.Stmt
	IncrementParagraphLength    *sqlx.Stmt
	IncrementStoryLength        *sqlx.Stmt
	GetParagraphsByStoryId      *sqlx.Stmt
	GetSentenceParaIdsByStoryId *sqlx.Stmt
}

var (
	statements PreparedStatements
)

func Init(dbConnMap map[string]*database.MasterSlave) {
	db = *dbConnMap[constant.DB_STORY_WRITER]

	log.Println("Initiating statements...")

	statements = PreparedStatements{
		InsertWord:                  db.Master.Preparex(insertWordQuery),
		InsertParagraph:             db.Master.Preparex(insertParagraphQuery),
		GetWordById:                 db.Slave.Preparex(getWordById),
		GetStoryById:                db.Slave.Preparex(getStoryById),
		InsertStory:                 db.Master.Preparex(insertStoryQuery),
		UpdateTitle:                 db.Master.Preparex(updateTitleQuery),
		InsertSentence:              db.Master.Preparex(insertSentenceQuery),
		GetStoryInProgress:          db.Slave.Preparex(getStoryInProgress),
		GetParagraphInProgress:      db.Slave.Preparex(getParagraphInProgress),
		GetSentenceInProgress:       db.Slave.Preparex(getSentenceInProgress),
		GetWordsBySentenceId:        db.Slave.Preparex(getWordsBySentenceId),
		IncrementSentenceLength:     db.Master.Preparex(incrementSentenceLength),
		IncrementParagraphLength:    db.Master.Preparex(incrementParagraphLength),
		IncrementStoryLength:        db.Master.Preparex(incrementStoryLength),
		GetParagraphsByStoryId:      db.Slave.Preparex(getParagraphsByStoryId),
		GetSentenceParaIdsByStoryId: db.Slave.Preparex(getSentenceParaIdsByStoryId),
	}

	log.Println("Initiated statements...")
}

var (
	insertWordQuery             = `INSERT INTO words (word, sentence_id) VALUES (?, ?)`
	getWordById                 = `SELECT id, sentence_id from words where id = ?`
	getStoryById                = `SELECT id, title, created_at, updated_at from stories where id = ?`
	insertParagraphQuery        = `INSERT INTO paragraphs (story_id, length, status) VALUES (?, ?, ?)`
	insertStoryQuery            = `INSERT INTO stories (title, title_length, length, status) VALUES (?, ?, ?, ?)`
	updateTitleQuery            = `Update stories set title = ?, title_length = ?, status = ? where id = ?`
	insertSentenceQuery         = `INSERT INTO sentences (paragraph_id, length, status) VALUES (?, ?, ?)`
	getStoryInProgress          = `SELECT id, title, title_length, length, status, created_at, updated_at from stories where status <= ?`
	getParagraphInProgress      = `SELECT id, story_id, length, status from paragraphs where status = ?`
	getSentenceInProgress       = `SELECT id, paragraph_id, length, status from sentences where status = ?`
	getWordsBySentenceId        = `SELECT id, word, sentence_id from words where sentence_id  = ?`
	incrementSentenceLength     = `Update sentences set length = ?, status = ? where id = ?`
	incrementParagraphLength    = `Update paragraphs set length = ?, status = ? where id = ?`
	incrementStoryLength        = `Update stories set length = ?, status = ? where id = ?`
	getParagraphsByStoryId      = `SELECT id from paragraphs where story_id = ?`
	getSentenceParaIdsByStoryId = `SELECT sentences.id as sentence_id, paragraphs.id as paragraph_id from sentences inner join paragraphs on sentences.paragraph_id = paragraphs.id where paragraphs.story_id = ?`
)
