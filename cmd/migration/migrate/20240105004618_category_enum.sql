CREATE TABLE category_enum(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(32) NOT NULL,
    name_i18n varchar(64) DEFAULT NULL,
    description varchar(128) NULL,
);

-- 为表添加注释
COMMENT ON TABLE category_enum IS 'categrory enum表 减少重复categore name';