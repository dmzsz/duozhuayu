CREATE TABLE user_to_role(
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name varchar(32) NULL,
    description varchar(128) NULL,
    user_id int8 NOT NULL,
    role_id int8 NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    FOREIGN KEY (role_id) REFERENCES "role"(id),
    CONSTRAINT unique_user_role UNIQUE (user_id, role_id)
);

-- 为表添加注释
COMMENT ON TABLE user_to_role IS '用户职位表';

