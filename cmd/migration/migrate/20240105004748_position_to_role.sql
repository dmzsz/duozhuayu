CREATE TABLE position_to_role(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name varchar(32) NOT NULL,
    description varchar(128),
    position_id int8 NOT NULL,
    role_id int8 NOT NULL,
    parent_id int8 NULL,
    FOREIGN KEY (position_id) REFERENCES position(id),
    FOREIGN KEY (role_id) REFERENCES "role"(id),
    FOREIGN KEY (parent_id) REFERENCES position_to_role(id),
    CONSTRAINT unique_position_role UNIQUE (position_id, role_id)
);

-- 为表添加注释
COMMENT ON TABLE position_to_role IS '职位角色表';

