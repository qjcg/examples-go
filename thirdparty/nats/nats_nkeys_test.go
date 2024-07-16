package nats

import (
	"testing"

	"github.com/nats-io/nkeys"
)

// Create team keys for decentralized JWT auth, including:
//   - 1x operator keypair with signing keys
//   - 1x system account keypair
//   - 1x system user keypair
//   - 1x team account keypair
//   - 3x team user/service keypairs
func TestCreateTeamKeys(t *testing.T) {
	check := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create operator KP.
	operatorKP, err := nkeys.CreateOperator()
	check(err)

	operatorSeed, err := operatorKP.Seed()
	check(err)
	t.Logf("operatorSeed: %s\n", operatorSeed)

	// Create operator signing KP.
	operatorSigningKP, err := nkeys.FromSeed(operatorSeed)
	check(err)
	operatorSigningSeed, err := operatorSigningKP.Seed()
	check(err)
	t.Logf("operatorSigningSeed: %s\n", operatorSigningSeed)

	// Create account KP.
	accountKP, err := nkeys.CreateAccount()
	check(err)

	accountSeed, err := accountKP.Seed()
	check(err)

	t.Logf("accountSeed: %s\n", accountSeed)
}
