package testimpl

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	managementGroupName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "management_group_name")
	managementGroupID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "management_group_id")
	require.NotEmpty(t, managementGroupName, "management_group_name output must not be empty")
	require.NotEmpty(t, managementGroupID, "management_group_id output must not be empty")

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err, "failed to create Azure credential")

	client, err := armmanagementgroups.NewClient(credential, nil)
	require.NoError(t, err, "failed to create management groups client")

	response, err := client.Get(context.Background(), managementGroupName, nil)
	require.NoError(t, err, "failed to get management group from Azure API")
	require.NotNil(t, response.Name, "management group name must not be nil")

	assert.Equal(t, managementGroupName, *response.Name, "management group name must match Terraform output")
	assert.Equal(t, managementGroupID, *response.ID, "management group ID must match Terraform output")
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	managementGroupName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "management_group_name")
	require.NotEmpty(t, managementGroupName, "management_group_name output must not be empty")

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err, "failed to create Azure credential")

	client, err := armmanagementgroups.NewClient(credential, nil)
	require.NoError(t, err, "failed to create management groups client")

	response, err := client.Get(context.Background(), managementGroupName, nil)
	require.NoError(t, err, "failed to get management group from Azure API")
	require.NotNil(t, response.Name, "management group name must not be nil")
	assert.Equal(t, managementGroupName, *response.Name, "management group name must match Terraform output")
}
