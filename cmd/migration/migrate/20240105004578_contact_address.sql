CREATE TABLE contact_address(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    user_id int8 DEFAULT NULL,
    region_id int8 DEFAULT NULL,
    address_type varchar(32) DEFAULT NULL,
    address varchar(256) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    FOREIGN KEY (region_id) REFERENCES region(id)
);

-- 为表添加注释
COMMENT ON TABLE contact_address IS '联系地址表';

COMMENT ON COLUMN contact_address.address_type IS '使用范围 微信联系方式  收发货联系方式 公司联系方式 部门联系方式';