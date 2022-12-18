package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	return emoji.NormalizeShortCode("Hello :world_map:")
}
