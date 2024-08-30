package config

// import (
// 	"database/sql"
// 	"fmt"
// 	"log/slog"

// 	_ "github.com/lib/pq"

// 	"github.com/casbin/casbin/v2"
// 	xormadapter "github.com/casbin/xorm-adapter/v2"
// )

// func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
// 	cfg := LoadConfig()

// 	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD)
// 	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
// 			cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER,  cfg.DB_PASSWORD)
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		logger.Error("Error connecting to database", "error", err.Error())
// 		return nil, err
// 	}
// 	defer db.Close()
	
// 	_, err = db.Exec("DROP DATABASE IF EXISTS casbin")
// 	if err != nil {
// 		logger.Error("Error dropping Casbin database", "error", err.Error())
// 		return nil, err
// 	}
	
// 	connector := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
// 			cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD)
// 	adapter, err := xormadapter.NewAdapter("postgres", connector)
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("Adapter databazaga ulanmadi: %v", err))
// 		return nil, err
// 	}

// 	enforcer, err := casbin.NewEnforcer("config/model.conf", adapter)
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("Enforcer error: %v", err))
// 		return nil, err
// 	}

// 	err = enforcer.LoadPolicy()
// 	if err != nil {
// 		logger.Error(err.Error())
// 		return nil, err
// 	}

// 	policies := [][]string{
// 		{"admin", "/user/update", "PUT"},
// 		{"doctor", "/user/update", "PUT"},
// 		{"patient", "/user/update", "PUT"},
// 		{"admin", "/user/get", "GET"},
// 		{"doctor", "/user/get", "GET"},
// 		{"patient", "/user/get", "GET"},
// 		{"patient", "/medical/getAll", "GET"},
// 		{"admin", "/medical/getAll", "GET"},
// 		{"doctor", "/medical/getAll", "GET"},
// 		{"admin", "/medical/get/:id", "GET"},
// 		{"admin", "/medical/create", "POST"},
// 		{"doctor", "/medical/create", "POST"},
// 		{"admin", "/medical/update/:id", "PUT"},
// 		{"doctor", "/medical/update/:id", "PUT"},
// 		{"admin", "/medical/delete/:id", "DELETE"},
// 		{"doctor", "/medical/delete/:id", "DELETE"},
// 		{"patient", "/lifestyle/getAll", "GET"},
// 		{"admin", "/lifestyle/getAll", "GET"},
// 		{"doctor", "/lifestyle/getAll", "GET"},
// 		{"admin", "/lifestyle/get/:id", "GET"},
// 		{"admin", "/lifestyle/create", "POST"},
// 		{"doctor", "/lifestyle/create", "POST"},
// 		{"admin", "/lifestyle/update/:id", "PUT"},
// 		{"doctor", "/lifestyle/update/:id", "PUT"},
// 		{"admin", "/lifestyle/delete/:id", "DELETE"},
// 		{"doctor", "/lifestyle/delete/:id", "DELETE"},
// 		{"patient", "/wearable/getAll", "GET"},
// 		{"admin", "/wearable/getAll", "GET"},
// 		{"doctor", "/wearable/getAll", "GET"},
// 		{"admin", "/wearable/get/:id", "GET"},
// 		{"admin", "/wearable/create", "POST"},
// 		{"doctor", "/wearable/create", "POST"},
// 		{"admin", "/wearable/update/:id", "PUT"},
// 		{"doctor", "/wearable/update/:id", "PUT"},
// 		{"admin", "/wearable/delete/:id", "DELETE"},
// 		{"doctor", "/wearable/delete/:id", "DELETE"},
// 		{"patient", "/health/getAll", "GET"},
// 		{"admin", "/health/getAll", "GET"},
// 		{"doctor", "/health/getAll", "GET"},
// 		{"admin", "/health/get/:id", "GET"},
// 		{"admin", "/health/create", "POST"},
// 		{"doctor", "/health/create", "POST"},
// 		{"admin", "/health/delete/:id", "DELETE"},
// 		{"doctor", "/health/delete/:id", "DELETE"},
// 		{"patient", "/monitoring/realtime", "GET"},
// 		{"admin", "/monitoring/realtime", "GET"},
// 		{"doctor", "/monitoring/realtime", "GET"},
// 		{"patient", "/monitoring/daily", "GET"},
// 		{"admin", "/monitoring/daily", "GET"},
// 		{"doctor", "/monitoring/daily", "GET"},
// 		{"patient", "/monitoring/weekly", "GET"},
// 		{"admin", "/monitoring/realtime", "GET"},
// 		{"doctor", "/monitoring/realtime", "GET"},
// 	}

// 	_, err = enforcer.AddPolicies(policies)
// 	if err != nil {
// 		logger.Error(err.Error())
// 		return nil, err
// 	}

// 	err = enforcer.SavePolicy()
// 	if err != nil {
// 		logger.Error(err.Error())
// 		return nil, err
// 	}

// 	return enforcer, nil
// }
