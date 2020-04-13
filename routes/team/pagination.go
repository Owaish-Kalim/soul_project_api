package  team

// import (
//     "http/net"
//     "encoding/json"
//     "strconv"

//     "github.com/gemcook/pagination-go"
// )


// type PageFetcher interface {
// 	Count(cond interface{}) (int, error)
// 	FetchPage(cond interface{}, input *PageFetchInput, result *PageFetchResult) error
// }

// type PageFetchInput struct {
// 	Limit  int
// 	Offset int
// 	Orders []*Order
// }



// type fruit struct {
// 	Name  string
// 	Price int
// }

// // fetcher

// type fruitsRepository struct {
// 	priceLowerLimit  int
// 	priceHigherLimit int
// }

// type fruitFetcher struct {
// 	repo *fruitsRepository
// }

// func newFruitFetcher() *fruitFetcher {
// 	return &fruitFetcher{
// 		repo: &fruitsRepository{},
// 	}
// }

// // cond

// type fruitCondition struct {
// 	PriceLowerLimit  *int
// 	PriceHigherLimit *int
// }


// func newFruitCondition(low, high int) *fruitCondition {
// 	return &fruitCondition{
// 		PriceLowerLimit:  &low,
// 		PriceHigherLimit: &high,
// 	}
// }

// func parseFruitCondition(queryStr string) *fruitCondition {
// 	u, err := url.Parse(queryStr)
// 	if err != nil {
// 		fmt.Println(err)
// 		low := -1 << 31
// 		high := 1<<31 - 1
// 		return newFruitCondition(low, high)
// 	}
// 	query := u.Query()

// 	if s := query.Get("price_range"); s != "" {
// 		prices := strings.Split(s, ",")
// 		low, err := strconv.Atoi(prices[0])
// 		if err != nil {
// 			panic(err)
// 		}
// 		high, err := strconv.Atoi(prices[1])
// 		if err != nil {
// 			panic(err)
// 		}
// 		return newFruitCondition(low, high)
// 	}

// 	low := -1 << 31
// 	high := 1<<31 - 1
// 	return newFruitCondition(low, high)
// }


// // CUSTOMIZE: GET DATA FROM DATABASE
// func (fr *fruitsRepository) GetFruits(orders []*pagination.Order) []fruit {
// 	result := make([]fruit, 0)
// 	for _, f := range dummyFruits {
// 		if fr.priceHigherLimit >= f.Price && f.Price >= fr.priceLowerLimit {
// 			result = append(result, f)
// 		}
// 	}

// 	for _, o := range orders {
// 		if o.ColumnName != "price" {
// 			continue
// 		}
// 		sort.SliceStable(result, func(i, j int) bool {
// 			if o.Direction == pagination.DirectionAsc {
// 				return result[i].Price < result[j].Price
// 			}

// 			return result[i].Price > result[j].Price
// 		})
// 	}

// 	return result
// }



// /// SETTING INTERFACES OF FETCHER

// // APPLY CONDITION
// func (ff *fruitFetcher) applyCondition(cond *fruitCondition) {
// 	if cond.PriceHigherLimit != nil {
// 		ff.repo.priceHigherLimit = *cond.PriceHigherLimit
// 	}
// 	if cond.PriceLowerLimit != nil {
// 		ff.repo.priceLowerLimit = *cond.PriceLowerLimit
// 	}
// }

// // COUNT
// func (ff *fruitFetcher) Count(cond interface{}) (int, error) {
// 	if cond != nil {
// 		ff.applyCondition(cond.(*fruitCondition))
// 	}
// 	orders := make([]*pagination.Order, 0, 0)
// 	fruits := ff.repo.GetFruits(orders)
// 	return len(fruits), nil
// }

// // FETCHPAGE
// func (ff *fruitFetcher) FetchPage(cond interface{}, input *pagination.PageFetchInput, result *pagination.PageFetchResult) error {
// 	if cond != nil {
// 		ff.applyCondition(cond.(*fruitCondition))
// 	}
// 	fruits := ff.repo.GetFruits(input.Orders)
// 	var toIndex int
// 	toIndex = input.Offset + input.Limit
// 	if toIndex > len(fruits) {
// 		toIndex = len(fruits)
// 	}
// 	for _, fruit := range fruits[input.Offset:toIndex] {
// 		*result = append(*result, fruit)
// 	}
// 	return nil
// }





// func handler(w http.ResponseWriter, r *http.Request) {
// 	// RequestURI: https://example.com/fruits?limit=10&page=1&price_range=100,300&sort=+price
// 	p := pagination.ParseQuery(r.URL.RequestURI())
// 	// cond := parseFruitCondition(r.URL.RequestURI())
// 	fetcher := newFruitFetcher()

// 	totalCount, totalPages, res, err := pagination.Fetch(fetcher, &pagination.Setting{
// 		Limit:      p.Limit,
// 		Page:       p.Page,
// 		// Cond:       cond,
// 		Orders:     p.Sort,
// 	})

// 	if err != nil {
// 		w.Header().Set("Content-Type", "text/html; charset=utf8")
// 		w.WriteHeader(400)
// 		fmt.Fprintf(w, "something wrong: %v", err)
// 		return
// 	}

// 	w.Header().Set("X-Total-Count", strconv.Itoa(totalCount))
// 	w.Header().Set("X-Total-Pages", strconv.Itoa(totalPages))
// 	w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count,X-Total-Pages")
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(200)
// 	resJSON, _ := json.Marshal(res)
// 	w.Write(resJSON)
// }