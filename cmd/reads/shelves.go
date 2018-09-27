package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/ajbosco/reads/utils"
)

var (
	shelvesCommand = cli.Command{
		Name:  "shelves",
		Usage: "view shelves and add books to them",
		Subcommands: []cli.Command{
			{
				Name:   "list",
				Usage:  "list your shelves",
				Action: listShelves,
			},
			{
				Name:   "show",
				Usage:  "show books on shelf",
				Action: listShelfBooks,
				Flags: []cli.Flag{
					cli.StringFlag{Name: "shelf, s",
						Usage:       "-s=shelf-name",
						Destination: &shelf},
				},
			},
			{
				Name:   "add",
				Usage:  "add a book to shelf",
				Action: addBookToShelf,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "shelf, s",
						Usage:       "-s=shelf-name",
						Destination: &shelf,
					},
					cli.StringFlag{
						Name:        "book-id, b",
						Usage:       "-b=book-id",
						Destination: &bookID,
					},
				},
			},
		},
	}

	shelf  string
	bookID string
)

func listShelves(c *cli.Context) error {
	user, err := client.GetCurrentUserID()
	if err != nil {
		logrus.Fatal(err)
	}

	shelves, err := client.ListShelves(user.ID)
	if err != nil {
		logrus.Fatal(err)
	}

	table := utils.NewTable("ID", "Shelf Name", "Book Count")

	for _, s := range shelves {
		table.Append([]string{s.ID, s.Name, s.BookCount})
	}
	table.Render()
	return nil
}

func listShelfBooks(c *cli.Context) error {
	if shelf == "" {
		fmt.Println("You didn't provide a shelf!")
		return nil
	}
	user, err := client.GetCurrentUserID()
	if err != nil {
		logrus.Fatal(err)
	}

	books, err := client.ListShelfBooks(shelf, user.ID)
	if err != nil {
		logrus.Fatal(err)
	}

	table := utils.NewTable("ID", "Title", "Average Rating", "URI")

	for _, b := range books {
		table.Append([]string{b.ID, b.Title, b.AvgRating, b.Link})
	}
	table.Render()
	return nil
}

func addBookToShelf(c *cli.Context) error {
	if shelf == "" {
		fmt.Println("You didn't provide a shelf!")
		return nil
	}
	if bookID == "" {
		fmt.Println("You didn't provide a book ID!")
		return nil
	}
	err := client.AddToShelf(shelf, bookID)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Fprintf(c.App.Writer, "Successfully added book id %q to shelf %q!\n", bookID, shelf)
	return nil
}
