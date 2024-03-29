package api

import "github.com/google/uuid"

// PrefectClient returns clients for different aspects of our API.
//
//nolint:interfacebloat // we'll accept a larger PrefectClient interface
type PrefectClient interface {
	Accounts(accountID uuid.UUID) (AccountsClient, error)
	AccountMemberships(accountID uuid.UUID) (AccountMembershipsClient, error)
	AccountRoles(accountID uuid.UUID) (AccountRolesClient, error)
	Collections() (CollectionsClient, error)
	Teams(accountID uuid.UUID) (TeamsClient, error)
	Workspaces(accountID uuid.UUID) (WorkspacesClient, error)
	WorkspaceAccess(accountID uuid.UUID, workspaceID uuid.UUID) (WorkspaceAccessClient, error)
	WorkspaceRoles(accountID uuid.UUID) (WorkspaceRolesClient, error)
	WorkPools(accountID uuid.UUID, workspaceID uuid.UUID) (WorkPoolsClient, error)
	WorkQueues(accountID uuid.UUID, workspaceID uuid.UUID, workPoolName string) (WorkQueuesClient, error)
	Variables(accountID uuid.UUID, workspaceID uuid.UUID) (VariablesClient, error)
	ServiceAccounts(accountID uuid.UUID) (ServiceAccountsClient, error)
}
