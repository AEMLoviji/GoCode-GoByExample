package main

import "fmt"

func main() {
	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	fmt.Printf("days=%v\n\n", days)

	weekdays := days[:5]
	weekdays[0] = "ALMA"
	fmt.Printf("weekdays=%v\n", weekdays)
	fmt.Printf("len(weekdays)=%d\n", len(weekdays))
	fmt.Printf("cap(weekdays)=%d\n\n", cap(weekdays))

	weekend := days[5:]
	fmt.Printf("weekend=%v\n", weekend)
	fmt.Printf("len(weekend)=%d\n", len(weekend))
	fmt.Printf("cap(weekend)=%d\n\n", cap(weekend))

	retired := days[:0]
	fmt.Printf("retired=%v\n", retired)
	fmt.Printf("len(retired)=%d\n", len(retired))
	fmt.Printf("cap(retired)=%d\n\n", cap(retired))

	fmt.Printf("my vacation=%v\n", days[3:5])
}
