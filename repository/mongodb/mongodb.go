package mongodb

import (
	"os"

	"online-food/domain"

	"github.com/Kamva/mgm/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoRepository struct {
	OrderRepo   domain.OrderRepository
	ProductRepo domain.ProductRepository
}

func New(l *zap.Logger) *MongoRepository {
	if err := godotenv.Load(); err != nil {
		l.Warn("No .env file found, using environment variables")
	}

	connectionString := os.Getenv("DATABASE")
	dbName := os.Getenv("DATABASE_NAME")

	if connectionString == "" || dbName == "" {
		l.Fatal("Required database environment variables not set",
			zap.Bool("connection_string_empty", connectionString == ""),
			zap.Bool("db_name_empty", dbName == ""))
	}

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(connectionString))
	if err != nil {
		l.Fatal("cannot set mgm config", zap.Error(err))
	}

	l.Info("Db connected")

	return &MongoRepository{
		OrderRepo:   NewOrderRepository(l),
		ProductRepo: NewProductRepository(l),
	}
}
