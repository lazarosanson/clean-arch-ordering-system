package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type OrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: orderRepository}
}

func (uc *ListOrdersUseCase) Execute() ([]OrderDTO, error) {
	orders, err := uc.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersOutput []OrderDTO
	for _, order := range orders {
		orderOutput := OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}

		ordersOutput = append(ordersOutput, orderOutput)
	}

	return ordersOutput, nil
}
