package main

import (
	"fmt"

	abstract_factory "github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory"
	"github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory/apple"
)

func main() {
	var deviceFactory abstract_factory.DeviceFactory = &apple.Apple{}
	smartphone := deviceFactory.CreateSmartphone()
	tablet := deviceFactory.CreateTablet()
	fmt.Println(smartphone.Price())
	fmt.Println(tablet.Price())
}
