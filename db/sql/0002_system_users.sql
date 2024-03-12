CREATE TABLE IF NOT EXISTS system_users (
    id serial,
    group_id integer NOT NULL,
    name character varying(50),
    description text,
    mail text,
    icon_path text,
    password character varying(50),
    sort_order integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT system_users_pkey PRIMARY KEY (id)
);

COMMENT ON COLUMN system_users.id IS 'ID';

COMMENT ON COLUMN system_users.group_id IS '所属ID';

COMMENT ON COLUMN system_users.name IS 'ユーザ名';

COMMENT ON COLUMN system_users.description IS 'ユーザ概要';

COMMENT ON COLUMN system_users.mail IS 'メールアドレス';

COMMENT ON COLUMN system_users.mail IS 'アイコンパス';

COMMENT ON COLUMN system_users.password IS 'パスワード';

COMMENT ON COLUMN system_users.sort_order IS '並び順';

-- INSERT INTO system_users (id, group_id, name, description, mail, password, sort_order) VALUES (1, 1, 'Admin', '管理ユーザー', 'admin@example.com', '123456', 1);
