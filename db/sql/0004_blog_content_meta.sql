CREATE TABLE IF NOT EXISTS blog_content_meta (
    id serial,
    id_blog_content integer NOT NULL,
    meta_key character varying(300),
    meta_value text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT blog_content_meta_pkey PRIMARY KEY (id, id_blog_content)
);

COMMENT ON COLUMN blog_content_meta.id IS 'ID';

COMMENT ON COLUMN blog_content_meta.id_blog_content IS '記事ID';

COMMENT ON COLUMN blog_content_meta.meta_key IS 'カスタムキー';

COMMENT ON COLUMN blog_content_meta.meta_value IS 'カスタム値';