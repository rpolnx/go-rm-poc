package adapter

type couponsAdapter struct {
	repo repository.CouponsRepository
}

func (adapter *couponsAdapter) GetCoupons(limit int, storeId string) (cp.AggregatedCouponsViewList, error) {
	return adapter.repo.GetListCoupons(limit, storeId)
}
func InitializeCouponsAdapter(repo repository.CouponsRepository) ports.CouponsPort {
	return &couponsAdapter{repo: repo}
}
