CREATE TABLE IF NOT EXISTS sub_sites (
    id serial,
    subdir_name text NOT NULL UNIQUE,
    icon_text   text NOT NULL,
    status smallint DEFAULT 0,
    title text,
    description text,
    sort_order integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT sub_sites_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE sub_sites IS 'サブサイト管理テーブル';

COMMENT ON COLUMN sub_sites.id IS 'ID';

COMMENT ON COLUMN sub_sites.subdir_name IS 'サブディレクトリ名';

COMMENT ON COLUMN sub_sites.icon_text IS 'アイコン名称';

COMMENT ON COLUMN sub_sites.status IS '状態 0: 非公開 1: 公開';

COMMENT ON COLUMN sub_sites.title IS 'サブサイトタイトル';

COMMENT ON COLUMN sub_sites.description IS 'サブサイト概要';

COMMENT ON COLUMN sub_sites.sort_order IS 'サブサイト並び順';