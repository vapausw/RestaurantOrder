create database if not exists restaurant;
use restaurant;

drop table if exists users;

create table `users`
(
    `id`        bigint(20) unsigned not null auto_increment primary key comment '主键',
    `user_id`   bigint(20) unsigned not null unique comment '用户id',
    `email`     varchar(50)         not null unique comment '邮箱',
    `password`  varchar(255)        not null comment '密码',
    `nick_name` varchar(50)         not null comment '昵称'
) comment '用户表';

drop table if exists merchants;
create table `merchants` (
    id bigint(20) unsigned not null auto_increment primary key comment '主键',
    merchant_id bigint(20) unsigned not null unique comment '商户id',
    merchant_email varchar(50) not null unique comment '商户邮箱',
    merchant_password varchar(255) not null comment '商户密码'
) comment '商户表';

drop table if exists shops;
create table `shops`
(
    `id`           bigint(20) unsigned not null auto_increment primary key comment '主键',
    `shop_id`      bigint(20) unsigned not null unique comment '商户id',
    `shop_name`    varchar(50)         not null comment '商户名称',
    `shop_address` varchar(255)        not null comment '商户地址',
    `shop_phone`   varchar(50)         not null comment '商户电话',
    `shop_desc`    longtext            not null comment '商户描述',
    `create_time`  timestamp           not null default current_timestamp comment '创建时间',
    `update_time`  timestamp           not null default current_timestamp on update current_timestamp comment '更新时间',
    foreign key (shop_id) references `merchants` (merchant_id)
) comment '商户信息表';

insert into shops(shop_id, shop_name, shop_address, shop_phone, shop_desc) values (1, '麦当劳', '北京市朝阳区', '010-12345678', '麦当劳是一家全球连锁快餐公司，总部位于美国伊利诺伊州芝加哥。');
insert into shops(shop_id, shop_name, shop_address, shop_phone, shop_desc) values (2, '肯德基', '北京市海淀区', '010-87654321', '肯德基是一家美国跨国连锁餐饮企业，总部位于美国肯塔基州路易斯维尔市。');


drop table if exists menus;
create table `menus`
(
    id          bigint(20) unsigned not null auto_increment primary key comment '主键',
    menu_id     bigint(20) unsigned not null unique comment '菜单id',
    shop_id     bigint(20) unsigned not null comment '商户id',
    menu_name   varchar(50)         not null comment '菜单名称',
    menu_price  decimal(10, 2)      not null comment '菜单价格',
    menu_desc   longtext            not null comment '菜单描述',
    menu_stock  int                 not null comment '菜单库存',
    create_time timestamp           not null default current_timestamp comment '创建时间',
    update_time timestamp           not null default current_timestamp on update current_timestamp comment '更新时间'
) comment '菜单表';


insert into menus(menu_id, shop_id, menu_name, menu_price, menu_desc, menu_stock) values (1, 1, '麦辣鸡腿堡', 18.0, '麦辣鸡腿堡是麦当劳的招牌汉堡之一。', 100);
insert into menus(menu_id, shop_id, menu_name, menu_price, menu_desc, menu_stock) values (2, 1, '薯条', 8.0, '薯条是麦当劳的招牌小吃之一。', 1000);
insert into menus(menu_id, shop_id, menu_name, menu_price, menu_desc, menu_stock) VALUES (3, 2, '香辣鸡腿堡', 20.0, '香辣鸡腿堡是肯德基的招牌汉堡之一。', 100);
insert into menus(menu_id, shop_id, menu_name, menu_price, menu_desc, menu_stock) VALUES (4, 2, '薯条', 10.0, '薯条是肯德基的招牌小吃之一。', 1000);

drop table if exists carts;
-- 购物车表 查询需要进行优化，因为用户id以及菜单id都不是唯一的，但这两个字段的组合是唯一的
-- 所以可以将这两个字段组合成一个唯一索引
create table `carts`
(
    id          bigint(20) unsigned not null auto_increment primary key comment '主键',
    menu_id     bigint(20) unsigned not null comment '菜单id',
    user_id     bigint(20) unsigned not null comment '用户id',
    menu_count  int                 not null comment '菜单数量',
    price       int                 not null comment '菜单价格',
    note        varchar(255)        not null comment '备注',
    create_time timestamp           not null default current_timestamp comment '创建时间',
    update_time timestamp           not null default current_timestamp on update current_timestamp comment '更新时间'
) comment '购物车表';
drop table if exists `orders`;
create table `orders`
(
    id       bigint(20) unsigned not null auto_increment primary key comment '主键',
    order_id bigint(20) unsigned not null unique comment '订单id',
    user_id  bigint(20) unsigned not null comment '用户id',
    shop_id  bigint(20) unsigned not null comment '商户id'
) comment '订单表';

drop table if exists `order_info`;
create table `order_info`
(
    id       bigint(20) unsigned not null auto_increment primary key comment '主键',
    order_id bigint(20) unsigned not null comment '订单id',
    menu_id  bigint(20) unsigned not null comment '菜单id',
    count    int                 not null comment '菜单数量',
    price    decimal(10, 2)      not null comment '菜单价格',
    note     varchar(255)        not null comment '备注',
    create_time timestamp       not null default current_timestamp comment '创建时间',
    update_time timestamp       not null default current_timestamp on update current_timestamp comment '更新时间',
    foreign key (order_id) references `orders` (order_id)
) comment '订单详情表';




