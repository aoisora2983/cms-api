CREATE TABLE IF NOT EXISTS system_user_links (
    id serial,
    user_id integer NOT NULL,
    url text,
    icon_path text,
    alt text,
    sort_order integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT system_user_links_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN system_user_links.id IS 'ID';

COMMENT ON COLUMN system_user_links.user_id IS 'ユーザーID ref) system_users.id';

COMMENT ON COLUMN system_user_links.url IS 'リンク先';

COMMENT ON COLUMN system_user_links.icon_path IS 'リンクアイコンのパス';

COMMENT ON COLUMN system_user_links.alt IS 'アイコンのaltテキスト';

COMMENT ON COLUMN system_user_links.sort_order IS '並び順';