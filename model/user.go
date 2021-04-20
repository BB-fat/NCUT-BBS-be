package model

/*
create table ncut_bbs.users
(
	id int auto_increment
		primary key,
	account_name varchar(50) null,
	password varchar(50) null,
	real_name varchar(30) null,
	sex int null comment '0:男 1:女',
	college varchar(50) null comment '学院',
	account_status int default 0 not null comment '0:未实名 1:待审核 2:实名未通过 3:激活 4:被封禁',
	grade int null comment '年级',
	create_time int not null,
	avatar text null,
	constraint user_account_name_uindex
		unique (account_name)
)
comment '用户信息';
*/

type User struct {
	ID            uint
	AccountName   string
	Avatar        string
	Password      string
	RealName      string
	Sex           int
	College       string
	AccountStatus int
	Grade         int
	CreateTime    int64 `gorm:"autoCreateTime"`
}
