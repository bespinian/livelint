package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/k8s-app-benchmarks/livelint/pkg/livelint"
	"github.com/k8s-app-benchmarks/livelint/pkg/opa"
	"github.com/k8s-app-benchmarks/livelint/pkg/promql"
)

func main() {
	checks := []livelint.Check{
		opa.NewCheck("my rego string"),
		promql.NewCheck("my promql query"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results := []livelint.CheckResult{}

		for _, c := range checks {
			r, err := c.Run()
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, r)
		}

		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
