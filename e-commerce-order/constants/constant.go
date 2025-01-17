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
