package models

import (
	"time"

	"github.com/google/uuid"
)

type AnnotationTask struct {
	ID             uuid.UUID
	OrigDocID      uuid.UUID
	TransDocID     uuid.UUID
	Description    string
	CreatedBy      uuid.UUID
	CreationDate   time.Time
	LastUpdateDate time.Time
}

type StructAnnotation struct {
	ID        uuid.UUID
	TaskID    uuid.UUID
	OrigDocID uuid.UUID
	BegSentNo int
	EndSentNo int
	Status    Status
	DoneBy    uuid.UUID
}

type Status string

type TermAnnotation struct {
	ID         uuid.UUID
	TaskID     uuid.UUID
	OrigDocID  uuid.UUID
	TransDocID uuid.UUID
	Status     Status
	DoneBy     uuid.UUID
}

type TermAnnotationPart struct {
	AnnotID         uuid.UUID
	PartNo          int
	OrigSentNo      int
	TransSentNo     int
	BegOrigTokenNo  int
	EndOrigTokenNo  int
	BegTransTokenNo int
	EndTransTokenNo int
}

func NewAnnotationTask(origDocID, transDocID uuid.UUID, desc string, createdBy uuid.UUID) *AnnotationTask {
	at := &AnnotationTask{
		ID:             uuid.New(),
		OrigDocID:      origDocID,
		TransDocID:     transDocID,
		Description:    desc,
		CreatedBy:      createdBy,
		CreationDate:   time.Now(),
		LastUpdateDate: time.Now(),
	}
	return at
}

func NewStructAnnotation(taskID, origDocID uuid.UUID, begSentNo, endSentNo int, status Status, doneBy uuid.UUID) *StructAnnotation {
	sa := &StructAnnotation{
		ID:        uuid.New(),
		TaskID:    taskID,
		OrigDocID: origDocID,
		BegSentNo: begSentNo,
		EndSentNo: endSentNo,
		Status:    status,
		DoneBy:    doneBy,
	}
	return sa
}

func NewTermAnnotation(taskID, origDocID, transDocID uuid.UUID, status Status, doneBy uuid.UUID) *TermAnnotation {
	ta := &TermAnnotation{
		ID:         uuid.New(),
		TaskID:     taskID,
		OrigDocID:  origDocID,
		TransDocID: transDocID,
		Status:     status,
		DoneBy:     doneBy,
	}
	return ta
}

func NewTermAnnotationPart(annotID uuid.UUID, partNo, origSentNo, transSentNo, begOrigTokenNo, endOrigTokenNo, begTransTokenNo, endTransTokenNo int) *TermAnnotationPart {
	chicken := &TermAnnotationPart{
		AnnotID:         annotID,
		PartNo:          partNo,
		OrigSentNo:      origSentNo,
		TransSentNo:     transSentNo,
		BegOrigTokenNo:  begOrigTokenNo,
		EndOrigTokenNo:  endOrigTokenNo,
		BegTransTokenNo: begTransTokenNo,
		EndTransTokenNo: endTransTokenNo,
	}
	return chicken
}
