CREATE TABLE IF NOT EXISTS system_groups (
    id serial,
    name character varying(50),
    edit_blog smallint DEFAULT 0,
    edit_category smallint DEFAULT 0,
    edit_tag smallint DEFAULT 0,
    edit_user smallint DEFAULT 0,
    admin smallint DEFAULT 0,
    sort_order integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT system_groups_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN system_groups.id IS 'ID';

COMMENT ON COLUMN system_groups.name IS 'グループ名';

COMMENT ON COLUMN system_groups.edit_blog IS 'ブログ編集権限 0:無 1:作成有り 2:作成・承認有り';

COMMENT ON COLUMN system_groups.edit_category IS 'カテゴリ編集権限 0:無 1:有り';

COMMENT ON COLUMN system_groups.edit_tag IS 'タグ編集権限 0:無 1:有り';

COMMENT ON COLUMN system_groups.edit_user IS 'ユーザー編集権限 0:無 1:有り';

COMMENT ON COLUMN system_groups.admin IS '管理者権限 0:無 1:有り';

COMMENT ON COLUMN system_groups.sort_order IS '並び順';

-- INSERT INTO system_groups(name, edit_blog, edit_category, edit_tag, edit_user, admin, sort_order) VALUES ('admin', 2, 1, 1, 1, 1, 1);