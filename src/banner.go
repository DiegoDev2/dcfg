package src

import (
	"fmt"
)

func Banner() {
	const banner = `
  _____        _   ______ _ _
 |  __ \      | | |  ____(_) |
 | |  | | ___ | |_| |__   _| | ___  ___
 | |  | |/ _ \| __|  __| | | |/ _ \/ __|
 | |__| | (_) | |_| |    | | |  __/\__ \
 |_____/ \___/ \__|_|    |_|_|\___||___/

    By: @DiegoDev2
	`
	fmt.Println(banner)
}
