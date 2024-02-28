-- 仓库
CREATE TABLE warehouse(
    id int8 PRIMARY KEY,
    name_text varchar(200) NOT NULL,
    contact_address_id in8,
    contact_number varchar(32) NULL,
    capacity integer NULL,
    manager_id int8 NULL,
    status store_opening_status_enum DEFAULT 'open',
    type warehouse_type_enum NULL,
    FOREIGN KEY (contact_address_id) REFERENCES contact_address(id),
    FOREIGN KEY (manager_id) REFERENCES "user"(id),
);

-- 为表添加注释
COMMENT ON TABLE warehouse IS '仓库表';

COMMENT ON COLUMN warehouse.manager_id IS '管理员';

COMMENT ON COLUMN warehouse.capacity IS '容量';

COMMENT ON COLUMN warehouse.type IS '仓库类型';

