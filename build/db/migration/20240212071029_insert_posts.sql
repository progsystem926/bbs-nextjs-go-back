-- +goose Up
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    text VARCHAR(255) DEFAULT NULL,
    user_id INT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX post_id on posts (id);

INSERT INTO posts (text, user_id, created_at)
VALUES
    ('post1,post1,post1,post1,post1,post1,post1,post1,post1,post1,post1,', 1, DATE_SUB(NOW(), INTERVAL FLOOR(RAND()*365) DAY)),
    ('post2,post2,post2,post2,post2,post2,post2,post2,post2,post2,post2,', 2, DATE_SUB(NOW(), INTERVAL FLOOR(RAND()*365) DAY)),
    ('post3,post3,post3,post3,post3,post3,post3,post3,post3,post3,post3,', 3, DATE_SUB(NOW(), INTERVAL FLOOR(RAND()*365) DAY));

-- +goose Down
DROP INDEX post_id on posts;
DROP TABLE posts;
