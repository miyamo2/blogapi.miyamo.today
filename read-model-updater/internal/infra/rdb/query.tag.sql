-- name: CreateTempTagsTable :exec
CREATE TEMP TABLE IF NOT EXISTS tmp_tags (
    id VARCHAR(144),
    name VARCHAR(35) NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id)
) ON COMMIT PRESERVE ROWS;

-- name: PrePutTags :copyfrom
INSERT INTO "tmp_tags" (
    "id"
    ,"name"
    ,"created_at"
    ,"updated_at"
)
VALUES (
    $1
    ,$2
    ,$3
    ,$4
);

-- name: PutTags :exec
INSERT INTO "tags" (
    "id"
    ,"name"
    ,"created_at"
    ,"updated_at"
)
SELECT id, name, created_at, updated_at FROM "tmp_tags" WHERE "tmp_tags"."id" = ANY($1::varchar[])
    ON CONFLICT DO NOTHING;

-- name: CreateTempArticlesTable :exec
CREATE TEMP TABLE IF NOT EXISTS tmp_articles (
    id VARCHAR(26),
    tag_id VARCHAR(144),
    title VARCHAR(255) NOT NULL,
    thumbnail VARCHAR(524271),
    created_at timestamp WITH TIME ZONE NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id, tag_id)
) ON COMMIT PRESERVE ROWS;

-- name: PrePutArticle :copyfrom
INSERT INTO "tmp_articles" (
    "id"
    ,"tag_id"
    ,"title"
    ,"thumbnail"
    ,"created_at"
    ,"updated_at"
)
VALUES (
    $1
    ,$2
    ,$3
    ,$4
    ,$5
    ,$6
);

-- name: PutArticle :exec
WITH "inserted" AS (
    INSERT INTO "articles" (
        "id"
        ,"tag_id"
        ,"title"
        ,"thumbnail"
        ,"created_at"
        ,"updated_at"
    )
    SELECT id, tag_id, title, thumbnail, created_at, updated_at FROM "tmp_articles" WHERE "tmp_articles"."id" = $1
    ON CONFLICT ("id","tag_id") DO UPDATE
        SET "title" = EXCLUDED.title
        ,"thumbnail" = EXCLUDED.thumbnail
        ,"updated_at" = EXCLUDED.updated_at
    RETURNING "tag_id"
)
DELETE
FROM
    "articles"
WHERE
    "articles"."id" = $1
AND
    "articles"."tag_id" NOT IN (SELECT "tag_id" FROM "inserted");