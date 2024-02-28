CREATE TABLE department(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(32) NOT NULL,
    name_i18n varchar(64) NULL,
    contact_number varchar(32) NULL,
    contact_address_id int8,
    parent_id int8 NULL,
    FOREIGN KEY (parent_id) REFERENCES department(id),
    FOREIGN KEY (contact_address_id) REFERENCES contact_address(id),
);

-- 为表添加注释
COMMENT ON TABLE department IS '部门表';

