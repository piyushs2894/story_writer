package model

import (
	"time"
)

type State int

const (
	TitleInProgress State = iota + 1
	TitleCompleted
	InProgress
	Completed
)

type Story struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Length      int       `db:"length" json:"length,omitempty"`             // indicates length of paragraphs also in story
	TitleLength int       `db:"title_length" json:"title_length,omitempty"` // indicates length of words in title story
	Status      State     `db:"status" json:"status,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Paragraph struct {
	ID        int64     `db:"id" json:"id"`
	StoryId   int64     `db:"story_id" json:"story_id"`
	Length    int       `db:"length" json:"length"` // indicates length of sentences in paragraph
	Status    State     `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Sentence struct {
	ID          int64     `db:"id" json:"id"`
	ParagraphId int64     `db:"paragraph_id" json:"paragraph_id"`
	Length      int       `db:"length" json:"length"` // indicates length of words in sentence
	Status      State     `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Word struct {
	ID         int64     `db:"id" json:"id"`
	Word       string    `db:"word" json:"word"`
	SentenceId int64     `db:"sentence_id" json:"sentence_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type AddWordResponse struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	CurrentSentence string `json:"current_sentence"`
}

type Params struct {
	Limit  int
	Offset int
	SortBy string
	Order  string
}

type StoriesResponse struct {
	Limit   int     `json:"limit"`
	Offset  int     `json:"offset"`
	Count   int     `json:"count"`
	Results []Story `json:"results"`
}

type StoryResponse struct {
	Id         int64               `json:"id"`
	Title      string              `json:"title"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	Paragraphs []ParagraphResponse `json:"paragraphs"`
}

type ParagraphResponse struct {
	Sentences []string `json:"sentences"`
}

type SentenceParaId struct {
	SentenceId  int64 `db:"sentence_id" json:"sentence_id"`
	ParagraphId int64 `db:"paragraph_id" json:"paragraph_id"`
}
