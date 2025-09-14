-- name: GetByID :one
SELECT
    "t".*,
    CAST(
        COALESCE(
            jsonb_agg(
                json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)
            ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
        ) AS json
    ) AS "articles"
FROM
    (SELECT * FROM "tags" WHERE "tags"."id" = $1) AS "t"
LEFT OUTER JOIN
    "articles" AS "a"
ON
    "t"."id" = "a"."tag_id"
GROUP BY "t"."id";

-- name: ListAfter :many
SELECT
    "t".*,
    CAST(
        COALESCE(
            jsonb_agg(
                json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)
            ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
        ) AS json
    ) AS "articles"
FROM
    (SELECT * FROM "tags" ORDER BY "tags"."id") AS "t"
LEFT OUTER JOIN
    "articles" AS "a"
ON
    "t"."id" = "a"."tag_id"
GROUP BY
    "t"."id"
ORDER BY
    "t"."id";

-- name: ListAfterWithLimit :many
SELECT
    "t".*,
    CAST(
        COALESCE(
            jsonb_agg(
                json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)
            ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
        ) AS json
    ) AS "articles"
FROM
    (SELECT * FROM "tags" ORDER BY "tags"."id" LIMIT $1) AS "t"
LEFT OUTER JOIN
    "articles" AS "a"
ON
    "t"."id" = "a"."tag_id"
GROUP BY
    "t"."id"
ORDER BY
    "t"."id";

-- name: ListAfterWithLimitAndCursor :many
SELECT
    "t".*,
    CAST(
        COALESCE(
            jsonb_agg(
                json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)
            ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
        ) AS json
    ) AS "articles"
FROM
    (
        SELECT
            *
        FROM
            "tags"
        WHERE
            EXISTS(
                SELECT id FROM "tags" WHERE "tags"."id" = $1
            )
        AND
            "tags"."id" > $1
        ORDER BY
            "tags"."id"
        LIMIT $2
    ) AS "t"
LEFT OUTER JOIN
    "articles" AS "a"
ON
    "t"."id" = "a"."tag_id"
GROUP BY
    "t"."id"
ORDER BY
    "t"."id";

-- name: ListBefore :many
SELECT
    "t".*,
    CAST(
        COALESCE(
            jsonb_agg(
                json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)
            ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
        ) AS json
    ) AS "articles"
FROM
    (SELECT * FROM "tags" ORDER BY "tags"."id" DESC) AS "t"
LEFT OUTER JOIN
    "articles" AS "a"
ON
    "t"."id" = "a"."tag_id"
GROUP BY
    "t"."id"
ORDER BY
    "t"."id" DESC;

-- name: ListBeforeWithLimit :many
SELECT
    "t".*,
    CAST(
        COALESCE(
            jsonb_agg(
                json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)
            ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
        ) AS json
    ) AS "articles"
FROM
    (SELECT * FROM "tags" ORDER BY "tags"."id" DESC LIMIT $1) AS "t"
LEFT OUTER JOIN
    "articles" AS "a"
ON
    "t"."id" = "a"."tag_id"
GROUP BY
    "t"."id"
ORDER BY
    "t"."id" DESC;

-- name: ListBeforeWithLimitAndCursor :many
SELECT
    "t".*,
    CAST(
        COALESCE(
            jsonb_agg(
                json_build_object('id', a.id, 'title', a.title, 'thumbnail', a.thumbnail, 'created_at', a.created_at, 'updated_at', a.updated_at)
            ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
        ) AS json
    ) AS "articles"
FROM
    (
        SELECT
            *
        FROM
            "tags"
        WHERE
            EXISTS(
                SELECT id FROM "tags" WHERE "tags"."id" = $1
            )
        AND
            "tags"."id" < $1
        ORDER BY
            "tags"."id" DESC
        LIMIT $2
    ) AS "t"
LEFT OUTER JOIN
    "articles" AS "a"
ON
    "t"."id" = "a"."tag_id"
GROUP BY
    "t"."id"
ORDER BY
    "t"."id" DESC;