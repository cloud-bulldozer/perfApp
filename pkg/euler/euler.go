package euler

import (
	"fmt"
	"math"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rsevilla87/perfapp/internal/perf"
	"github.com/rsevilla87/perfapp/pkg/utils"
)

// HandleEuler Handle requests to compute euler number aproximation
func HandleEuler(w http.ResponseWriter, r *http.Request) {
	log.Println("Computing euler approximation")
	now := time.Now()
	calcEuler()
	elapsed := time.Now().Sub(now)
	insert := fmt.Sprintf("INSERT INTO euler VALUES ('%s', '%f')", now.Format(time.RFC3339), elapsed.Seconds())
	if err := perf.QueryDB(insert); err != nil {
		utils.ErrorHandler(err)
	} else {
		fmt.Fprintln(w, "Ok")
		log.Printf("Euler approximation computed in %f seconds", elapsed.Seconds())
		perf.HTTPRequestDuration.Observe(elapsed.Seconds())
	}
}

func calcEuler() {
	var n float64
	var x float64
	for math.E > x {
		x = math.Pow((1 + 1/n), n)
		n++
	}
}
