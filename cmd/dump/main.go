//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/yutopp/go-flv"
	flvtag "github.com/yutopp/go-flv/tag"
	"io"
	"os"
)

var (
	inFilePathOpt = flag.String("i", "", "Input file")
)

func init() {
	flag.Parse()
}

func main() {
	log.Printf("Open file: %s", *inFilePathOpt)

	file, err := os.Open(*inFilePathOpt)
	if err != nil {
		log.Fatalf("Failed to open file: %+v", err)
	}
	defer file.Close()

	dec, err := flv.NewDecoder(file)
	if err != nil {
		log.Fatalf("Failed to create decoder: %+v", err)
	}
	log.Printf("Header: %+v", dec.Header())

	var flvTag flvtag.FlvTag
	for {
		if err := dec.Decode(&flvTag); err != nil {
			if err == io.EOF {
				break
			}
			log.Warnf("Failed to decode: %+v", err)
		}

		log.Printf("Tag: %+v", flvTag)

		flvTag.Close() // Discard unread buffers
	}
}
