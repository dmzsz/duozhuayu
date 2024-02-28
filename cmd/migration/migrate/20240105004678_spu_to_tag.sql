-- spu tag对应表
CREATE TABLE spu_to_tag(
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    spu_id int8,
    tag_id varchar,
    description varchar(128),
    theme_color varchar(32) NULL,
    FOREIGN KEY (spu_id) REFERENCES spu(id),
    FOREIGN KEY (tag_id) REFERENCES tag(id),
    CONSTRAINT unique_spu_tag UNIQUE (spu_id, tag_id),
);

-- 为表添加注释
COMMENT ON TABLE spu_to_tag IS 'spu tag对应表';