-- name: PutArticle :exec
INSERT INTO "articles" (
    "id"
    ,"title"
    ,"body"
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
)
ON CONFLICT ("id") DO UPDATE
SET "title" = EXCLUDED.title
    ,"body" = EXCLUDED.body
    ,"thumbnail" = EXCLUDED.thumbnail
    ,"updated_at" = EXCLUDED.updated_at;

-- name: CreateTempTagsTable :exec
CREATE TEMP TABLE IF NOT EXISTS tmp_tags (
    id VARCHAR(144),
    article_id VARCHAR(26),
    name VARCHAR(35) NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id, article_id)
) ON COMMIT PRESERVE ROWS;

-- name: PreAttachTags :copyfrom
INSERT INTO "tmp_tags" (
    "id"
    ,"article_id"
    ,"name"
    ,"created_at"
    ,"updated_at"
) VALUES (
    $1
    ,$2
    ,$3
    ,$4
    ,$5
);

-- name: AttachTags :exec
WITH "inserted" AS (
    INSERT INTO "tags" (
        "id"
        ,"article_id"
        ,"name"
        ,"created_at"
        ,"updated_at"
    )
    SELECT id, article_id, name, created_at, updated_at FROM "tmp_tags" WHERE "tmp_tags"."article_id" = $1
    ON CONFLICT ("id","article_id") DO UPDATE
    SET "updated_at" = EXCLUDED.updated_at
    RETURNING "id"
)
DELETE
FROM
    "tags"
WHERE
    "tags"."article_id" = $1
AND
    "tags"."id" NOT IN (SELECT id FROM "inserted");