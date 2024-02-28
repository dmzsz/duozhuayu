-- tag表
CREATE TABLE tag(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(32) NOT NULL,
    name_i18n varchar(64),
);

-- 为表添加注释
COMMENT ON TABLE tag IS 'tag对应表';

