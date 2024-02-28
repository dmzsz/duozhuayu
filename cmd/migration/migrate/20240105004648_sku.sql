CREATE TABLE sku(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    name_text varchar(32) NULL,
    name_i18n varchar(64) NULL,
    spu_id int8 NULL,
    price DECIMAL(10, 2),
    stock_quantity int8, -- 库存数量
    sort real DEFAULT 0,
    sku_attrs json DEFAULT NULL,
    shelf_status sale_status_enum DEFAULT 'not_on_sale',
    FOREIGN KEY (spu_id) REFERENCES spu(id),
);

-- 为表添加注释
COMMENT ON TABLE sku IS 'Standard Product Unit（标准化产品单元）';

-- 为列添加注释
COMMENT ON COLUMN sku.sku_attrs IS 'sku属性（商品属性）';

-- 默认是二手，还有参杂自营的情况 要排在第一个
COMMENT ON COLUMN sku.sort IS '排名';

COMMENT ON COLUMN sku.stock_quantity IS '库存';

COMMENT ON COLUMN sku.shelf_status IS '上架状态 "not_on_sale":未上架,"on_sale" 已上架,"off_sale" 已下架';

