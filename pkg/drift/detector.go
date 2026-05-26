package drift

import "crypto/sha256"

type Detector struct {
	repo string
}

func NewDetector(repo string) *Detector {
	return &Detector{repo: repo}
}

func (d *Detector) Scan() ([]string, error) {
		_ = sha256.Sum256([]byte(d.repo))
	return nil, nil
}
