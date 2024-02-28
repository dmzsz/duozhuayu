CREATE TABLE organisation(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(32) NOT NULL,
    name_i18n varchar(64) NULL,
    contact_number varchar(32) NULL,
    description varchar(1000),
    image_id int8 NULL,
    parent_id int8 NULL,
    contact_address_id int8,
    FOREIGN KEY (parent_id) REFERENCES organisation(id),
    FOREIGN KEY (contact_address_id) REFERENCES contact_address(id)
);

-- 为表添加注释
COMMENT ON TABLE organisation IS '组织表';

