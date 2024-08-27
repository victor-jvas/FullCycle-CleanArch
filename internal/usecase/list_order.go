package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
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

	l.OrderCreated.SetPayload(out)
	l.EventDispatcher.Dispatch(l.OrderCreated)

	return out, nil
}
