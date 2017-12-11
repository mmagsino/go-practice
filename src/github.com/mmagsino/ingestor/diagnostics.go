package ingestor
import (
	"fmt"
	"github.com/mmagsino/ingestor/internal/repository"
)

func DiagnosticGroupById(id int32) {
	group := repository.FindById(id)
	fmt.Printf("%#v\n\n", group)
}
func DiagnosticGroups() {
	groups := repository.FindAll()
	for _, group := range groups {
		fmt.Printf("%#v\n\n", group)
	}
}