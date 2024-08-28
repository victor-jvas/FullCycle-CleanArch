package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {

	var out []OrderOutputDTO
	orders, err := l.OrderRepository.GetAll()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}

		out = append(out, dto)
	}

	l.OrdersListed.SetPayload(out)
	l.EventDispatcher.Dispatch(l.OrdersListed)

	return out, nil
}
