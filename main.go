package main

import (
 "context"
 "encoding/json"
 "log"
 "net/http"
 "os"
 "time"

 "go.mongodb.org/mongo-driver/v2/mongo"
 "go.mongodb.org/mongo-driver/v2/mongo/options"
 "go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
 uri := os.Getenv("MONGODB_URI")
 if uri == "" { log.Fatal("MONGODB_URI nao definida. Configure em Render > Environment.") }
 port := os.Getenv("PORT")
 if port == "" { port = "10000" }

 client, err := mongo.Connect(options.Client().ApplyURI(uri))
 if err != nil { log.Fatalf("Falha ao inicializar cliente MongoDB: %v", err) }
 defer client.Disconnect(context.Background())

 mux := http.NewServeMux()
 mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]string{"status":"running", "service":"colegio-unico-sync"})
 })
 mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
  ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
  defer cancel()
  if err := client.Ping(ctx, readpref.Primary()); err != nil {
   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusServiceUnavailable)
   json.NewEncoder(w).Encode(map[string]string{"status":"database_unavailable"})
   return
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]string{"status":"database_connected"})
 })
 mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  http.NotFound(w, r)
 })

 srv := &http.Server{Addr:":"+port, Handler:mux, ReadHeaderTimeout:5*time.Second}
 log.Printf("Servidor de teste iniciado na porta %s",port)
 log.Fatal(srv.ListenAndServe())
}
