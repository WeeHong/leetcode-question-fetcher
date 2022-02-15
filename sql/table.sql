SET timezone TO 'Asia/Singapore';

DROP TABLE IF EXISTS Records;
DROP TABLE IF EXISTS QuestionsTags;
DROP TABLE IF EXISTS Tags;
DROP TABLE IF EXISTS Questions;

CREATE TABLE IF NOT EXISTS Records (
    id SERIAL UNIQUE NOT NULL PRIMARY KEY,
    total INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Tags (
	id VARCHAR(255) UNIQUE NOT NULL PRIMARY KEY,
	name VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Questions (
    id INTEGER UNIQUE NOT NULL PRIMARY KEY,
	title VARCHAR(255),
	slug VARCHAR(255),
    difficulty VARCHAR(20),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS QuestionsTags (
    id SERIAL UNIQUE NOT NULL PRIMARY KEY,
    question_id INTEGER,
    tag_id VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_question
        FOREIGN KEY(question_id)
            REFERENCES Questions(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_tag
        FOREIGN KEY(tag_id)
            REFERENCES tags(id)
            ON DELETE CASCADE
);

INSERT INTO Records(total) VALUES (0);