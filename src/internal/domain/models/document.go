package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id         uuid.UUID
	Url        Url
	Title      string
	Lang       Lang
	OrigDocId  uuid.UUID
	UploadedBy uuid.UUID
	UploadDate time.Time
}

type Sentence struct {
	DocId   uuid.UUID
	SentNo  int
	Content string
}

type Token struct {
	DocId   uuid.UUID
	SentNo  int
	TokenNo int
	Content string
}

type Url string
type Lang string

func NewDocument(url Url, title string, lang Lang, orig_doc_id uuid.UUID, uploaded_by uuid.UUID) *Document {
	d := &Document{
		Id:         uuid.New(),
		Url:        url,
		Title:      title,
		Lang:       lang,
		OrigDocId:  orig_doc_id,
		UploadedBy: uploaded_by,
		UploadDate: time.Now(),
	}
	return d
}

func NewSentence(doc_id uuid.UUID, sent_no int, content string) *Sentence {
	s := &Sentence{
		DocId:   doc_id,
		SentNo:  sent_no,
		Content: content,
	}
	return s
}

func NewToken(doc_id uuid.UUID, sent_no int, token_no int, content string) *Token {
	t := &Token{
		DocId:   doc_id,
		SentNo:  sent_no,
		TokenNo: token_no,
		Content: content,
	}
	return t
}
