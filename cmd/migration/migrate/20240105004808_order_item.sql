-- 订单项目表
CREATE TABLE order_item(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    price decimal(10, 2) NOT NULL,
    count integer NOT NULL DEFAULT 1,
    is_locked bool DEFAULT FALSE,
    locked_number integer DEFAULT 0,
    order_id int8 NULL,
    sku_id int8 NULL,
    bundle_sku_id int8 NULL,
    image_id int8 NULL,
    FOREIGN KEY (order_id) REFERENCES "order"(id),
    FOREIGN KEY (sku_id) REFERENCES sku(id),
    FOREIGN KEY (bundle_sku_id) REFERENCES bundle_sku(id),
    FOREIGN KEY (image_id) REFERENCES image(id),
);

-- 为表添加注释
COMMENT ON TABLE order_item IS '订单项目表';

COMMENT ON COLUMN order_item.order_id IS '订单id';

COMMENT ON COLUMN order_item.image_id IS 'spu 页面快照图id';

COMMENT ON COLUMN order_item.is_locked IS '锁定sku';

COMMENT ON COLUMN order_item.locked_number IS '锁定的数量';

