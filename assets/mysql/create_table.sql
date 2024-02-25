CREATE TABLE users
(
    id           INT                 NOT NULL AUTO_INCREMENT,
    username     VARCHAR(255) UNIQUE NOT NULL,
    real_name    VARCHAR(255)        NOT NULL,
    email        VARCHAR(255)        NOT NULL,
    phone_number VARCHAR(255)        NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
