package model

/*
create table ncut_bbs.post_comments
(
	id int auto_increment
		primary key,
	author_id int null,
	post_id int null,
	create_time int null,
	content text null,
	likes int null,
	unlikes int null,
	pictures text null,
	constraint post_comments_posts_id_fk
		foreign key (post_id) references ncut_bbs.posts (id),
	constraint post_comments_users_id_fk
		foreign key (author_id) references ncut_bbs.users (id)
);
*/

type PostComment struct {
	ID         uint
	AuthorID   uint
	PostID     uint
	CreateTime int64 `gorm:"autoCreateTime"`
	Content    string
	Likes      int
	Unlikes    int
	Pictures   string
}
