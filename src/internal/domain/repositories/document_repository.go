package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
)

type DocumentRepository struct {
	DB *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

func (r *DocumentRepository) CreateDocument(d models.Document) error {
	query := `
		INSERT INTO document
		(id, url, title, lang, orig_doc_id, uploaded_by, upload_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.Exec(query, d.Id, d.Url, d.Title, d.OrigDocId, d.UploadedBy, d.UploadDate)
	return err
}

// XXX there should be a better place for that than document repository
func (r *DocumentRepository) CreateSentence(s models.Sentence) error {
	query := `
		INSERT INTO sentence
		(doc_id, sent_no, content)
		VALUES ($1, $2, $3)
	`
	_, err := r.DB.Exec(query, s.DocId, s.SentNo, s.Content)
	return err
}

func (r *DocumentRepository) CreateToken(t models.Token) error {
	query := `
		INSERT INTO token
		(doc_id, sent_no, token_no, content)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.DB.Exec(query, t.DocId, t.SentNo, t.TokenNo, t.Content)
	return err
}

func (r *DocumentRepository) GetDocumentText(doc_id uuid.UUID) (string, error) {
	query := `
		SELECT content
		FROM sentence
		WHERE doc_id = $1
		ORDER BY sent_no
	`
	rows, err := r.DB.Query(query, doc_id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var text string
	for rows.Next() {
		var sentence string
		if err := rows.Scan(&sentence); err != nil {
			return "", err
		}
		text += sentence + " "
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return text, nil
}
