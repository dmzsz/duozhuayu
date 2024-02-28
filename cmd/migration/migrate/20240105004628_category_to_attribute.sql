-- 针对某个category包含的attribute 的记录表
CREATE TABLE category_to_attribute(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    category_id int8 NOT NULL,
    attribute_id int8 NOT NULL,
    is_required bool NOT NULL DEFAULT FALSE,
    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (attribute_id) REFERENCES attribute(id),
    CONSTRAINT unique_category_attribute UNIQUE (category_id, attribute_id)
);

-- 为表添加注释
COMMENT ON TABLE category_to_attribute IS '针对某个category包含的attribute 的记录表';