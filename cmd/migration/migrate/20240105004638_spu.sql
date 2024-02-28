CREATE TABLE spu(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    brand_id int8 DEFAULT NULL,
    author_id int8 DEFAULT NULL,
    "name" varchar(200) NOT NULL,
    sort real DEFAULT 0,
    category_id int8 DEFAULT NULL,
    shelf_status sale_status_enum DEFAULT 'not_on_sale', -- 上架状态
    is_sold_out bool NOT NULL DEFAULT FALSE,
    original_price numeric NULL, -- 原价
    new_condition_price numeric NULL, -- 自营全新价格
    status approval_status_enum FALSE, -- 审核状态
    spu_specs json DEFAULT NULL, -- 商品规格
    creator_id int8 DEFAULT NULL, -- 创建者
    FOREIGN KEY (category_id) REFERENCES category(id)
);

-- 为表添加注释
COMMENT ON TABLE spu IS 'Standard Product Unit （标准化产品单元）';

-- 为列添加注释
COMMENT ON COLUMN spu.sort IS '权重排名';

COMMENT ON COLUMN spu.spu_specs IS '商品规格 包含sku属性值列表color:[white,red]';

COMMENT ON COLUMN spu.shelf_status IS '上架状态 "not_on_sale":未上架,"on_sale" 已上架,"off_sale" 已下架';

COMMENT ON COLUMN spu.is_deleted IS '是否删除,0:未删除，1：已删除';

COMMENT ON COLUMN spu.status IS '审核状态，0：未审核，1：已审核，2：审核不通过';

COMMENT ON COLUMN spu.is_sold_out IS '是否售罄';

COMMENT ON COLUMN spu.original_price IS '原价';

COMMENT ON COLUMN spu.new_condition_price IS '自营全新价格';

-- 创建函数 用于更新 updated_at字段
-- CREATE OR REPLACE FUNCTION update_spu_updated_at()
--     RETURNS TRIGGER
--     AS $$
-- BEGIN
--     NEW.updated_at = CURRENT_TIMESTAMP;
--     RETURN NEW;
-- END;
-- $$
-- LANGUAGE plpgsql;
-- 创建表触发器，在update之前执行update_spu_updated_at函数
-- CREATE TRIGGER spu_update_at_trigger
--     BEFORE UPDATE ON spu
--     FOR EACH ROW
--     EXECUTE FUNCTION update_spu_updated_at();
