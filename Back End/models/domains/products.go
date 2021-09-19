package domains

type Product struct {
	ID            int
	NAME          string
	Price         int
	Size          []string
	Quantity      int
	Details       string
	CategoryID    int
	SubCategoryID int
}
