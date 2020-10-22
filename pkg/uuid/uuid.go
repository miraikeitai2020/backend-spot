package uuid

import(
	"log"
	"strings"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(uuidObj.String(), "-")[0]
}
