ALTER TABLE
    blog_contents
ADD
    COLUMN page_type smallint DEFAULT 0;

COMMENT ON COLUMN blog_contents.page_type IS 'ページの種類 0: ページ, 1: プライバシーポリシー, 2: 運営者情報';