/* vim:set ts=4 sts=2 sw=4 tw=0 fenc=utf-8: */

package main

import (
	invoice "handler/invoice"
)

func main() {
	invoice := invoice.NewHandler()
	invoice.Main()
	return
}

