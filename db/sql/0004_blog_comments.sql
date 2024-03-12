CREATE TABLE IF NOT EXISTS blog_comments (
    id serial,
    id_blog_content integer NOT NULL,
    id_replay int,
    user_name character varying(300),
    comment text,
    ip text,
    good int default 0,
    status smallint default 0,
    comment_time timestamp with time zone,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT blog_comments_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN blog_comments.id IS 'ID';

COMMENT ON COLUMN blog_comments.id_blog_content IS 'コンテンツID';

COMMENT ON COLUMN blog_comments.id_replay IS '返信先ID';

COMMENT ON COLUMN blog_comments.user_name IS '投稿者名';

COMMENT ON COLUMN blog_comments.comment IS 'コメント内容';

COMMENT ON COLUMN blog_comments.ip IS '投稿者IP';

COMMENT ON COLUMN blog_comments.status IS 'コメント状態';

COMMENT ON COLUMN blog_comments.comment_time IS 'コメント投稿時刻';