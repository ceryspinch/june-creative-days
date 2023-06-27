package main

import(
	"time"
	
	pb "git "

	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func init() {
	DatabaseConnection()
 }

 var DB *gorm.DB
 var err error
 
 type Ticket struct{
	ID					string `gorm:"primarykey"`
	Purchaser		 	string
	IsBringingGuest		bool
	HasReceivedTicket	bool
	CreatedAt			time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt 			time.Time `gorm:"autoUpdateTime:false"`
 }

 func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "pass1234"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Movie{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
 }

 var (
	port = flag.Int("port", 50051, "gRPC server port")
 )
  
 type server struct {
	pb.UnimplementedMovieServiceServer
 }