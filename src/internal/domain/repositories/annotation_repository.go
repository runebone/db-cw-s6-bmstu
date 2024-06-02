package repositories

import (
	"database/sql"

	"github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
)

type AnnotationRepository struct {
	DB *sql.DB
}

func NewAnnotationRepository(db *sql.DB) *AnnotationRepository {
	return &AnnotationRepository{DB: db}
}

func (r *AnnotationRepository) CreateAnnotationTask(at models.AnnotationTask) error {
	query := `
		INSERT INTO annotation_task
		(id, orig_doc_id, trans_doc_id, description, created_by, creation_date, last_update_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.Exec(query, at.ID, at.OrigDocID, at.TransDocID, at.Description, at.CreatedBy, at.CreationDate, at.LastUpdateDate)
	return err
}

func (r *AnnotationRepository) CreateStructAnnotation(sa models.StructAnnotation) error {
	query := `
		INSERT INTO struct_annotation
		(id, task_id, orig_doc_id, beg_sent_no, end_sent_no, status, done_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.Exec(query, sa.ID, sa.TaskID, sa.OrigDocID, sa.BegSentNo, sa.EndSentNo, sa.Status, sa.DoneBy)
	return err
}

func (r *AnnotationRepository) CreateTermAnnotation(ta models.TermAnnotation) error {
	query := `
		INSERT INTO term_annotation
		(id, task_id, orig_doc_id, trans_doc_id, status, done_by)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.Exec(query, ta.ID, ta.TaskID, ta.OrigDocID, ta.TransDocID, ta.Status, ta.DoneBy)
	return err
}

func (r *AnnotationRepository) CreateTermAnnotationPart(chicken models.TermAnnotationPart) error {
	query := `
		INSERT INTO term_annotation_part
		(annot_id, part_no, orig_sent_no, trans_sent_no, beg_orig_token_no, end_orig_token_no, beg_trans_token_no, end_trans_token_no)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.DB.Exec(query, chicken.AnnotID, chicken.PartNo, chicken.OrigSentNo, chicken.TransSentNo, chicken.BegOrigTokenNo, chicken.EndOrigTokenNo, chicken.BegTransTokenNo, chicken.EndTransTokenNo)
	return err
}
