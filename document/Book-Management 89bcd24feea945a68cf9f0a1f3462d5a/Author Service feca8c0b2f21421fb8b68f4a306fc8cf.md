# Author Service

## Category-svc

Service that serves requests related to book lists and searches, adding books, editing books, deleting books, borrowing, returning books.

### General Response:

- Error handling
    - protobuf status (INTERNAL)
    - protobuf message (cant process request)
- Success handling
    - protobuf status (OK)
    - protobuf message (success)
    - body (data)

### Protobuf Contract:

- **TOKEN**
    
    Url: localhost:50053/GetToken
    
    Request:
    
    ```json
    {
        "username": "wandi",
        "password": "cobacoba"
    }
    ```
    
- **VERIFY**
    
    Url: localhost:50053/ValidateToken
    
    Request:
    
    ```
    {
        "token": "qBPsKzEpSMDBV8_vzqQGKnzhTV1P8rRpmMtHWUlFGrU"
    }
    ```
    
- **CREATE**
    
    Url: localhost:50053/CreateUser
    
    Request:
    
    ```
    {
        "first_name": "irwandi",
        "last_name": "paputungan",
        "username": "wandi",
        "email": "irwandi@mail.com",
        "password": "cobacoba"
    }
    ```