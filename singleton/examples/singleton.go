package singleton

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	cities map[string]int
}

func (db singletonDatabase) GetPopulation(name string) int {
	return db.cities[name]
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		db := singletonDatabase{}
		cities, err := readData("./cities.txt")
		if err != nil {
			panic(err)
		}

		db.cities = cities
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}
