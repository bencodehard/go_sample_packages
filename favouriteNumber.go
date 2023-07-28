package favoriteNumber

import (
	"fmt"
	"math/rand"
)

func FavoriteNumber() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
