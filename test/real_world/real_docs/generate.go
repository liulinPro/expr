package main

import (
	"fmt"

	"github.com/liulinpro/expr/docgen"
	"github.com/liulinpro/expr/test/real_world"
)

func main() {
	doc := docgen.CreateDoc(real_world.NewEnv())

	fmt.Println(doc.Markdown())
}
