DROP   TABLE IF EXISTS     user_data;
CREATE TABLE IF NOT EXISTS user_data (
    id                     UUID,
    username               VARCHAR(32),
    email                  VARCHAR(64),
    pwd_hash               VARCHAR(256),
    trust_level            FLOAT,
    creation_date          TIMESTAMP,
    last_update_date       TIMESTAMP
    );

DROP   TABLE IF EXISTS     document;
CREATE TABLE IF NOT EXISTS document (
    id                     UUID,
    url                    VARCHAR(256),
    title                  VARCHAR(256),
    lang                   VARCHAR(2),
    orig_doc_id            UUID,
    uploaded_by            UUID,
    upload_date            TIMESTAMP
    );

DROP   TABLE IF EXISTS     document_author;
CREATE TABLE IF NOT EXISTS document_author (
    doc_id                 UUID,
    author_id              UUID
    );

DROP   TABLE IF EXISTS     author;
CREATE TABLE IF NOT EXISTS author (
    id                     UUID,
    first_name             VARCHAR(64),
    last_name              VARCHAR(64),
    mid_name               VARCHAR(64),
    email                  VARCHAR(64),
    creation_date          TIMESTAMP
    );

DROP   TABLE IF EXISTS     annotation_task;
CREATE TABLE IF NOT EXISTS annotation_task (
    id                     UUID,
    orig_doc_id            UUID,
    trans_doc_id           UUID,
    description            VARCHAR(256),
    created_by             UUID,
    creation_date          TIMESTAMP,
    last_update_date       TIMESTAMP
    );

DROP   TABLE IF EXISTS     struct_annotation;
CREATE TABLE IF NOT EXISTS struct_annotation (
    id                     UUID,
    task_id                UUID,
    orig_doc_id            UUID,
    beg_sent_no            INT,
    end_sent_no            INT,
    status                 VARCHAR(64),
    done_by                UUID
    );

DROP   TABLE IF EXISTS     term_annotation;
CREATE TABLE IF NOT EXISTS term_annotation (
    id                     UUID,
    task_id                UUID,
    orig_doc_id            UUID,
    trans_doc_id           UUID,
    status                 VARCHAR(64),
    done_by                UUID
    );

DROP   TABLE IF EXISTS     term_annotation_part;
CREATE TABLE IF NOT EXISTS term_annotation_part (
    annot_id               UUID,
    part_no                INT,
    orig_sent_no           INT,
    trans_sent_no          INT,
    beg_orig_token_no      INT,
    end_orig_token_no      INT,
    beg_trans_token_no     INT,
    end_trans_token_no     INT
    );

DROP   TABLE IF EXISTS     sentence;
CREATE TABLE IF NOT EXISTS sentence (
    doc_id                 UUID,
    sent_no                INT,
    content                VARCHAR(256)
    );

DROP   TABLE IF EXISTS     token;
CREATE TABLE IF NOT EXISTS token (
    doc_id                 UUID,
    sent_no                INT,
    token_no               INT,
    content                VARCHAR(256)
    );
