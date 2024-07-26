package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/serenite11/market/proto/services/order_service_v1"
	order_model "github.com/serenite11/market/services/order-service/internal/domain/order/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package   github.com/serenite11/market/services/order-service/internal/domain/order/mapper
// goverter:name OrderConverter
// goverter:useZeroValueOnPointerInconsistency
// goverter:enum no
type Converter interface {
	// goverter:ignore state sizeCache unknownFields
	// goverter:map Id | ConvertId
	// goverter:map UserId | ConvertId
	// goverter:map CreatedAt | ConvertTime
	// goverter:map UpdatedAt | ConvertTime
	// goverter:map CompletedAt | ConvertPQTime
	FromEntityOrderToProto(order *order_model.Order) *order_service_v1.Order
	FromEntityOrdersToProto(orders []*order_model.Order) []*order_service_v1.Order
}

func ConvertId(uuid uuid.UUID) string {
	return uuid.String()
}

func ConvertTime(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}

func ConvertPQTime(time pq.NullTime) *timestamppb.Timestamp {
	return timestamppb.New(time.Time)
}
