func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    if n == 0 {
        return 0
    }

    // pair each car's position with its speed
    type Car struct {
        position int
        speed    int
    }

    cars := make([]Car, n)
    for i := range cars {
        cars[i] = Car{position[i], speed[i]}
    }

    // sort descending by position — closest to target first
    sort.Slice(cars, func(i, j int) bool {
        return cars[i].position > cars[j].position
    })

    // monotonic stack — track arrival times of each fleet
    stack := []float64{}
    for _, car := range cars {
        timeToTarget := float64(target-car.position) / float64(car.speed)

        // a car catches up to the car ahead → same fleet, discard
        if len(stack) == 0 || timeToTarget > stack[len(stack)-1] {
            stack = append(stack, timeToTarget)
        }
    }

    return len(stack)
}