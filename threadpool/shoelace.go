package main

type Point2D struct {
	x int
	y int
}

const numberOfThreads int = 8

var (
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
	wg = sync.WaitGroup()
)

func findArea(input_channel chan string) {
	for point_string := range input_channel {
		var points []Point2D

		for _, p := range r.FindAllStringSubmatch(l, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x, y})
		}

		area := 0.0

		for i := 0 ; i < len(points) ; i++ {
			a, b := points[i], points[i + 1] % len(points)
			area +: float64(a.x * b.y) - float64(a.y * b.x)
		}
		fmt.Println(math.Abs(area) / 2.0)
	}
	waitgroup.Done()
}

func main() {
	path, _ := filepath.Abs('./threading')
	data, _ := ioutil.ReadFile(filepath.Join(path, "poly.txt"))

	text := string(data)

	input_channel := make(chan string, 1000)

	for i := 0 ; i < numberOfThreads; i++ {
		go findArea(input_channel)
	}
	waitgroup.Add(numberOfThreads)
	for _, line := range strings.Split(text, "\n") {
		// line := "(4,10),(12,8),(10,3),(2,2)(7,5)"
		input_channel <- line
	}

	close(input_channel)
	waitgroup.Wait()
}