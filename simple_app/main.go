package main

import (
	"fmt"
	util "simple_app/util"
)

func main() {

	print("\n")

	fmt.Println(util.TitleLower("Lorem Ipsum Dolor Sit Amet"))

	print("\n\n")

	res := util.WordStatistic(getArticle())

	fmt.Printf("%#v", res["TotalWordCounts"])
	print("\n")
	fmt.Printf("%#v", res["WordCountEveryWord"])
	print("\n")
	fmt.Printf("%#v", res["NumberOfWordShowUpOnce"])
	print("\n")
	fmt.Printf("%#v", res["WordHighestCount"])
	print("\n")
	fmt.Printf("%#v", res["WordSmallestCount"])
	print("\n")
}

func getArticle() string {
	return `Go, also known as Golang,
			[14]
			is a statically typed, compiled programming language designed at

			Google[15]
			by Robert Griesemer, Rob Pike, and Ken Thompson.
			[12]
			Go is syntactically similar to C,

			but with memory safety, garbage collection, structural typing,
			[6]
			and CSP-style concurrency.
			[16]

			Go was designed at Google in 2007 to improve programming productivity in an era of multicore,
			networked machines and large codebases.
			[23]
			The designers wanted to address criticism of other

			languages in use at Google, but keep their useful characteristics:[24]
			● Static typing and run-time efficiency (like C++)
			● Readability and usability (like Python or JavaScript)
			[25]
			● High-performance networking and multiprocessing
			The designers were primarily motivated by their shared dislike of C++.
			[26][27][28]

			Go was publicly announced in November 2009,[29]

			and version 1.0 was released in March

			2012.[30][31]

			Go is widely used in production at Google[32]

			and in many other organizations and

			open-source projects.
			In November 2016, the Go and Go Mono fonts which are sans-serif and monospaced respectively
			were released by type designers Charles Bigelow and Kris Holmes. Both fonts adhere to WGL4 and
			were designed to be legible with a large x-height and distinct letterforms by conforming to the DIN
			1450 standard.[33][34]
			In April 2018, the original logo was replaced with a stylized GO slanting right with trailing
			streamlines. However, the Gopher mascot remained the same.[35]
			In August 2018, the Go principal contributors published two ′′draft designs′′ for new language
			features, Generics and error handling, and asked Go users to submit feedback on them.[36][37]
			Lack
			of support for generic programming and the verbosity of error handling in Go 1.x had drawn
			considerable criticism.`
}
