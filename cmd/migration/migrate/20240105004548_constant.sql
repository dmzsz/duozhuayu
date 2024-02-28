CREATE TYPE approval_status_enum AS ENUM(
    'pending', -- 未审核
    'reviewed', -- 已审核
    'rejected', -- 审核不通过
);

CREATE TYPE attr_type_enum AS ENUM(
    'product_attr', -- 商品属性
    'sales_attr', -- 销售属性
    'special_attr', -- 特殊属性
);

CREATE TYPE sale_status_enum AS ENUM(
    'not_on_sale', -- 未上架
    'on_sale', -- 已上架
    'off_sale', -- 已下架
);

CREATE TYPE gender_enum AS ENUM(
    'female', -- 女
    'male', -- 男
    'other', -- 其他
);

CREATE TYPE productable_type_enum AS ENUM(
    'sku', -- 绑定sku
    'spu', -- 绑定商品 方便join表查询标识
);

CREATE TYPE imageable_type_enum AS ENUM(
    'spu',
    'sku',
    'user_avatar',
    ''
);

-- ui input 输入框类型
CREATE TYPE ui_input_type_enum AS ENUM(
    'input_string', -- 普通输入框 具体是数字还是字符串 需要参照value_type的值
    'input_number', -- 普通输入框 具体是数字还是字符串 需要参照value_type的值
    'textarea', -- 大文本输入
    'single_choice', -- 单选下拉框 或者 radio input
    'multiple_choice', -- 可以多选的下拉框 或者  checkbox
    "object," -- 对象
    "list_object", -- 对象数组
);

-- ui input 输入框类型
CREATE TYPE value_type_enum AS ENUM(
    'string', -- 字符串
    'number', -- 数字
    'bool' -- 布尔
);

-- 图片大小
CREATE TYPE image_size_enum AS ENUM(
    'small',
    'medium',
    'large',
);

-- 权限 action
CREATE TYPE permission_action_enum AS ENUM(
    'read',
    'write',
    'all',
    ''
);

-- 范围操作符
CREATE TYPE operating_range_enum AS ENUM(
    'including', -- 包含
    'excluding', -- 排除
    ''
);

CREATE TYPE store_opening_status_enum AS ENUM(
    'open', -- 开启
    'close', -- 关闭
    'disalbe', -- 禁用
);

CREATE TYPE warehouse_type_enum AS ENUM(
    'front', -- 前置仓库
    'district', -- 大区仓库
    'overseas', -- 海外仓库
);

CREATE TYPE payment_status_enum AS ENUM(
    'buyer_pending_payment',
    'buyer_confirms_payment',
    'pay_to_seller',
    'refund_to_buyer',
)
CREATE TYPE order_type_enum AS ENUM(
    'buy',
    'sell',
)
-- SELECT table_name, column_name
-- FROM information_schema.columns
-- WHERE table_name IN (
--     SELECT table_name
--     FROM information_schema.columns
--     WHERE data_type = 'image_size'
-- );
-- ALTER TABLE your_table
-- ALTER COLUMN your_enum_column TYPE new_permissions_action_enum
-- USING image_size:varchar:image_size;
