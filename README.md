# **GoBookstore**

GoBookstore is a small project to experiment with **Go** by building a simple RESTful API for managing a collection of books.

## **Getting Started**

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/gobookstore.git
   cd gobookstore
   ```

2. Run the application with Docker:
   ```bash
   docker-compose up --build
   ```

3. Access the API at `http://localhost:8080`.

## **Testing**

### Local Testing
1. Make sure you have Go installed on your machine
2. Run the tests:
   ```bash
   go test ./...
   ```
3. For verbose output:
   ```bash
   go test -v ./...
   ```
4. For test coverage:
   ```bash
   go test -cover ./...
   ```

## **Features**
- Add, view, and delete books.
- Uses MongoDB as the database.

## **Endpoints**
- `GET /books` - Retrieve all books.
- `GET /books/{id}` - Retrieve a book by ID.
- `POST /books` - Add a new book.
- `DELETE /books/{id}` - Delete a book by ID.
- `POST /books/random` - Create a random book.

## **License**
Feel free to use or modify this project however you like.