package main

import (
	"net/http"
	"os"

	"aetherium.com/shared/go/logger"
	"aetherium.com/user-service/app/api"
	"aetherium.com/user-service/app/features/signup"

	"aetherium.com/user-service/app/externals/database"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const defaultPort = "8081"

func main() {
	// 0. Initialize our beautiful logger
	log.Logger = logger.New()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msg("Logger initialized")

	// Load environment variables from a .env file.
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No .env file found, relying on environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal().Msg("FATAL: DATABASE_URL is not set in environment")
	}

	// Establish a connection to the database.
	pool, err := database.NewPostgresConnection(dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer pool.Close()
	log.Info().Msg("Successfully connected to the database")

	// --- Dependency Injection: Construct application layers ---
	// 1. Create the repository instance.
	signupRepo := signup.NewSignUpRepository(pool)
	// 2. Create the service instance, injecting the repository.
	signupService := signup.NewSignUpService(signupRepo)
	// 3. Create the GraphQL resolver, injecting the service.
	resolver := &api.Resolver{SignUpService: signupService}
	// --- End of Dependency Injection ---

	// Set up the HTTP router.
	router := chi.NewRouter()

	// Create the GraphQL server.
	srv := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: resolver}))

	// Register the GraphQL endpoints.
	router.Handle("/", playground.Handler("Aetherium User Service", "/graphql"))
	router.Handle("/graphql", srv)

	log.Info().Str("port", port).Msgf("GraphQL playground running at http://localhost:%s/", port)

	// Start the server.
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}