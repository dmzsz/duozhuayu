-- 订单信息
CREATE TABLE "order"(
    id int8 PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    is_deleted bool NOT NULL DEFAULT FALSE,
    order_no int8 NOT NULL,
    amount decimal(10, 2) NULL,
    discount_total decimal(10, 2) NULL,
    payment_status payment_status_enum NULL,
    paymented_at timestamp NULL,
    payment_type varchar(32) NULL,
    is_anonyous bool FALSE,
    order_type order_type_enum NOT NULL DEFAULT 'buy',
    finished_at timestamp NULL,
    user_id int8 NOT NULL,
    address_id int8 NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    FOREIGN KEY (address_id) REFERENCES contact_address(id)
);

-- 为表添加注释
COMMENT ON TABLE "order" IS '物流信息表';

COMMENT ON COLUMN "order".order_no IS '订单编号';

COMMENT ON COLUMN "order".amount IS '实付款';

COMMENT ON COLUMN "order".discount_total IS '优惠总金额';

COMMENT ON COLUMN "order".payment_status IS '支付状态';

COMMENT ON COLUMN "order".paymented_at IS '支付啊hi见';

COMMENT ON COLUMN "order".payment_type IS '支付方式';

COMMENT ON COLUMN "order".is_anonyous IS '是否匿名';

COMMENT ON COLUMN "order".order_type IS '订单类型 buy(购买)还是sell(出售) 默认buy';

COMMENT ON COLUMN "order".finished_at IS '关闭订单时间';

