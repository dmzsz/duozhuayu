-- 地区表
CREATE TABLE region(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    -- 洲
    continent_text varchar(50),
    continent_i18n varchar(50),
    -- 国家
    country_text varchar(50),
    country_iso_code varchar(50),
    country_i18n varchar(50),
    -- 国家代码
    country_code varchar(12),
    -- 省
    province varchar(50),
    -- 市
    city varchar(50),
    -- 区
    district varchar(50),
    -- 街道
    street varchar(50),
    -- 座机区号
    area_code varchar(4),
    -- 邮编
    postal_code varchar(10),
    parent_region_id int8 NULL,
    FOREIGN KEY (parent_region_id) REFERENCES region(id),
);

-- 为表添加注释
COMMENT ON TABLE region IS '地区信息表';

COMMENT ON COLUMN region.area_code IS '座机区号';

COMMENT ON COLUMN region.postal_code IS '邮编';

COMMENT ON COLUMN region.continent_text IS '洲 亚洲 非洲 欧洲';

COMMENT ON COLUMN region.country_text IS '国家';

COMMENT ON COLUMN region.country_iso_code IS '国家缩写 中国 CHN';

COMMENT ON COLUMN region.country_code IS '全球国际电话服务的区号列表 手机号前面的数字 中国是 +86';

COMMENT ON COLUMN region.province IS '省';

COMMENT ON COLUMN region.city IS '市';

COMMENT ON COLUMN region.district IS '区';

COMMENT ON COLUMN region.street IS '街道';

