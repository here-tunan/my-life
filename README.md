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