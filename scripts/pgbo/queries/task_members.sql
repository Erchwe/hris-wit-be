-- name: CreateTaskMember :one
INSERT INTO task_members (
    task_memb_id, project_memb_id, task_id
) VALUES (
    @task_memb_id, @project_memb_id, @task_id
) RETURNING *;

-- name: GetTaskMemberByID :one
SELECT * FROM task_members
WHERE task_memb_id = @task_memb_id;

-- name: ListTaskMembers :many
SELECT * FROM task_members;

