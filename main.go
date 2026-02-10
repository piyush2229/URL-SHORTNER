package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Url struct {
	ID           int    `json:"id" bson:"id"`
	OriginalUrl  string `json:"original_url" bson:"original_url"`
	ShortUrl     string `json:"short_url" bson:"short_url"`
	CreationDate string `json:"creation_date" bson:"creation_date"`
}

var (
	urlDB      = make(map[string]Url)
	urlsCol    *mongo.Collection
	baseURL    = "http://localhost:3000"
	mongoReady = false
)

// --------- UTILS ----------

func generateShortUrl(originalUrl string) string {
	hash := md5.Sum([]byte(originalUrl))
	return hex.EncodeToString(hash[:])[:8]
}

func createUrl(originalUrl string) Url {
	shortUrl := generateShortUrl(originalUrl)

	if existing, ok := urlDB[shortUrl]; ok {
		return existing
	}

	url := Url{
		ID:           len(urlDB) + 1,
		OriginalUrl:  originalUrl,
		ShortUrl:     shortUrl,
		CreationDate: time.Now().Format("2006-01-02"),
	}
	return url
}

func loadEnvOrDefault(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func connectMongo(ctx context.Context) (*mongo.Client, error) {
	uri := loadEnvOrDefault("MONGO_URI", "mongodb://localhost:27017")
	dbName := loadEnvOrDefault("MONGO_DB", "url_shortener")
	colName := loadEnvOrDefault("MONGO_COLLECTION", "urls")
	baseURL = loadEnvOrDefault("BASE_URL", baseURL)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	urlsCol = client.Database(dbName).Collection(colName)
	mongoReady = true
	return client, nil
}

// --------- MIDDLEWARE ----------

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

// --------- HANDLERS ----------

func handleShorten(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.URL == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	url := createUrl(body.URL)
	urlDB[url.ShortUrl] = url

	if mongoReady && urlsCol != nil {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		var existing Url
		err := urlsCol.FindOne(ctx, bson.M{"short_url": url.ShortUrl}).Decode(&existing)
		if err == nil {
			url = existing
		} else {
			_, _ = urlsCol.InsertOne(ctx, url)
		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": fmt.Sprintf("%s/%s", baseURL, url.ShortUrl),
	})
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]

	url, ok := urlDB[shortUrl]
	if !ok && mongoReady && urlsCol != nil {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		err := urlsCol.FindOne(ctx, bson.M{"short_url": shortUrl}).Decode(&url)
		if err == nil {
			urlDB[shortUrl] = url
			ok = true
		}
	}
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

// --------- MAIN ----------

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env not loaded:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := connectMongo(ctx)
	if err != nil {
		fmt.Println("Mongo connection failed:", err)
	} else {
		defer func() {
			_ = client.Disconnect(context.Background())
		}()
	}

	port := loadEnvOrDefault("PORT", "3000")
	addr := ":" + port
	fmt.Println("URL Shortener running on", addr)

	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", handleRedirect)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}


