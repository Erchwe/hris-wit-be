-- name: CreateProjectMember :one
INSERT INTO project_members (
    project_memb_id, id_data_pegawai, project_id, project_role
) VALUES (
    @project_memb_id, @id_data_pegawai, @project_id, @project_role::project_role_enum
) RETURNING *;

-- name: GetProjectMemberByID :one
SELECT * FROM project_members
WHERE project_memb_id = @project_memb_id;

-- name: ListProjectMembers :many
SELECT * FROM project_members;

-- name: UpdateProjectMember :one
UPDATE project_members
SET
    project_role = @project_role::project_role_enum
WHERE project_memb_id = @project_memb_id
RETURNING *;

