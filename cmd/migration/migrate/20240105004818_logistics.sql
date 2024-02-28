-- 物流信息
CREATE TABLE logistics(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    tracking_no varchar(32),
    logistics_company varchar(32),
    delivered_at timestamp NULL,
    receipted_at timestamp NULL,
    node_info json NULL,
    sku_ids int8[] NULL,
    order_id int8[] NULL,
    warehouse_id int8,
    parent_id int8 NULL,
    FOREIGN KEY (order_id) REFERENCES "order"(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id),
    FOREIGN KEY (parent_id) REFERENCES logistics(id)
);

-- 为表添加注释
COMMENT ON TABLE logistics IS '物流信息表';

COMMENT ON COLUMN logistics.tracking_no IS '快递单号';

COMMENT ON COLUMN logistics.logistics_company IS '货运公司';

COMMENT ON COLUMN logistics.delivered_at IS '发货时间';

COMMENT ON COLUMN logistics.receipted_at IS '签收时间';

COMMENT ON COLUMN logistics.node_info IS '货运节点信息，物流公司那里获取到的';

COMMENT ON COLUMN logistics.sku_ids IS '包裹包含的sku商品id';

COMMENT ON COLUMN logistics.order_id IS '订单id';

COMMENT ON COLUMN logistics.parent_id IS '拆包父级id';

