package b221020

import "time"

func Print() {
    println(time.Now().UTC().Format(time.RFC3339))
}
