package supertoken

import (
	"github.com/supertokens/supertokens-golang/supertokens"
)

func SuperTokens() {
	err := supertokens.Init(SuperTokensConfig)

	if err != nil {
		panic(err.Error())
	}
}
