# Book Service

## About

Service that serves requests related to book lists and searches, adding books, editing books, deleting books, borrowing, returning books.

### General Response:

- Error handling
    - protobuf status: INTERNAL
    - protobuf message: cant process request
- Success handling
    - protobuf status: OK
    - protobuf message: success
    - body: data

### Protobuf Contract:

- GET
    
    Url: localhost:50051/GetBook
    
    Request:
    
    ```
    {
        "book_id": "1"
    }
    ```
    
- GET ALL
    
    Url: localhost:50051/UpdateBook
    
    Request:
    
    ```
    {
        "search": {
            "value": ""
        },
        "limit": 15,
        "page": 1
    }
    ```
    
- CREATE
    
    Url: localhost:50051/CreateBook
    
    Request:
    
    ```jsx
    {
        "author_id": "1",
        "category_id": "1",
        "description": "no",
        "isbn": "ABC",
        "published_date": {"nanos": 0, "seconds": "0"},
        "stock_id": "1",
        "title": "testing 3"
    }
    ```
    
- UPDATE
    
    Url: localhost:50051/GetAllBooks
    
    Request:
    
    ```jsx
    {
        "book_id": 3,
        "book": {
             "author_id": "1",
            "category_id": "1",
            "description": "update",
            "isbn": "ABB",
            "published_date": {"nanos": 0, "seconds": "0"},
            "stock_id": "1",
            "title": "Coba"
        }
    }
    ```
    
- DELETE
    
    Url: localhost:50051/DeleteBook
    
    Request:
    
    ```jsx
    {
        "book_id": "1"
    }
    ```
    
- BORROWING
    
    Url: localhost:50051/BorrowingBook
    
    Request:
    
    ```jsx
    {
        "book_id": [
            3
        ],
        "user_id": {
            "value": "1"
        },
        "is_borrow": false
    }
    ```
    
- CHANGE STOCK
    
    Url: localhost:50051/ChangeStock
    
    Request:
    
    ```
    {
      "book_id": "1",
      "total_qty": "100",
      "avaible_qty": "15"
    }
    ```
    
- GET STOCK
    
    Url: localhost:50051/GetStock
    
    Request:
    
    ```
    {
        "book_id": {
            "value": "1"
        }
    }
    ```