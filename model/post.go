package model

/*
create table ncut_bbs.posts
(
	id int auto_increment
		primary key,
	author_id int null,
	create_time int null,
	update_time int null,
	title text null,
	views int null,
	likes int null,
	content text null,
	unlikes int null,
	pictures text null,
	constraint posts_users_id_fk
		foreign key (author_id) references ncut_bbs.users (id)
);
*/

type Post struct {
	ID         uint
	AuthorID   uint
	CreateTime int64 `gorm:"autoCreateTime"`
	UpdateTime int64 `gorm:"autoUpdateTime"`
	Title      string
	Views      int
	Likes      int
	Content    string
	Unlikes    int
	Pictures   string
}
