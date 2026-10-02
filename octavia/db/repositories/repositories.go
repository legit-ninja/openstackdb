package octavia

import "context"

// PoolRepository provides read operations for pools.
type PoolRepository struct{ queries *Queries }

func NewPoolRepository(db DBTX) *PoolRepository { return &PoolRepository{queries: New(db)} }
func (r *PoolRepository) GetAll(ctx context.Context) ([]PoolGetAllRow, error) {
	return r.queries.PoolGetAll(ctx)
}

// LoadBalancerRepository provides read operations for load balancers.
type LoadBalancerRepository struct{ queries *Queries }

func NewLoadBalancerRepository(db DBTX) *LoadBalancerRepository {
	return &LoadBalancerRepository{queries: New(db)}
}
func (r *LoadBalancerRepository) GetAllWithVIP(ctx context.Context) ([]LoadBalancerGetAllWithVIPRow, error) {
	return r.queries.LoadBalancerGetAllWithVIP(ctx)
}

// AmphoraRepository provides read operations for amphorae.
type AmphoraRepository struct{ queries *Queries }

func NewAmphoraRepository(db DBTX) *AmphoraRepository { return &AmphoraRepository{queries: New(db)} }
func (r *AmphoraRepository) GetAll(ctx context.Context) ([]AmphoraGetAllRow, error) {
	return r.queries.AmphoraGetAll(ctx)
}

// QuotaRepository provides read operations for project quotas.
type QuotaRepository struct{ queries *Queries }

func NewQuotaRepository(db DBTX) *QuotaRepository { return &QuotaRepository{queries: New(db)} }
func (r *QuotaRepository) GetAll(ctx context.Context) ([]QuotaGetAllRow, error) {
	return r.queries.QuotaGetAll(ctx)
}
func (r *QuotaRepository) LoadBalancerCountByProject(ctx context.Context) ([]LoadBalancerCountByProjectRow, error) {
	return r.queries.LoadBalancerCountByProject(ctx)
}
func (r *QuotaRepository) PoolCountByProject(ctx context.Context) ([]PoolCountByProjectRow, error) {
	return r.queries.PoolCountByProject(ctx)
}
func (r *QuotaRepository) ListenerCountByProject(ctx context.Context) ([]ListenerCountByProjectRow, error) {
	return r.queries.ListenerCountByProject(ctx)
}
