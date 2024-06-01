ALTER TABLE user_data
    ALTER COLUMN id SET NOT NULL,
    ADD CONSTRAINT pk_user_data_id
    PRIMARY KEY (id),
    ALTER COLUMN username SET NOT NULL,
    ADD CONSTRAINT unique_user_data_username
    UNIQUE (username),
    ALTER COLUMN email SET NOT NULL,
    ALTER COLUMN pwd_hash SET NOT NULL,
    ALTER COLUMN trust_level SET NOT NULL,
    ALTER COLUMN trust_level SET DEFAULT 0,
    ALTER COLUMN creation_date SET NOT NULL,
    ALTER COLUMN last_update_date SET DEFAULT NULL;

ALTER TABLE document
    ALTER COLUMN id SET NOT NULL,
    ADD CONSTRAINT pk_document_id
    PRIMARY KEY (id),
    ALTER COLUMN url SET NOT NULL,
    ALTER COLUMN title SET NOT NULL,
    ALTER COLUMN lang SET NOT NULL,
    ADD CONSTRAINT fk_document_orig_doc_id
    FOREIGN KEY (orig_doc_id)
    REFERENCES document(id),
    ALTER COLUMN uploaded_by SET NOT NULL,
    ADD CONSTRAINT fk_document_uploaded_by
    FOREIGN KEY (uploaded_by)
    REFERENCES user_data(id),
    ALTER COLUMN upload_date SET NOT NULL;

ALTER TABLE author
    ALTER COLUMN id SET NOT NULL,
    ADD CONSTRAINT pk_author_id
    PRIMARY KEY (id),
    ALTER COLUMN creation_date SET NOT NULL;

ALTER TABLE document_author
    ALTER COLUMN doc_id SET NOT NULL,
    ADD CONSTRAINT fk_document_author_doc_id
    FOREIGN KEY (doc_id)
    REFERENCES document(id)
    ON DELETE CASCADE,
    ALTER COLUMN author_id SET NOT NULL,
    ADD CONSTRAINT fk_document_author_author_id
    FOREIGN KEY (author_id)
    REFERENCES author(id);

ALTER TABLE annotation_task
    ALTER COLUMN id SET NOT NULL,
    ADD CONSTRAINT pk_annotation_task_id
    PRIMARY KEY (id),
    ALTER COLUMN orig_doc_id SET NOT NULL,
    ADD CONSTRAINT fk_annotation_task_orig_doc_id
    FOREIGN KEY (orig_doc_id)
    REFERENCES document(id),
    ADD CONSTRAINT fk_annotation_task_trans_doc_id
    FOREIGN KEY (trans_doc_id)
    REFERENCES document(id),
    ALTER COLUMN description SET NOT NULL,
    ADD CONSTRAINT fk_annotation_task_created_by
    FOREIGN KEY (created_by)
    REFERENCES user_data(id),
    ALTER COLUMN creation_date SET NOT NULL,
    ALTER COLUMN last_update_date SET NOT NULL;

ALTER TABLE struct_annotation
    ALTER COLUMN id SET NOT NULL,
    ADD CONSTRAINT pk_struct_annotation_id
    PRIMARY KEY (id),
    ALTER COLUMN task_id SET NOT NULL,
    ADD CONSTRAINT fk_struct_annotation_task_id
    FOREIGN KEY (task_id)
    REFERENCES annotation_task(id),
    ALTER COLUMN orig_doc_id SET NOT NULL,
    ADD CONSTRAINT fk_struct_annotation_orig_doc_id
    FOREIGN KEY (orig_doc_id)
    REFERENCES document(id),
    ALTER COLUMN beg_sent_no SET NOT NULL,
    ALTER COLUMN end_sent_no SET NOT NULL,
    ALTER COLUMN done_by SET NOT NULL,
    ADD CONSTRAINT fk_struct_annotation_done_by
    FOREIGN KEY (done_by)
    REFERENCES user_data(id);

ALTER TABLE term_annotation
    ALTER COLUMN id SET NOT NULL,
    ADD CONSTRAINT pk_term_annotation_id
    PRIMARY KEY (id),
    ALTER COLUMN task_id SET NOT NULL,
    ADD CONSTRAINT fk_term_annotation_task_id
    FOREIGN KEY (task_id)
    REFERENCES annotation_task(id),
    ALTER COLUMN orig_doc_id SET NOT NULL,
    ADD CONSTRAINT fk_term_annotation_orig_doc_id
    FOREIGN KEY (orig_doc_id)
    REFERENCES document(id),
    ADD CONSTRAINT fk_term_annotation_trans_doc_id
    FOREIGN KEY (trans_doc_id)
    REFERENCES document(id),
    ALTER COLUMN done_by SET NOT NULL,
    ADD CONSTRAINT fk_term_annotation_done_by
    FOREIGN KEY (done_by)
    REFERENCES user_data(id);

ALTER TABLE term_annotation_part
    ADD CONSTRAINT pk_term_annotation_part_annot_id_part_no
    PRIMARY KEY (annot_id, part_no),
    ADD CONSTRAINT fk_term_annotation_part_annot_id
    FOREIGN KEY (annot_id)
    REFERENCES term_annotation(id)
    ON DELETE CASCADE,
    ALTER COLUMN orig_sent_no SET NOT NULL,
    ALTER COLUMN trans_sent_no SET NOT NULL,
    ALTER COLUMN beg_orig_token_no SET NOT NULL,
    ALTER COLUMN end_orig_token_no SET NOT NULL,
    ALTER COLUMN beg_trans_token_no SET NOT NULL,
    ALTER COLUMN end_trans_token_no SET NOT NULL;

ALTER TABLE sentence
    ADD CONSTRAINT pk_sentence_doc_id_sent_no
    PRIMARY KEY (doc_id, sent_no),
    ADD CONSTRAINT fk_sentence_doc_id
    FOREIGN KEY (doc_id)
    REFERENCES document(id),
    ALTER COLUMN content SET NOT NULL;

ALTER TABLE token
    ADD CONSTRAINT pk_token_doc_id_sent_no_token_no
    PRIMARY KEY (doc_id, sent_no, token_no),
    ADD CONSTRAINT fk_token_doc_id
    FOREIGN KEY (doc_id)
    REFERENCES document(id),
    ALTER COLUMN content SET NOT NULL;
