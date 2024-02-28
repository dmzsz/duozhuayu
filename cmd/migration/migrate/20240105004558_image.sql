-- 图片表
CREATE TABLE image(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    name varchar(64),
    alt_text varchar(64),
    imageable_id int8 NULL,
    imageable_type varchar(32),
    size image_size_enum NULL,
    imageRatio real, -- 长宽比
    sort_num real, -- 排序
    url varchar(256),
    parent_id int8 NULL,
    FOREIGN KEY (parent_id) REFERENCES image(id),
);

