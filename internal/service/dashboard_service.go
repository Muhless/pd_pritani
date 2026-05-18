package service

import (
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type DashboardData struct {
	TotalRevenue     decimal.Decimal `json:"total_revenue"`
	TotalExpense     decimal.Decimal `json:"total_expense"`
	TotalCustomer    int64           `json:"total_customer"`
	TotalProduct     int64           `json:"total_product"`
	LowStockProducts []model.Product `json:"low_stock_products"`
	RecentSales      interface{}     `json:"recent_sales"`
	Month            int             `json:"month"`
	Year             int             `json:"year"`
}

type DashboardService interface {
	GetDashboard() (*DashboardData, error)
}

type dashboarService struct {
	salesRepo    repository.SalesRepository
	purchaseRepo repository.PurchaseRepository
	productRepo  repository.ProductRepository
	customerRepo repository.CustomerRepository
}

func NewDashboardService(
	salesRepo repository.SalesRepository,
	purchaseRepo repository.PurchaseRepository,
	productRepo repository.ProductRepository,
	customerRepo repository.CustomerRepository,
) DashboardService {
	return &dashboarService{salesRepo, purchaseRepo, productRepo, customerRepo}
}

func (s *dashboarService) GetDashboard() (*DashboardData, error) {
	now := time.Now()
	month := int(now.Month())
	year := now.Year()

	data := &DashboardData{
		Month: month,
		Year:  year,
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error

	setError := func(err error) {
		mu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		mu.Unlock()
	}

	wg.Add(6)

	// total this month
	go func() {
		defer wg.Done()
		total, err := s.salesRepo.GetTotalRevenue(month, year)
		if err != nil {
			setError(err)
			return
		}
		mu.Lock()
		data.TotalRevenue = total
		mu.Unlock()
	}()

	// total expense this month
	go func() {
		defer wg.Done()
		total, err := s.purchaseRepo.GetTotalExpense(month, year)
		if err != nil {
			setError(err)
			return
		}
		mu.Lock()
		data.TotalExpense = total
		mu.Unlock()
	}()

	// get total customer
	go func() {
		defer wg.Done()
		count, err := s.customerRepo.Count()
		if err != nil {
			setError(err)
			return
		}
		mu.Lock()
		data.TotalCustomer = count
		mu.Unlock()
	}()

	// get total product
	go func() {
		defer wg.Done()
		count, err := s.productRepo.Count()
		if err != nil {
			setError(err)
			return
		}
		mu.Lock()
		data.TotalProduct = count
		mu.Unlock()
	}()

	// low stock product
	go func() {
		defer wg.Done()
		products, err := s.productRepo.GetLowStock(5)
		if err != nil {
			setError(err)
			return
		}
		mu.Lock()
		data.LowStockProducts = products
		mu.Unlock()
	}()

	// recent sales
	go func() {
		defer wg.Done()
		sales, err := s.salesRepo.GetRecentSales(5)
		if err != nil {
			setError(err)
			return
		}
		mu.Lock()
		data.RecentSales = sales
		mu.Unlock()
	}()

	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}
	return data, nil
}
