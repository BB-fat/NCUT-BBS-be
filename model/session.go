package model

/*
create table ncut_bbs.sessions
(
	id int auto_increment
		primary key,
	user_id int not null,
	update_time int not null,
	token varchar(100) not null,
	constraint sessions_token_uindex
		unique (token),
	constraint sessions_user_id_uindex
		unique (user_id),
	constraint sessions_users_id_fk
		foreign key (user_id) references ncut_bbs.users (id)
);


*/

type Session struct {
	ID         uint
	UserID     uint
	UpdateTime int64 `gorm:"autoUpdateTime"`
	Token      string
}
