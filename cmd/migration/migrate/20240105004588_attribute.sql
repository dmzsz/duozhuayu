CREATE TABLE attribute(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    code varchar(32) NOT NULL,
    attr_type attr_type_enum NOT NULL DEFAULT 'product_attr',
    attr_type_other varchar(32) NULL,
    name_text varchar(32) NOT NULL,
    name_i18n varchar(64) DEFAULT NULL,
    ui_input_type ui_input_type_enum NOT NULL 'input_string',
    value_type value_type_enum DEFAULT 'string',
    is_filter bool NOT NULL DEFAULT FALSE,
    sort real DEFAULT 0,
    shelf_status sale_status_enum DEFAULT 'not_on_sale',
    parent_attribute_id int8 NULL,
    FOREIGN KEY (parent_attribute_id) REFERENCES attribute(id)
);

-- 为表添加注释
COMMENT ON TABLE attribute IS '商品属性名表';

-- 为列添加注释
COMMENT ON COLUMN attribute.code IS '属性code 用于表示json的key';

COMMENT ON COLUMN attribute.name_text IS '属性中文名称';

COMMENT ON COLUMN attribute.attr_type IS '属性应用:product_attr商品属性,sales_attr销售属性,special_attr特殊属性';

COMMENT ON COLUMN attribute.attr_type_other IS '其他自定义属性应用';

COMMENT ON COLUMN attribute.sort IS '排序';

COMMENT ON COLUMN attribute.shelf_status IS '启用状态';

COMMENT ON COLUMN attribute.is_filter IS '是否支持前台筛选';

-- materialGroups：[{ name: "棉", percentage: 50}, { name: "莫代尔", percentage: 50}]
COMMENT ON COLUMN attribute.parent_attr_id IS 'group 属性父级id';

