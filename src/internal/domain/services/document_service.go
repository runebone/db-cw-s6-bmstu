package services

import (
	"github.com/google/uuid"
	m "github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
	r "github.com/runebone/db-cw-s6-bmstu/internal/domain/repositories"
)

type DocumentService struct {
	DocumentRepo *r.DocumentRepository
}

func NewDocumentService(repo *r.DocumentRepository) *DocumentService {
	return &DocumentService{DocumentRepo: repo}
}

func (s *DocumentService) CreateDocument(d *m.Document) error {
	return s.DocumentRepo.CreateDocument(d)
}

func (s *DocumentService) CreateSentence(snt *m.Sentence) error {
	return s.DocumentRepo.CreateSentence(snt)
}

func (s *DocumentService) GetDocumentText(documentID uuid.UUID) (string, error) {
	return s.DocumentRepo.GetDocumentText(documentID)
}

func (s *DocumentService) GetDocumentsByContent(str string) ([]uuid.UUID, error) {
	return s.DocumentRepo.GetDocumentsByContent(str)
}
