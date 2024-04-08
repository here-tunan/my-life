# my-life
一个生活管理系统，目前支持多人一起记账和查账。
## 功能截图
### 登录页面
![image](https://github.com/here-tunan/my-life/assets/40956738/ead55814-d91f-49aa-a759-a88e7776f81a)

### 首页 （个人基本信息和家庭基本信息）
![image](https://github.com/here-tunan/my-life/assets/40956738/d5caa1fa-76e9-4c19-b8b6-9a0d07f50c45)

### 记账页面
![image](https://github.com/here-tunan/my-life/assets/40956738/3df447ab-33e8-4e3b-b775-fdc8a03858e2)

### 快速手工记账
![image](https://github.com/here-tunan/my-life/assets/40956738/e399307b-e31b-443f-bf80-bb6efcc133bf)

### 支持微信/支付宝账单导入（且导入时会根据描述信息，基于ES的分词功能和历史记账信息匹配出最优类型）
![image](https://github.com/here-tunan/my-life/assets/40956738/6e9b891c-03ee-4bfb-9d85-7e6f7552cfb3)

### 个人收支统计
![image](https://github.com/here-tunan/my-life/assets/40956738/d0a0cc80-eadc-4227-8bb9-b34ba27a27bb)

### 家庭收支统计（可以看到家庭内的其他成员的收支明细）
![image](https://github.com/here-tunan/my-life/assets/40956738/e3c28bef-d02a-474c-b9b1-da3c75a6dac4)


## Mysql 建表语句
```sql
create table user
(
    id           bigint auto_increment
        primary key,
    account      varchar(20)  not null,
    password     varchar(50)  not null,
    name         varchar(255) null,
    `desc`       varchar(255) null,
    avatar       varchar(200) null comment '头像',
    extra        varchar(255) null,
    is_deleted   tinyint(1)   null,
    gmt_create   datetime     null,
    gmt_modified datetime     null
);

create table family
(
    id           bigint auto_increment comment '主键id'
        primary key,
    name         varchar(100) null comment '家庭名称',
    `desc`       varchar(400) not null comment '家庭描述',
    avatar       varchar(200) null comment '头像',
    gmt_create   timestamp    null,
    gmt_modified timestamp    null,
    is_deleted   tinyint(1)   null
)
    comment '家庭表';


create table family_member
(
    id           bigint auto_increment comment '自增id'
        primary key,
    family_id    bigint     null,
    user_id      bigint     null,
    is_creator   tinyint(1) null comment '是这个家庭的创建者么',
    gmt_create   timestamp  null,
    gmt_modified timestamp  null,
    is_deleted   tinyint(1) null
)
    comment '家庭成员表';


create table transaction
(
    id           bigint auto_increment
        primary key,
    amount       decimal(10, 2) null,
    description  varchar(255)   null,
    user_id      bigint         null,
    type         int            null,
    category     int            null,
    account      int            null,
    time         datetime       null,
    is_deleted   tinyint(1)     null,
    gmt_create   datetime       null,
    gmt_modified datetime       null
);

create table transaction_account
(
    id           bigint auto_increment
        primary key,
    name         varchar(255) null,
    `desc`       varchar(255) null,
    is_deleted   tinyint(1)   null,
    gmt_create   datetime     null,
    gmt_modified datetime     null
)
    charset = utf8mb3;

create table transaction_category
(
    id           bigint auto_increment
        primary key,
    name         varchar(255) charset utf8mb3 null,
    type         int                          null,
    `desc`       varchar(255) charset utf8mb3 null,
    is_deleted   tinyint(1)                   null,
    gmt_create   datetime                     null,
    gmt_modified datetime                     null
);
```

## Mysql 基础数据
```sql
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (1, '微信支付', '微信支付', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (2, '支付宝', '支付宝', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (3, '现金', '现金', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (4, '美团', '美团', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (5, '医保/保险', '医保/保险', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (6, '中国银行', '中国银行', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (7, '工商银行', '工商银行', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (8, '建设银行', '建设银行', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (9, '农业银行', '建设银行', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (10, '招商银行', '招商银行', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (11, '杭州银行', '杭州银行', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (12, '关爱通', '关爱通', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');
insert into life.transaction_account (id, name, `desc`, is_deleted, gmt_create, gmt_modified) values (101, '其他', '其他', 0, '2023-09-18 14:55:52', '2023-09-18 14:55:52');

insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (1, '饮食', 1, '平日里的饮食消费（粮油菜也算）', 0, '2023-09-13 15:05:37', '2023-09-13 15:05:40');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (2, '生活用品', 1, '购买一些生活用品', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (3, '保险医疗', 1, '用于保险、治病、买药等的支出', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (4, '运动健身', 1, '用于买运动器械装备、课程等', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (5, '聚会娱乐', 1, '平日的聚会、旅游、购物，用于娱乐项目的支出等', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (6, '成长教育', 1, '用于提升自我的一些消费，买书买课等', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (7, '水电煤气费', 1, '水电煤气费', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (8, '话费通讯', 1, '流量、话费等充值', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (9, '电子产品', 1, '手机、耳机、手表、笔记本等电子产品购买', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (10, '房租', 1, '房租', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (11, '运输交通', 1, '寄快递、打车、公共交通', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (12, '人情往来', 1, '发红包、随礼、压岁钱等', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (13, '孝敬父母', 1, '给家里买东西、给钱等支出', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (14, '工资收入', 2, '工资收入', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (15, '红包收入', 2, '收到的一些红包', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (16, '返现收入', 2, '购物的一些返现', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (17, '补贴收入', 2, '政府的补贴、失业金等的收入', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (101, '其他支出', 1, '其他支出', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');
insert into life.transaction_category (id, name, type, `desc`, is_deleted, gmt_create, gmt_modified) values (102, '其他收入', 2, '其他收入', 0, '2023-09-18 16:32:55', '2023-09-18 16:33:02');

insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 123.00, '1231', 2, 2, 2, 1, '2024-02-01 00:00:00', 0, '2024-02-06 16:51:44', '2024-02-06 16:51:44');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 123.00, '1231', 2, 2, 2, 1, '2024-02-01 00:00:00', 0, '2024-02-06 16:50:17', '2024-02-06 16:50:17');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 123.00, '1231', 2, 2, 2, 1, '2024-02-01 00:00:00', 0, '2024-02-06 16:22:17', '2024-02-06 16:22:17');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 123.00, '1231', 2, 2, 2, 1, '2024-02-01 00:00:00', 0, '2024-02-06 16:21:29', '2024-02-06 16:21:29');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 55.00, 'sheng sheng 生活！！！', 2, 2, 2, 3, '2024-02-18 00:00:00', 0, '2024-02-06 13:16:30', '2024-02-06 16:15:42');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 567.00, '测试', 2, 2, 2, 3, '2024-02-18 00:00:00', 1, '2024-02-06 13:16:30', '2024-02-06 13:16:53');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 86.00, '好好啊哈', 2, 2, 2, 3, '2024-02-18 00:00:00', 0, '2024-02-06 13:16:30', '2024-02-06 13:16:30');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 867.00, '电饭锅', 2, 2, 3, 1, '2024-02-08 00:00:00', 0, '2024-02-06 13:13:50', '2024-02-06 13:13:50');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 465.00, '官方_环境_有点好', 1, 2, 3, 1, '2024-02-08 00:00:00', 0, '2024-02-06 13:13:50', '2024-02-06 13:14:20');
insert into life.transaction ( amount, description, user_id, type, category, account, time, is_deleted, gmt_create, gmt_modified) values ( 234.00, 'fghj', 2, 2, 3, 1, '2024-02-08 00:00:00', 0, '2024-02-06 13:13:50', '2024-02-06 13:13:50');

```