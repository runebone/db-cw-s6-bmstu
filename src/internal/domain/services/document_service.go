package services

import (
	"github.com/google/uuid"
	r "github.com/runebone/db-cw-s6-bmstu/internal/domain/repositories"
)

type DocumentService struct {
	DocumentRepo *r.DocumentRepository
}

func NewDocumentService(repo *r.DocumentRepository) *DocumentService {
	return &DocumentService{DocumentRepo: repo}
}

func (s *DocumentService) GetDocumentText(documentID uuid.UUID) (string, error) {
	return s.DocumentRepo.GetDocumentText(documentID)
}
