package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	Mymap      map[string]string
	MasterName string
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		nline := strings.TrimSpace(line)
		if len(nline) == 0 {
			continue
		}

		// ignore the ones that start with '#' or ';'
		if nline[0] == '#' || nline[0] == ';' {
			continue
		}
		fmt.Println("nline[0]:", string(nline[0]))

		//find the c.MasterName
		if strings.HasPrefix(nline, "[") && strings.HasSuffix(nline, "]") {
			c.MasterName = nline[1 : len(nline)-1]
		}

		n := strings.Index(nline, "=")
		if n < 0 {
			continue
		}
		nb := strings.TrimSpace(nline[0:n])
		ne := strings.TrimSpace(nline[n:len(nline)])

		if np := strings.Index(nline, "#"); np > -1 {
			ne = strings.TrimSpace(nline[n+1 : np])
		}
		key := c.MasterName + "=" + nb
		c.Mymap[key] = ne

	}
}

func (c *Config) Read(node, key string) string {
	key = node + "=" + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}
