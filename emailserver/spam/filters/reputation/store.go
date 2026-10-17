package reputation

import "net"

type Store struct {
	ipScore     map[string]float64
	domainScore map[string]float64
}

func NewStore() *Store {
	return &Store{
		ipScore:     make(map[string]float64),
		domainScore: make(map[string]float64),
	}
}

func (s *Store) SetIP(ip string, score float64) {
	s.ipScore[ip] = score
}

func (s *Store) SetDomain(domain string, score float64) {
	s.domainScore[domain] = score
}

func (s *Store) Score(ip, domain string) float64 {
	return s.ipScore[ip] + s.domainScore[domain]
}

func ExtractIP(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return addr
	}
	return host
}
