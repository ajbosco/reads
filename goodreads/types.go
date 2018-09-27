package goodreads

// Response defines the data struct for an API response object.
type Response struct {
	Search  Search  `xml:"search,omitempty"`
	Shelves Shelves `xml:"shelves,omitempty"`
	User    User    `xml:"user,omitempty"`
	Reviews Reviews `xml:"reviews,omitempty"`
}

// Search defines the struct for the search object
type Search struct {
	Works []Work `xml:"results>work,omitempty" json:"work,omitempty"`
}

// Work defines the struct for the work object
type Work struct {
	ID           string `xml:"best_book>id,omitempty"`
	Title        string `xml:"best_book>title,omitempty"`
	Author       string `xml:"best_book>author>name,omitempty"`
	RatingsCount string `xml:"ratings_count,omitempty"`
	AvgRating    string `xml:"average_rating,omitempty"`
}

// Shelves defines the struct for the shelves object
type Shelves struct {
	UserShelves []UserShelf `xml:"user_shelf,omitempty"`
}

// UserShelf defines the struct for the user shelf object
type UserShelf struct {
	ID        string `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	BookCount string `xml:"book_count,omitempty"`
}

// User defines the struct for the user object
type User struct {
	ID string `xml:"id,attr"`
}

// Reviews defines the struct for the reviews object
type Reviews struct {
	Books []Book `xml:"review>book,omitempty"`
}

// Book defines the struct for the book object
type Book struct {
	ID        string `xml:"id,omitempty"`
	Title     string `xml:"title,omitempty"`
	AvgRating string `xml:"average_rating,omitempty"`
	Link      string `xml:"link,omitempty"`
}
