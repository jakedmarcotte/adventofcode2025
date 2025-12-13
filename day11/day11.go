package main

import (
	"bufio"
	"os"
	"strings"
)

type Server struct {
	id         string
	subservers []*Server
}

func (s Server) Visit() int {
	result := 0
	for _, sub := range s.subservers {
		if sub.id == "out" {
			result++
			continue
		} else if sub.id == "you" {
			continue
		}
		result += sub.Visit()
	}
	return result
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)

	serverLookup := make(map[string]*Server)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")
		parentID := strings.TrimSpace(splitLine[0])
		childrenStr := strings.TrimSpace(splitLine[1])
		childIDs := strings.Fields(childrenStr)
		if _, exists := serverLookup[parentID]; !exists {
			serverLookup[parentID] = &Server{
				id: parentID,
			}
		}
		parentServer := serverLookup[parentID]
		for _, childID := range childIDs {
			if _, exists := serverLookup[childID]; !exists {
				serverLookup[childID] = &Server{
					id: childID,
				}
			}
			childServer := serverLookup[childID]
			parentServer.subservers = append(parentServer.subservers, childServer)
		}

	}
	you := serverLookup["you"]
	result := you.Visit()
	println("Total servers reachable from 'you':", result)
}
