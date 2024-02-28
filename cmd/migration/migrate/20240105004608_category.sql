CREATE TABLE category(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(200) NOT NULL,
    name_id int8,
    level int4 NOT NULL DEFAULT 1,
    path varchar(32)[] DEFAULT NULL,
    theme_color varchar(32) NULL,
    description varchar(128) NULL,
    parent_category_id int8, -- 父类别ID，用于归类
    FOREIGN KEY (parent_category_id) REFERENCES category(id),
    FOREIGN KEY (name_id) REFERENCES category_enum(id)
);

COMMENT ON COLUMN category.path IS '完整父级路径：父父id_父id';

