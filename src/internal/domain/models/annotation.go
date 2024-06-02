package models

import (
	"time"

	"github.com/google/uuid"
)

type AnnotationTask struct {
	Id             uuid.UUID
	OrigDocId      uuid.UUID
	TransDocId     uuid.UUID
	Description    string
	CreatedBy      uuid.UUID
	CreationDate   time.Time
	LastUpdateDate time.Time
}

type StructAnnotation struct {
	Id        uuid.UUID
	TaskId    uuid.UUID
	OrigDocId uuid.UUID
	BegSentNo int
	EndSentNo int
	Status    Status
	DoneBy    uuid.UUID
}

type Status string

type TermAnnotation struct {
	Id         uuid.UUID
	TaskId     uuid.UUID
	OrigDocId  uuid.UUID
	TransDocId uuid.UUID
	Status     Status
	DoneBy     uuid.UUID
}

type TermAnnotationPart struct {
	AnnotId         uuid.UUID
	PartNo          int
	OrigSentNo      int
	TransSentNo     int
	BegOrigTokenNo  int
	EndOrigTokenNo  int
	BegTransTokenNo int
	EndTransTokenNo int
}

func NewAnnotationTask(orig_doc_id, trans_doc_id uuid.UUID, desc string, created_by uuid.UUID) *AnnotationTask {
	at := &AnnotationTask{
		Id:             uuid.New(),
		OrigDocId:      orig_doc_id,
		TransDocId:     trans_doc_id,
		Description:    desc,
		CreatedBy:      created_by,
		CreationDate:   time.Now(),
		LastUpdateDate: time.Now(),
	}
	return at
}

func NewStructAnnotation(task_id, orig_doc_id uuid.UUID, beg_sent_no, end_sent_no int, status Status, done_by uuid.UUID) *StructAnnotation {
	sa := &StructAnnotation{
		Id:        uuid.New(),
		TaskId:    task_id,
		OrigDocId: orig_doc_id,
		BegSentNo: beg_sent_no,
		EndSentNo: end_sent_no,
		Status:    status,
		DoneBy:    done_by,
	}
	return sa
}

func NewTermAnnotation(task_id, orig_doc_id, trans_doc_id uuid.UUID, status Status, done_by uuid.UUID) *TermAnnotation {
	ta := &TermAnnotation{
		Id:         uuid.New(),
		TaskId:     task_id,
		OrigDocId:  orig_doc_id,
		TransDocId: trans_doc_id,
		Status:     status,
		DoneBy:     done_by,
	}
	return ta
}

func NewTermAnnotationPart(annot_id uuid.UUID, part_no, orig_sent_no, trans_sent_no, beg_orig_token_no, end_orig_token_no, beg_trans_token_no, end_trans_token_no int) *TermAnnotationPart {
	chicken := &TermAnnotationPart{
		AnnotId:         annot_id,
		PartNo:          part_no,
		OrigSentNo:      orig_sent_no,
		TransSentNo:     trans_sent_no,
		BegOrigTokenNo:  beg_orig_token_no,
		EndOrigTokenNo:  end_orig_token_no,
		BegTransTokenNo: beg_trans_token_no,
		EndTransTokenNo: end_trans_token_no,
	}
	return chicken
}
