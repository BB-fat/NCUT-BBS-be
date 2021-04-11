package model

/*
create table ncut_bbs.verify_info
(
	id int auto_increment
		primary key,
	user_id int not null,
	image varchar(100) null,
	remark text null,
	pass int null,
	constraint verify_info_users_id_fk
		foreign key (user_id) references ncut_bbs.users (id)
);
*/

type VerifyInfo struct {
	ID     uint
	UserID uint
	Image  string
	Remark string
	Pass   uint
}
