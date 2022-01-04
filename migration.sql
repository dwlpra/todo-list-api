CREATE TABLE activities (
id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
email VARCHAR(255) NOT NULL,
title VARCHAR(255) NOT NULL,
created_at TIMESTAMP,
updated_at TIMESTAMP,
deleted_at TIMESTAMP
);

INSERT INTO activities (email, title, created_at,updated_at) VALUES ('ad0286a7-bec4-405c-96e2-cd472c18a1e7@skyshi.com', 'coba 4', '2021-11-30 05:29:24','2021-11-30 05:29:24','2021-11-30 05:29:24');