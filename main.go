package main

import "fmt"

const (
	// Library Information
	LibraryName  = "British Library"
	Location     = "United Kingdom"
	WorkingHours = "09:00 AM - 06:00 PM"

	// Book Status
	AvailableStatus = "Available"
	BorrowedStatus  = "Borrowed"

	// Book Categories
	ProgrammingCategory = "Programming"
	DatabaseCategory    = "Database"
	NetworkingCategory  = "Networking"
	WebCategory         = "Web Development"
	AICategory          = "Artificial Intelligence"

	// Library Rules
	MaxBooks      = 100
	MaxBorrowDays = 14
)

type Book struct {
	BookID          string
	Title           string
	Author          string
	Category        string
	PublicationYear int
	Pages           int
	Status          string
}

var (
	library []Book

	totalBooks     int
	availableBooks int
	borrowedBooks  int
)

func libraryInfo() (
	string,
	string,
	string,
) {
	libraryName := LibraryName
	location := Location
	workingHours := WorkingHours

	return libraryName, location, workingHours
}

func createBooks() {

	// Create Book 1
	book1 := Book{
		BookID:          "BK001",
		Title:           "Learning Go",
		Author:          "Jon Bodner",
		Category:        ProgrammingCategory,
		PublicationYear: 2021,
		Pages:           350,
		Status:          AvailableStatus,
	}

	// Create Book 2
	book2 := Book{
		BookID:          "BK002",
		Title:           "Clean Code",
		Author:          "Robert C. Martin",
		Category:        ProgrammingCategory,
		PublicationYear: 2008,
		Pages:           464,
		Status:          BorrowedStatus,
	}

	// Create Book 3
	book3 := Book{
		BookID:          "BK003",
		Title:           "Computer Networks",
		Author:          "Andrew S. Tanenbaum",
		Category:        NetworkingCategory,
		PublicationYear: 2011,
		Pages:           960,
		Status:          AvailableStatus,
	}

	// Create Book 4
	book4 := Book{
		BookID:          "BK004",
		Title:           "Learning SQL",
		Author:          "Alan Beaulieu",
		Category:        DatabaseCategory,
		PublicationYear: 2020,
		Pages:           394,
		Status:          AvailableStatus,
	}

	// Create Book 5
	book5 := Book{
		BookID:          "BK005",
		Title:           "Artificial Intelligence",
		Author:          "Stuart Russell",
		Category:        AICategory,
		PublicationYear: 2021,
		Pages:           1136,
		Status:          BorrowedStatus,
	}

	// Add all books into library
	library = append(library,
		book1,
		book2,
		book3,
		book4,
		book5,
	)

	// Update statistics
	totalBooks = len(library)

	for _, book := range library {

		if book.Status == AvailableStatus {
			availableBooks++
		} else {
			borrowedBooks++
		}
	}
}

func addBook() {

	var newBook Book

	fmt.Println("===================================")
	fmt.Println("          ADD NEW BOOK")
	fmt.Println("===================================")

	fmt.Print("Enter book ID: ")
	fmt.Scanln(&newBook.BookID)

	fmt.Print("Enter Book Title: ")
	fmt.Scanln(&newBook.Title)

	fmt.Print("Enter Author Name: ")
	fmt.Scanln(&newBook.Author)

	fmt.Print("Enter Category: ")
	fmt.Scanln(&newBook.Category)

	fmt.Print("Enter Publication Year: ")
	fmt.Scanln(&newBook.PublicationYear)

	fmt.Print("Enter Total Pages: ")
	fmt.Scanln(&newBook.Pages)

	newBook.Status = AvailableStatus

	library = append(library, newBook)

	totalBooks++
	availableBooks++

	fmt.Println()
	fmt.Println("Book added Successfully.")
}

func removeBook() {

	var bookID string

	fmt.Println("===================================")
	fmt.Println("         REMOVE BOOK")
	fmt.Println("===================================")

	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&bookID)

	for i, book := range library {
		if book.BookID == bookID {
			if book.Status == AvailableStatus {
				availableBooks--
			} else {
				borrowedBooks--
			}

			library = append(library[:i], library[i+1:]...)

			totalBooks--

			fmt.Println()
			fmt.Println("Book Removed Successfully.")
			return
		}
	}

	fmt.Println()
	fmt.Println("Book Not Found")
}

