#  Library Management System

A simple **Library Management System** built with **Go (Golang)**.

This project was created to practice the core concepts of Go programming such as **Structs, Slices, Functions, Loops, CRUD Operations, Global Variables, Constants, and User Input**.

---

#  Features

-  Display library information
-  Add a new book
-  Remove a book
-  Search a book by ID
-  Display all books
-  Borrow a book
-  Return a borrowed book
-  Print library report
-  Library statistics

---

# Technologies

- Go (Golang)
- VS Code
- Standard Go Library

---

#  Project Structure

```
library-management-system/
│
├── main.go
├── README.md
└── go.mod
```

---

#  Book Structure

Each book contains the following information:

- Book ID
- Title
- Author
- Category
- Publication Year
- Pages
- Status

---

#  Sample Books

The application starts with five sample books.

| Book ID | Title | Status |
|---------|---------------------------|------------|
| BK001 | Learning Go | Available |
| BK002 | Clean Code | Borrowed |
| BK003 | Computer Networks | Available |
| BK004 | Learning SQL | Available |
| BK005 | Artificial Intelligence | Borrowed |

---

# ⚙️ Available Functions

### Library Information

- `libraryInfo()`

Returns basic library information.

---

### Book Management

- `createBooks()`
- `addBook()`
- `removeBook()`
- `borrowBook()`
- `returnBook()`
- `searchBook()`
- `printBooks()`

---

### Report

- `printReport()`

Displays the final library report.

---

#  Example Output

```
======================================================
          LIBRARY MANAGEMENT SYSTEM REPORT
======================================================
Library Name      : British Library
Location          : United Kingdom
Working Hours     : 09:00 AM - 06:00 PM

Total Books       : 5
Available Books   : 3
Borrowed Books    : 2

Library Status    : Active
Maximum Books     : 100
Borrow Limit      : 14 Days
======================================================
```

---

#  Concepts Practiced

This project helped me practice:

- Variables
- Constants
- Structs
- Functions
- Multiple Return Values
- Global Variables
- Slices
- append()
- len()
- for-range
- if / else
- CRUD Operations
- User Input
- fmt package

---

#  Future Improvements

- Interactive menu system
- Book ID validation
- Duplicate book checking
- Better input handling
- Data persistence (JSON/File)
- Database integration (PostgreSQL/MySQL)
- Search by title
- Search by author
- Update book information
- Colored terminal output

---

# Learning Goal

This project is part of my Go learning journey.

The main goal of this project was to strengthen my understanding of Go fundamentals by building a real-world console application.

---

# Author

**Ulug'bek Zokirov**

Backend Developer (Go Learner)
