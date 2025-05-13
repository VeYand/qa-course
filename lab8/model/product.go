package model

type Product struct {
	ID          string `json:"id"`
	CategoryID  string `json:"category_id"` // число от 1 до 15
	Title       string `json:"title"`
	Alias       string `json:"alias"` // формируется из поля title через транслит на латиницу. Но если такой алиас существует, то добавляется префикс -0
	Content     string `json:"content"`
	Price       string `json:"price"`
	OldPrice    string `json:"old_price"`
	Status      string `json:"status"` // (0, 1)
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Hit         string `json:"hit"` // (0, 1)
}
