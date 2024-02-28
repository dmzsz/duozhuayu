CREATE TABLE "role"(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name varchar(32) NOT NULL,
    description varchar(128) NULL,
    parent_id int8 NULL,
    position_id int8 NULL,
    FOREIGN KEY (parent_id) REFERENCES "role"(id),
    FOREIGN KEY (position_id) REFERENCES position(id)
);

-- 为表添加注释
COMMENT ON TABLE "role" IS '角色表';

