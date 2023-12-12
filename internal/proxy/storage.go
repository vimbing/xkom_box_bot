package proxy

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"xkomopener/internal/utils/helpers"
)

func (s *Storage) GetRandomProxy() string {
	return s.Proxies[helpers.RandomInt(0, len(s.Proxies)-1)]
}

func (s *Storage) AddProxy(proxy string) {
	s.Proxies = append(s.Proxies, proxy)
}

func (s *Storage) LoadFromFile() error {
	readFile, err := os.Open(os.Getenv("PROXIES_PATH"))

	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		proxySplit := strings.Split(fileScanner.Text(), ":")

		if len(proxySplit) != 4 {
			continue
		}

		host := proxySplit[0]
		port := proxySplit[1]
		user := proxySplit[2]
		password := proxySplit[3]

		proxyString := fmt.Sprintf("http://%s:%s@%s:%s", user, password, host, port)

		s.AddProxy(proxyString)
	}

	readFile.Close()

	return nil
}