func borrowBook() {

	var bookID string

	fmt.Println("===================================")
	fmt.Println("          BORROW BOOK")
	fmt.Println("===================================")

	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&bookID)

	for i := range library {
		if library[i].BookID == bookID {
			if library[i].Status == BorrowedStatus {
				fmt.Println()
				fmt.Println("This Book is Already Borrowed.")
				return
			}

			library[i].Status = BorrowedStatus

			availableBooks--
			borrowedBooks++

			fmt.Println()
			fmt.Println("Book Borrowed Successfully.")
			return
		}
	}

	fmt.Println()
	fmt.Println("Book Not Found.")
}

func returnBook() {

	var bookID string

	fmt.Println("===================================")
	fmt.Println("          RETURN BOOK")
	fmt.Println("===================================")

	fmt.Println("Enter Book ID: ")
	fmt.Scanln(&bookID)

	for i := range library {

		if library[i].BookID == bookID {
			if library[i].Status == AvailableStatus {

				fmt.Println()
				fmt.Println("This Book is Already Available.")
				return
			}

			library[i].Status = AvailableStatus

			availableBooks++
			borrowedBooks--

			fmt.Println()
			fmt.Println("Book Returned Successfully.")
			return
		}
	}

	fmt.Println()
	fmt.Println("Book Not Found")
}

func searchBook() {

	var bookID string

	fmt.Println("===================================")
	fmt.Println("          SEARCH BOOK")
	fmt.Println("===================================")

	fmt.Println("Enter Book ID: ")
	fmt.Scanln(&bookID)

	for _, book := range library {
		if book.BookID == bookID {

			fmt.Println()
			fmt.Println("========== BOOK INFORMATION ==========")
			fmt.Println("Book ID          :", book.BookID)
			fmt.Println("Title            :", book.Title)
			fmt.Println("Author           :", book.Author)
			fmt.Println("Category         :", book.Category)
			fmt.Println("Publication Year :", book.PublicationYear)
			fmt.Println("Pages            :", book.Pages)
			fmt.Println("Status           :", book.Status)
			fmt.Println("======================================")

			return
		}
	}

	fmt.Println()
	fmt.Println("Book Not Found.")
}

func printBooks() {

	fmt.Println("==============================================================")
	fmt.Println("                    LIBRARY BOOK LIST")
	fmt.Println("==============================================================")

	if len(library) == 0 {
		fmt.Println("No books Available in the Library")
		return
	}

	for i, book := range library {
		fmt.Printf("Book %d\n", i+1)
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("Book ID          :", book.BookID)
		fmt.Println("Title            :", book.Title)
		fmt.Println("Author           :", book.Author)
		fmt.Println("Category         :", book.Category)
		fmt.Println("Publication Year :", book.PublicationYear)
		fmt.Println("Pages            :", book.Pages)
		fmt.Println("Status           :", book.Status)
		fmt.Println()
	}
	fmt.Println("==============================================================")
	fmt.Printf("Total Books: %d\n", len(library))
	fmt.Println("==============================================================")
}

func printReport(

	libraryName string,
	location string,
	workingHours string,

	totalBooks int,
	availableBooks int,
	borrowedBooks int,
) {
	fmt.Println("======================================================")
	fmt.Println("          LIBRARY MANAGEMENT SYSTEM REPORT")
	fmt.Println("======================================================")

	fmt.Println("Library Name      :", libraryName)
	fmt.Println("Location          :", location)
	fmt.Println("Working Hours     :", workingHours)

	fmt.Println("------------------------------------------------------")

	fmt.Println("Total Books       :", totalBooks)
	fmt.Println("Available Books   :", availableBooks)
	fmt.Println("Borrowed Books    :", borrowedBooks)

	fmt.Println("------------------------------------------------------")

	fmt.Println("Library Status    : Active")
	fmt.Println("Maximum Books     :", MaxBooks)
	fmt.Println("Borrow Limit      :", MaxBorrowDays, "Days")

	fmt.Println("======================================================")
	fmt.Println("        THANK YOU FOR USING OUR LIBRARY")
	fmt.Println("======================================================")
}

func main() {

	defer fmt.Println("\nProgram Finished Successfully.")

	// Get library information
	libraryName, location, workingHours := libraryInfo()

	// Create sample books
	createBooks()

	// Print all books
	printBooks()

	fmt.Println()

	// Search a book
	searchBook()

	fmt.Println()

	// Add a new book
	addBook()

	fmt.Println()

	// Borrow a book
	borrowBook()

	fmt.Println()

	// Return a borrowed book
	returnBook()

	fmt.Println()

	// Remove a book
	removeBook()

	fmt.Println()

	// Print all books again
	printBooks()

	fmt.Println()

	// Print final report
	printReport(
		libraryName,
		location,
		workingHours,
		totalBooks,
		availableBooks,
		borrowedBooks,
	)
}
