package main

import (
	"fmt"
	"strings"

	"github.com/ajbosco/reads/utils"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	searchCommand = cli.Command{
		Name:   "search",
		Usage:  "search for a book by title, author, or id",
		Action: searchBooks,
	}
)

const baseBookURI = "https://www.goodreads.com/book/show"

func searchBooks(c *cli.Context) error {
	if len(c.Args()) <= 0 {
		fmt.Println("You didn't search for anything!")
		return nil
	}
	args := append([]string{c.Args().First()}, c.Args().Tail()...)
	searchText := strings.Join(args, " ")
	books, err := client.SearchBooks(searchText)
	if err != nil {
		logrus.Fatal(err)
	}

	table := utils.NewTable("ID", "Title", "Author", "Average Rating", "Ratings Count", "URI")

	for _, b := range books {
		uri := fmt.Sprintf("%s/%s", baseBookURI, b.ID)
		table.Append([]string{b.ID, b.Title, b.Author, b.AvgRating, b.RatingsCount, uri})
	}
	fmt.Fprintf(c.App.Writer, "Results for %q\n", searchText)
	table.Render()
	return nil
}
