package constants

const (
	RedisKeyProducts      = "ecommerce:product"
	RedisKeyProductDetail = "ecommerce:product:%d"
)

const (
	OrderStatusSuccess = "SUCCESS"
	OrderStatusPending = "PENDING"
	OrderStatusFailed  = "FAILED"
	OrderStatusRefund  = "REFUND"
)

var MappingOrderStatus = map[string]bool{
	OrderStatusSuccess: true,
	OrderStatusPending: true,
	OrderStatusFailed:  true,
	OrderStatusRefund:  true,
}

var MappingFlowOrderStatus = map[string][]string{
	OrderStatusPending: {OrderStatusSuccess, OrderStatusFailed},
	OrderStatusSuccess: {OrderStatusRefund},
	OrderStatusFailed:  {OrderStatusRefund},
}
