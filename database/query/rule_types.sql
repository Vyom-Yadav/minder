-- name: CreateRuleType :one
INSERT INTO rule_type (
    name,
    project_id,
    description,
    guidance,
    definition,
    severity_value,
    subscription_id,
    display_name,
    release_phase,
    short_failure_message
) VALUES (
    $1,
    $2,
    $3,
    $4,
    sqlc.arg(definition)::jsonb,
    sqlc.arg(severity_value),
    sqlc.narg(subscription_id),
    sqlc.arg(display_name),
    sqlc.arg(release_phase),
    sqlc.arg(short_failure_message)
) RETURNING *;

-- name: ListRuleTypesByProject :many
SELECT * FROM rule_type WHERE project_id = $1;

-- name: GetRuleTypeByID :one
SELECT * FROM rule_type WHERE id = $1;

-- name: GetRuleTypeByName :one
SELECT * FROM rule_type WHERE  project_id = ANY(sqlc.arg(projects)::uuid[]) AND lower(name) = lower(sqlc.arg(name));

-- name: DeleteRuleType :exec
DELETE FROM rule_type WHERE id = $1;

-- name: UpdateRuleType :one
UPDATE rule_type
    SET description = $2, definition = sqlc.arg(definition)::jsonb, severity_value = sqlc.arg(severity_value), display_name = sqlc.arg(display_name), release_phase = sqlc.arg(release_phase), short_failure_message = sqlc.arg(short_failure_message)
    WHERE id = $1
    RETURNING *;

-- name: GetRuleTypesByEntityInHierarchy :many
SELECT rt.* FROM rule_type AS rt
JOIN rule_instances AS ri ON ri.rule_type_id = rt.id
WHERE ri.entity_type = $1
AND ri.project_id = ANY(sqlc.arg(projects)::uuid[]);

-- intended as a temporary transition query
-- this will be removed once the evaluation history tables replace the old state tables
-- name: GetRuleTypeNameByID :one
SELECT name FROM rule_type
WHERE id = $1;
