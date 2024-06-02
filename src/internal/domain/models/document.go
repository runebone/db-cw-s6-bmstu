package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID         uuid.UUID
	Url        Url
	Title      string
	Lang       Lang
	OrigDocID  uuid.UUID
	UploadedBy uuid.UUID
	UploadDate time.Time
}

type Sentence struct {
	DocID   uuid.UUID
	SentNo  int
	Content string
}

type Token struct {
	DocID   uuid.UUID
	SentNo  int
	TokenNo int
	Content string
}

type Url string
type Lang string

func NewDocument(url Url, title string, lang Lang, origDocID uuid.UUID, uploadedBy uuid.UUID) *Document {
	d := &Document{
		ID:         uuid.New(),
		Url:        url,
		Title:      title,
		Lang:       lang,
		OrigDocID:  origDocID,
		UploadedBy: uploadedBy,
		UploadDate: time.Now(),
	}
	return d
}

func NewSentence(docID uuid.UUID, sentNo int, content string) *Sentence {
	s := &Sentence{
		DocID:   docID,
		SentNo:  sentNo,
		Content: content,
	}
	return s
}

func NewToken(docID uuid.UUID, sentNo int, tokenNo int, content string) *Token {
	t := &Token{
		DocID:   docID,
		SentNo:  sentNo,
		TokenNo: tokenNo,
		Content: content,
	}
	return t
}
