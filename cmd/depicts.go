package main

import(
	"flag"
	"github.com/sfomuseum/go-sfomuseum-geojson/feature"
	"github.com/sfomuseum/go-sfomuseum-geojson/properties/sfomuseum"	
	"log"
)

func main() {

	flag.Parse()

	for _, path := range flag.Args() {

		f, err := feature.LoadFeatureFromFile(path)

		if err != nil {
			log.Fatal(err)
		}

		ids, err := sfomuseum.Depicts(f)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(ids)
	}
	
}
