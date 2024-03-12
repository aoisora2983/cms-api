CREATE TABLE IF NOT EXISTS blog_contents (
    id integer NOT NULL,
    id_branch integer NOT NULL,
    id_user integer NOT NULL,
    title character varying(300),
    content text,
    description text,
    status smallint DEFAULT 0,
    thumbnail text,
    published_start_time timestamp with time zone NOT NULL,
    published_end_time timestamp with time zone,
    published_updated_time timestamp with time zone,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT blog_contents_pkey PRIMARY KEY (id, id_branch)
);

COMMENT ON COLUMN blog_contents.id IS 'ID';

COMMENT ON COLUMN blog_contents.id_user IS '作成者ID';

COMMENT ON COLUMN blog_contents.title IS '記事タイトル';

COMMENT ON COLUMN blog_contents.content IS '記事内容';

COMMENT ON COLUMN blog_contents.description IS '記事概要';

COMMENT ON COLUMN blog_contents.status IS '記事状態: 下書き, 公開中, etc...';

COMMENT ON COLUMN blog_contents.thumbnail IS 'サムネイル, etc...';

COMMENT ON COLUMN blog_contents.published_start_time IS '公開開始時間';

COMMENT ON COLUMN blog_contents.published_end_time IS '公開終了時間';

COMMENT ON COLUMN blog_contents.published_end_time IS '更新時間';