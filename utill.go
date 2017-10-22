package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
)

func logTopicErr(topic string, msg string, err error) {
	logrus.WithFields(logrus.Fields{
		"topic": topic,
		"event": msg,
	}).Error(err)
}

func readHosts(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, nil
}
