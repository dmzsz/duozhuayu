-- 如果是entity_type是 spu 需要取goods_id 和 spu表join。
CREATE TABLE bundle_sku(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(32) NULL,
    name_i18n varchar(64) NULL,
    spu_id int8 NOT NULL,
    sku_id int8 NOT NULL,
    quantity real DEFAULT NULL,
    price DECIMAL(10, 2) DEFAULT NULL,
    sort real DEFAULT 0,
    shelf_status sale_status_enum NULL,
    parent_id int8 NULL,
    FOREIGN KEY (spu_id) REFERENCES spu(id),
    FOREIGN KEY (sku_id) REFERENCES sku(id),
    FOREIGN KEY (parent_id) REFERENCES bundle_sku(id),
);

-- 为表添加注释
COMMENT ON TABLE bundle_sku IS '捆绑销售的商品 或者 组合不同sku销售';

COMMENT ON COLUMN bundle_sku.quantity IS '包含的数量';

COMMENT ON COLUMN bundle_sku.price IS '组合后的价格。一般组合后会比单独购买sku有一定的优惠 需要在entity_type为spu的记录上填写价格';

COMMENT ON COLUMN bundle_sku.shelf_status IS '上架状态 "not_on_sale":未上架,"on_sale" 已上架,"off_sale" 已下架';

