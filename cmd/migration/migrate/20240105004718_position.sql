CREATE TABLE position (
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(32) NOT NULL,
    name_i18n varchar(64) NULL,
    parent_id int8 NULL,
    FOREIGN KEY (parent_id) REFERENCES position(id),
);

-- 为表添加注释
COMMENT ON TABLE "position" IS '职位表';

