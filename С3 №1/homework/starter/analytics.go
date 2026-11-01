package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func AnalyzeSales(sales []Sale) Report {
	report := Report{
		RevenueByCategory: make(map[string]int64),
	}
	topProducts := make(map[string]int)
	for _, unit := range sales {
		if unit.Cancelled {
			continue
		} 
		if unit.Product == "" || unit.Category == "" || unit.Quantity < 1 || unit.UnitPriceKopecks < 1 {
			report.InvalidSales++
		} else {
			productRevenue := unit.UnitPriceKopecks * int64(unit.Quantity)
			topProducts[unit.Product] += unit.Quantity
			report.RevenueKopecks += productRevenue
			report.UnitsSold += unit.Quantity
			report.RevenueByCategory[unit.Category] += productRevenue
		}
	}

	if len(topProducts) > 0 {
		topName := ""
		topCount := 0
		for name, count := range topProducts {
			if count > topCount || (count == topCount && name < topName) {
				topName = name
				topCount = count
			}
		}
		report.TopProduct = topName
	}

	return report
}
