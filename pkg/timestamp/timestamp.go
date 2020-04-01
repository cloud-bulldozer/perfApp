package timestamp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rsevilla87/perfapp/internal/perf"
	"github.com/rsevilla87/perfapp/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// Tables Euler workload required tables
var Tables = map[string]string{"ts": "CREATE TABLE IF NOT EXISTS ts (date TIMESTAMP)"}

// Handler Handle timestamp requests
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Inserting timestamp record in table")
	now := time.Now()
	elapsed := time.Now().Sub(now)
	insert := fmt.Sprintf("INSERT INTO ts VALUES ('%s')", now.Format(time.RFC3339))
	if err := perf.QueryDB(insert); err != nil {
		utils.ErrorHandler(err)
	} else {
		fmt.Fprintln(w, "Ok")
		log.Printf("Timestamp inserted in %v ns", elapsed.Nanoseconds())
		perf.HTTPRequestDuration.Observe(elapsed.Seconds())
	}
}
