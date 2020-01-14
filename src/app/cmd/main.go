/* vim:set ts=4 sts=2 sw=4 tw=0 fenc=utf-8: */

package main

import (
	estimate "handler/estimate"
)

func main() {
	handler := estimate.NewHandler()
	handler.Main()
	return
}

