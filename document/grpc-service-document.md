# Book-Management

## BookService dan Book Database

**Tables:**

- **Books**
    - `book_id` (primary key)
    - `title`
    - `isbn`
    - `author_id` (foreign key to Author Database)
    - `category_id` (foreign key to Category Database)
    - `published_date`
    - `description`
    - `stock_id` (foreign key to Stock table in the same database)
    - `created_at`
    - `updated_at`
- **Stock**
    - `stock_id` (primary key)
    - `book_id` (foreign key to Books table)
    - `total_quantity`
    - `available_quantity`
    - `created_at`
    - `updated_at`

## **AuthorService dan Author Database**

**Tables:**

- **Authors**
    - `author_id` (primary key)
    - `first_name`
    - `last_name`
    - `biography`
    - `date_of_birth`
    - `created_at`
    - `updated_at`

## CategoryService dan Category Database

**Tables:**

- **Categories**
    - `category_id` (primary key)
    - `category_name`
    - `description`
    - `created_at`
    - `updated_at`

## UserService/AuthService dan User Database

**Tables:**

- **Users**
    - `user_id` (primary key)
    - `username`
    - `password`
    - `email`
    - `first_name`
    - `last_name`
    - `created_at`
    - `updated_at`
- **Borrowing**
    - `borrowing_id` (primary key)
    - `user_id` (foreign key to Users table)
    - `book_id` (foreign key to Book Database)
    - `borrow_date`
    - `due_date`
    - `return_date`
    - `created_at`
    - `updated_at`
- **Tokens**
    - `token_id` (primary key)
    - `user_id` (foreign key to Users table)
    - `token`
    - `expired`
    - `created_at`
    - `updated_at`
- **UserRecommendations**
    - `recommendation_id` (primary key)
    - `user_id` (foreign key to Users table)
    - `book_id` (foreign key to Book Database)
    - `created_at`
    - `updated_at`

[Book Service](Book-Management%2089bcd24feea945a68cf9f0a1f3462d5a/Book%20Service%20c1afff5c0f874ef7916e1a690ece7328.md)

[User Service](Book-Management%2089bcd24feea945a68cf9f0a1f3462d5a/User%20Service%204115e18d556f4c19b92f3a1284a6a6a5.md)

[Author Service](Book-Management%2089bcd24feea945a68cf9f0a1f3462d5a/Author%20Service%20feca8c0b2f21421fb8b68f4a306fc8cf.md)