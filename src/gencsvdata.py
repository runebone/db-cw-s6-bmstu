from faker import Faker
import uuid
import csv

fake = Faker()

def create_csv_title(filename, data):
    with open(filename, 'w', newline='', encoding='utf-8') as file:
        writer = csv.writer(file)
        writer.writerow(data)

def write_to_csv(filename, data):
    with open(filename, 'a', newline='', encoding='utf-8') as file:
        writer = csv.writer(file)
        writer.writerow(data)

user_data_ids = []
USER_DATA_CSV = "csv/user_data.csv"
create_csv_title(USER_DATA_CSV, [
    "id",
    "username",
    "email",
    "pwd_hash",
    "trust_level",
    "creation_date",
    "last_update_date",
    ])
def gen_user_data(csv=USER_DATA_CSV):
    id = str(uuid.uuid4())
    user_data_ids.append(id)
    username = fake.user_name() \
    + str(fake.random.randint(1, 9999))
    email = fake.email()
    pwd_hash = fake.sha256()
    trust_level = 0
    creation_date = fake.date_time()
    last_update_date = creation_date
    write_to_csv(csv, [
        id,
        username,
        email,
        pwd_hash,
        trust_level,
        creation_date,
        last_update_date
        ])

document_ids = []
DOCUMENT_CSV = "csv/document.csv"
create_csv_title(DOCUMENT_CSV, [
    "id",
    "url",
    "title",
    "lang",
    "orig_doc_id",
    "uploaded_by",
    "upload_date",
    ])
def gen_document(csv=DOCUMENT_CSV):
    id = str(uuid.uuid4())
    document_ids.append(id)
    url = fake.url()
    title = fake.text(max_nb_chars=256)
    lang = fake.random_element(['ru', 'en', 'de', 'zh', 'ar'])
    orig_doc_id = fake.random_element(document_ids)
    uploaded_by = fake.random_element(user_data_ids)
    upload_date = fake.date_time()
    write_to_csv(csv, [
        id,
        url,
        title,
        lang,
        orig_doc_id,
        uploaded_by,
        upload_date
        ])

author_ids = []
AUTHOR_CSV = "csv/author.csv"
create_csv_title(AUTHOR_CSV, [
    "id",
    "first_name",
    "last_name",
    "mid_name",
    "email",
    "creation_date",
    ])
def gen_author(csv=AUTHOR_CSV):
    id = str(uuid.uuid4())
    author_ids.append(id)
    name = fake.name().split()
    first_name = name[0]
    mid_name = name[1 % len(name)]
    last_name = name[2 % len(name)]
    email = fake.email()
    creation_date = fake.date_time()
    write_to_csv(csv, [
        id,
        first_name,
        mid_name,
        last_name,
        email,
        creation_date
        ])

DOCUMENT_AUTHOR_CSV = "csv/document_author.csv"
create_csv_title(DOCUMENT_AUTHOR_CSV, [
    "doc_id",
    "author_id",
    ])
def gen_document_author(csv=DOCUMENT_AUTHOR_CSV):
    doc_id = fake.random_element(document_ids)
    author_id = fake.random_element(author_ids)
    write_to_csv(csv, [
        doc_id,
        author_id
        ])

annotation_task_ids = []
ANNOTATION_TASK_CSV = "csv/annotation_task.csv"
create_csv_title(ANNOTATION_TASK_CSV, [
    "id",
    "orig_doc_id",
    "trans_doc_id",
    "description",
    "created_by",
    "creation_date",
    "last_update_date",
    ])
def gen_annotation_task(csv=ANNOTATION_TASK_CSV):
    id = str(uuid.uuid4())
    annotation_task_ids.append(id)
    orig_doc_id = fake.random_element(document_ids)
    trans_doc_id = fake.random_element(document_ids)
    description = fake.text(max_nb_chars=256)
    created_by = fake.random_element(user_data_ids)
    creation_date = fake.date_time()
    last_update_date = fake.date_time()
    write_to_csv(csv, [
        id,
        orig_doc_id,
        trans_doc_id,
        description,
        created_by,
        creation_date,
        last_update_date
        ])

STRUCT_ANNOTATION_CSV = "csv/struct_annotation.csv"
create_csv_title(STRUCT_ANNOTATION_CSV, [
    "id",
    "task_id",
    "orig_doc_id",
    "beg_sent_no",
    "end_sent_no",
    "status",
    "done_by",
    ])
