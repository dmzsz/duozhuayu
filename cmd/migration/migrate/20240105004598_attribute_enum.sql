CREATE TABLE attribute_enum(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    attr_id int8 NOT NULL, -- 用于前端url显示或者前端标识的key
    code attr_type_enum NOT NULL, -- spu_spec json的values。比如 red white
    name_text varchar(32) NOT NULL, -- 用于页面显示的中文值
    name_i18n varchar(64) DEFAULT NULL,
    sort real DEFAULT 0,
    FOREIGN KEY (attr_id) REFERENCES attribute(id)
);

-- 为表添加注释
COMMENT ON TABLE attribute_enum IS '商品属性名表';

-- 为列添加注释
COMMENT ON COLUMN attribute_enum.attr_id IS '用于前端url显示或者前端标识的key';

COMMENT ON COLUMN attribute_enum.enum_code IS 'json的key';

COMMENT ON COLUMN attribute_enum.enum_code IS 'spu_spec json的values。比如 red white';

COMMENT ON COLUMN attribute_enum.sort IS '排序';

COMMENT ON COLUMN attribute_enum.name_text IS 'spu_spec json的values。用于页面显示的中文值,比如 红色 白色 基础班';

