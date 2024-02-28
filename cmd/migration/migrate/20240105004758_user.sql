CREATE TABLE "user"(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    first_name varchar(32) NULL,
    last_name varchar(32) NULL,
    username varchar(60) NOT NULL UNIQUE,
    avatar varchar(256) NULL,
    email varchar(32) NULL UNIQUE,
    gender gender_enum NULL,
    phone_number varchar(256),
    password varchar(32) NULL,
    reset_password_token varchar(32) NULL,
    reset_password_expires int4 NULL,
    is_verified bool NOT NULL DEFAULT FALSE,
    is_online bool NOT NULL DEFAULT FALSE,
    is_locked bool NOT NULL DEFAULT FALSE,
    reason varchar(256) NULL,
    is_active bool NOT NULL DEFAULT FALSE, 
    access_token varchar(32) NULL,
    refresh_token varchar(32) NULL,
    openid varchar(32) NULL,
    balance decimal(10, 2) NULL,
    unionid varchar(32) NULL,
    level varchar(32) NULL,
);

COMMENT ON COLUMN user.openid IS '微信openid';

COMMENT ON COLUMN user.unionid IS '微信unionid';

COMMENT ON COLUMN user.balance IS '余额';

COMMENT ON COLUMN user.position IS '等级';

