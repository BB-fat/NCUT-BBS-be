package model

/*
create table ncut_bbs.post_likes
(
	post_id int null,
	user_id int null,
	constraint post_likes_posts_id_fk
		foreign key (post_id) references ncut_bbs.posts (id),
	constraint post_likes_users_id_fk
		foreign key (user_id) references ncut_bbs.users (id)
);
*/

type PostLike struct {
	UserID uint
	PostID uint
}