def gen_struct_annotation(csv=STRUCT_ANNOTATION_CSV):
    id = str(uuid.uuid4())
    task_id = fake.random_element(annotation_task_ids)
    orig_doc_id = fake.random_element(document_ids)
    beg_sent_no = str(fake.random.randint(1, 9999))
    end_sent_no = str(int(beg_sent_no) + fake.random.randint(1, 9999))
    status = fake.random_element(["approved", "rejected", ""])
    done_by = fake.random_element(user_data_ids)
    write_to_csv(csv, [
        id,
        task_id,
        orig_doc_id,
        beg_sent_no,
        end_sent_no,
        status,
        done_by
        ])

term_annotation_ids = []
TERM_ANNOTATION_CSV = "csv/term_annotation.csv"
create_csv_title(TERM_ANNOTATION_CSV, [
    "id",
    "task_id",
    "orig_doc_id",
    "trans_doc_id",
    "status",
    "done_by",
    ])
def gen_term_annotation(csv=TERM_ANNOTATION_CSV):
    id = str(uuid.uuid4())
    term_annotation_ids.append(id)
    task_id = fake.random_element(annotation_task_ids)
    orig_doc_id = fake.random_element(document_ids)
    trans_doc_id = fake.random_element(document_ids)
    status = fake.random_element(["approved", "rejected", ""])
    done_by = fake.random_element(user_data_ids)
    write_to_csv(csv, [
        id,
        task_id,
        orig_doc_id,
        trans_doc_id,
        status,
        done_by
        ])

TERM_ANNOTATION_PART_CSV = "csv/term_annotation_part.csv"
create_csv_title(TERM_ANNOTATION_PART_CSV, [
    "annot_id",
    "part_no",
    "orig_sent_no",
    "trans_sent_no",
    "beg_orig_token_no",
    "end_orig_token_no",
    "beg_trans_token_no",
    "end_trans_token_no",
    ])
def gen_term_annotation_part(csv=TERM_ANNOTATION_PART_CSV):
    annot_id = fake.random_element(term_annotation_ids)
    part_no = str(fake.random.randint(1, 9999))
    orig_sent_no = str(fake.random.randint(1, 9999))
    trans_sent_no = str(fake.random.randint(1, 9999))
    beg_orig_token_no = str(fake.random.randint(1, 9999))
    end_orig_token_no = str(fake.random.randint(1, 9999))
    beg_trans_token_no = str(fake.random.randint(1, 9999))
    end_trans_token_no = str(fake.random.randint(1, 9999))
    write_to_csv(csv, [
        annot_id,
        part_no,
        orig_sent_no,
        trans_sent_no,
        beg_orig_token_no,
        end_orig_token_no,
        beg_trans_token_no,
        end_trans_token_no
        ])

SENTENCE_CSV = "csv/sentence.csv"
create_csv_title(SENTENCE_CSV, [
    "doc_id",
    "sent_no",
    "content",
    ])
def gen_sentence(csv=SENTENCE_CSV):
    doc_id = fake.random_element(document_ids)
    sent_no = str(fake.random.randint(1, 9999))
    content = fake.text(max_nb_chars=256)
    write_to_csv(csv, [
        doc_id,
        sent_no,
        content
        ])

TOKEN_CSV = "csv/token.csv"
create_csv_title(TOKEN_CSV, [
    "doc_id",
    "sent_no",
    "token_no",
    "content",
    ])
def gen_token(csv=TOKEN_CSV):
    doc_id = fake.random_element(document_ids)
    sent_no = str(fake.random.randint(1, 9999))
    token_no = str(fake.random.randint(1, 9999))
    content = fake.text(max_nb_chars=256)
    write_to_csv(csv, [
        doc_id,
        sent_no,
        token_no,
        content
        ])

if __name__ == "__main__":
    for _ in range(100):
        gen_user_data()
        gen_document()
        gen_author()
    for _ in range(100):
        gen_document_author()
        gen_annotation_task()
    for _ in range(100):
        gen_struct_annotation()
        gen_term_annotation()
    for _ in range(100):
        gen_term_annotation_part()
        gen_sentence()
        gen_token()
