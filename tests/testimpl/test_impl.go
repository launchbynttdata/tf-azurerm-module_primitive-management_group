package common

import (
	"os"
	"testing"

	"github.com/launchbynttdata/lcaf-component-terratest/types"
)

func TestComposableManagementGroup(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}
}
