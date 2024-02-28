CREATE TABLE user_to_position(
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    user_id int8 NOT NULL,
    position_id int8 NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    FOREIGN KEY (position_id) REFERENCES position(id),
    CONSTRAINT unique_user_position UNIQUE (user_id, position_id)
);

-- 为表添加注释
COMMENT ON TABLE user_to_position IS '用户职位表';

