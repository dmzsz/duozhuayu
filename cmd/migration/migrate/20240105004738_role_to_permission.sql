CREATE TABLE role_to_permission(
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    role_id int8 NOT NULL,
    permission_id int8 NOT NULL,
    FOREIGN KEY (role_id) REFERENCES "role"(id),
    FOREIGN KEY (permission_id) REFERENCES permission(id),
    CONSTRAINT unique_role_permission UNIQUE (role_id, permission_id)
);

-- 为表添加注释
COMMENT ON TABLE role_to_permission IS '角色权限表';

