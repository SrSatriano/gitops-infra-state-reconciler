package main

import (
	"flag"
	"log"
	"time"

	"github.com/xerife/gitops-reconciler/pkg/drift"
)

func main() {
	repo := flag.String("repo", "", "git repository URL")
	interval := flag.Duration("interval", 60*time.Second, "poll interval")
	flag.Parse()

	if *repo == "" {
		log.Fatal("--repo required")
	}

	d := drift.NewDetector(*repo)
	log.Printf("reconciler started repo=%s interval=%s", *repo, *interval)
	for {
		changes, err := d.Scan()
		if err != nil {
			log.Printf("scan error: %v", err)
		} else if len(changes) > 0 {
			log.Printf("drift detected: %d files — reconciling", len(changes))
			// TODO: git checkout + apply
		}
		time.Sleep(*interval)
	}
}
