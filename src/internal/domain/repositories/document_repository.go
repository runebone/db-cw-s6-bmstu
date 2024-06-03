package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	m "github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
)

type DocumentRepository struct {
	DB *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

func (r *DocumentRepository) CreateDocument(d *m.Document) error {
	query := `
		INSERT INTO document
		(id, url, title, lang, orig_doc_id, uploaded_by, upload_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.Exec(query, d.ID, d.Url, d.Title, d.OrigDocID, d.UploadedBy, d.UploadDate)
	return err
}

// XXX there should be a better place for that than document repository
func (r *DocumentRepository) CreateSentence(s *m.Sentence) error {
	query := `
		INSERT INTO sentence
		(doc_id, sent_no, content)
		VALUES ($1, $2, $3)
	`
	_, err := r.DB.Exec(query, s.DocID, s.SentNo, s.Content)
	return err
}

func (r *DocumentRepository) CreateToken(t *m.Token) error {
	query := `
		INSERT INTO token
		(doc_id, sent_no, token_no, content)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.DB.Exec(query, t.DocID, t.SentNo, t.TokenNo, t.Content)
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

	fmt.Printf("\n\n %s \n\n", text)

	return text, nil
}

func (r *DocumentRepository) GetDocumentsByContent(s string) ([]uuid.UUID, error) {
	query := `
		SELECT doc_id
		FROM sentence
		WHERE content
		LIKE $1
	`
	rows, err := r.DB.Query(query, "%"+s+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docIDs []uuid.UUID
	for rows.Next() {
		var docID uuid.UUID
		if err := rows.Scan(&docID); err != nil {
			return nil, err
		}
		docIDs = append(docIDs, docID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// fmt.Printf("\n\n %s \n\n", docIDs)

	return docIDs, nil
}
