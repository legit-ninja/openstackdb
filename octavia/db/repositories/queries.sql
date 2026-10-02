-- name: PoolGetAll :many
SELECT
    id,
    project_id,
    name,
    protocol,
    lb_algorithm,
    operating_status,
    load_balancer_id,
    provisioning_status
FROM
    pool
WHERE
    provisioning_status != 'DELETED';

-- name: LoadBalancerGetAllWithVIP :many
SELECT
    lb.id,
    lb.project_id,
    lb.name,
    lb.provisioning_status,
    lb.operating_status,
    lb.provider,
    v.ip_address as vip_address
FROM
    load_balancer lb
    LEFT JOIN vip v ON lb.id = v.load_balancer_id
WHERE
    lb.provisioning_status != 'DELETED';

-- name: AmphoraGetAll :many
SELECT
    id,
    compute_id,
    status,
    load_balancer_id,
    lb_network_ip,
    ha_ip,
    role,
    cert_expiration
FROM
    amphora
WHERE
    status != 'DELETED';

-- name: QuotaGetAll :many
-- Per-project Octavia quota limits. NULL column means no explicit override.
SELECT
    project_id,
    health_monitor,
    listener,
    load_balancer,
    member,
    pool,
    l7policy,
    l7rule
FROM
    quotas;

-- name: LoadBalancerCountByProject :many
SELECT
    project_id,
    COUNT(*) as count
FROM
    load_balancer
WHERE
    provisioning_status != 'DELETED'
GROUP BY
    project_id;

-- name: PoolCountByProject :many
SELECT
    project_id,
    COUNT(*) as count
FROM
    pool
WHERE
    provisioning_status != 'DELETED'
GROUP BY
    project_id;

-- name: ListenerCountByProject :many
SELECT
    project_id,
    COUNT(*) as count
FROM
    listener
WHERE
    provisioning_status != 'DELETED'
GROUP BY
    project_id;
