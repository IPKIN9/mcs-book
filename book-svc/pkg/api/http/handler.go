package http

import (
	"book-svc/pkg/protos"
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func SetupRoutes(app *fiber.App) {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := protos.NewBookServiceClient(conn)

	app.Get("/books", func(c *fiber.Ctx) error {
		req := new(protos.GetAllBookRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		res, err := client.GetAllBooks(context.Background(), req)

		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				// Map gRPC status code to HTTP status code
				switch st.Code() {
				case codes.NotFound:
					return c.Status(fiber.StatusNotFound).SendString(st.Message())
				case codes.InvalidArgument:
					return c.Status(fiber.StatusBadRequest).SendString(st.Message())
				case codes.Internal:
					return c.Status(fiber.StatusInternalServerError).SendString(st.Message())
				// Add more cases as needed for other gRPC error codes
				default:
					return c.Status(fiber.StatusInternalServerError).SendString("Unexpected error")
				}
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Unknown error")
		}

		protoBytes, err := proto.Marshal(res)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Set(fiber.HeaderContentType, "application/x-protobuf")
		return c.Send(protoBytes)
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		book := new(protos.CreateBookRequest)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		res, err := client.CreateBook(context.Background(), book)

		protoBytes, err := proto.Marshal(res)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Set(fiber.HeaderContentType, "application/x-protobuf")
		return c.Send(protoBytes)
	})

	app.Put("/books/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
		book := new(protos.UpdateBookRequest)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		book.BookId = id
		res, err := client.UpdateBook(context.Background(), book)

		protoBytes, err := proto.Marshal(res)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Set(fiber.HeaderContentType, "application/x-protobuf")
		return c.Send(protoBytes)
	})

	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
		req := &protos.DeleteBookRequest{BookId: id}
		_, err := client.DeleteBook(context.Background(), req)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendStatus(200)
	})

	app.Get("/borrow", func(c *fiber.Ctx) error {
		req := new(protos.BorrowingBookRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		res, err := client.BorrowingBook(context.Background(), req)

		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				// Map gRPC status code to HTTP status code
				switch st.Code() {
				case codes.NotFound:
					return c.Status(fiber.StatusNotFound).SendString(st.Message())
				case codes.InvalidArgument:
					return c.Status(fiber.StatusBadRequest).SendString(st.Message())
				case codes.Internal:
					return c.Status(fiber.StatusInternalServerError).SendString(st.Message())
				// Add more cases as needed for other gRPC error codes
				default:
					return c.Status(fiber.StatusInternalServerError).SendString("Unexpected error")
				}
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Unknown error")
		}

		protoBytes, err := proto.Marshal(res)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Set(fiber.HeaderContentType, "application/x-protobuf")
		return c.Send(protoBytes)
	})
}
