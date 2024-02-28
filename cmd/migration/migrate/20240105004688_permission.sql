CREATE TABLE permission(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name varchar(32) NOT NULL,
    code varchar(32) NULL,
    menu_path varchar(32)[] NULL,
    source_operating_range operating_range_enum NULL,
    source varchar(32)[] NULL,
    source_field_operating_range operating_range_enum NULL,
    source_field varchar(32)[] NULL,
    api_name varchar(32)[] NULL,
    allow_action permission_action_enum NULL,
    parent_id int8 NULL,
    FOREIGN KEY (parentId) REFERENCES permission(id)
);

-- 为表添加注释
COMMENT ON TABLE permission IS '权限表';

COMMENT ON COLUMN permission.menu_path IS '可以访问的前端菜单路径';

COMMENT ON COLUMN permission.source_operating_range IS '排除数据库范围。include 包含 excluding 排除';

COMMENT ON COLUMN permission.source IS '可以操作的数据库表';

COMMENT ON COLUMN permission.source_operating_range IS '字段选择范围。include 包含 excluding 排除';

COMMENT ON COLUMN permission.source_field IS '可以操作的数据库表字段';

COMMENT ON COLUMN permission.api_name IS '接口调用名称， 功能名称';

COMMENT ON COLUMN permission.allow_action IS '允许的操作 可读 可写 所有(删除权限包含读写权限，也就是所有权限)';

