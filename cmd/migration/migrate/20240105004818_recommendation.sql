-- 推荐表
CREATE TABLE recommendation(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    user_id int8,
    spu_id int8,
    rating int8,
    review text,
    tag varchar(64),
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    FOREIGN KEY (spu_id) REFERENCES spu(id)
);

