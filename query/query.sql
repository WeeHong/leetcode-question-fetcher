SET timezone TO 'Asia/Singapore';

DROP TABLE IF EXISTS records;
DROP TABLE IF EXISTS questions_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS questions;

CREATE TABLE IF NOT EXISTS records (
    id SERIAL UNIQUE NOT NULL PRIMARY KEY,
    total INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tags (
	id VARCHAR(255) UNIQUE NOT NULL PRIMARY KEY,
	name VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS questions (
    id INTEGER UNIQUE NOT NULL PRIMARY KEY,
	title VARCHAR(255),
	slug VARCHAR(255),
    difficulty VARCHAR(20),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS questions_tags (
    id SERIAL UNIQUE NOT NULL PRIMARY KEY,
    question_id INTEGER,
    tag_id VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_question
        FOREIGN KEY(question_id)
            REFERENCES questions(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_tag
        FOREIGN KEY(tag_id)
            REFERENCES tags(id)
            ON DELETE CASCADE
);

INSERT INTO records(total) VALUES (0);

SELECT QuestionsTags.question_id, Questions.title, Questions.slug, Questions.difficulty, string_agg(tag.name, ', ')
FROM QuestionsTags
LEFT JOIN question ON Questions.id = QuestionsTags.question_id
LEFT JOIN tag ON tag.id = QuestionsTags.tag_id
GROUP BY QuestionsTags.question_id, Questions.title, Questions.slug, Questions.difficulty;
