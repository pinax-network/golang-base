package base_repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/auth0/go-auth0/management"
	"github.com/eosnationftw/eosn-base-api/log"
	"go.uber.org/zap"
)

type Auth0Repository struct {
	auth0Management *management.Management
}

func NewAuth0Repository(config *Auth0Config) (repository *Auth0Repository, err error) {
	ctx := context.Background()
	man, err := management.New(config.ManagementDomain, management.WithClientCredentials(ctx, config.Client, config.Secret))
	if err != nil {
		return
	}

	repository = &Auth0Repository{
		auth0Management: man,
	}

	return
}

func (a *Auth0Repository) ListUsers(ctx context.Context, page int) *management.UserList {
	users, err := a.auth0Management.User.List(ctx, management.Page(page))
	if err != nil {
		panic(fmt.Errorf("failed to get user list from auth0 management api: %e", err))
	}

	return users
}

func (a *Auth0Repository) GetAuth0UserByFullId(ctx context.Context, id string) (*management.User, error) {
	user, err := a.auth0Management.User.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user with id %q from auth0 management api: %v", id, err)
	}

	return user, err
}

func (a *Auth0Repository) GetAuth0UserByIdProvider(ctx context.Context, provider, id string) (*management.User, error) {
	return a.GetAuth0UserByFullId(ctx, ConvertToFullId(provider, id))
}

func (a *Auth0Repository) ResendVerificationEmail(ctx context.Context, auth0Id string) {

	err := a.auth0Management.Job.VerifyEmail(ctx, &management.Job{
		UserID: &auth0Id,
	})

	if err != nil {
		log.Warn("error sending verification mail", zap.Error(err))
	}
}

func (a *Auth0Repository) UpdateEosnId(ctx context.Context, eosnId, auth0FullId string) {

	auth0User, err := a.auth0Management.User.Read(ctx, auth0FullId)
	if err != nil {
		panic(fmt.Errorf("failed to get user from auth0 management api: %e", err))
	}

	if auth0User.AppMetadata == nil {
		auth0User.AppMetadata = &map[string]interface{}{}
	}

	newAuth0User := &management.User{
		AppMetadata: auth0User.AppMetadata,
	}
	(*newAuth0User.AppMetadata)["eosn_id"] = eosnId

	err = a.auth0Management.User.Update(ctx, auth0FullId, newAuth0User)
	if err != nil {
		panic(fmt.Errorf("failed to update user with new metadata: %e", err))
	}
}

func (a *Auth0Repository) LinkAccount(ctx context.Context, fullAuth0Id, targetIdToken string) error {

	link, err := a.auth0Management.User.Link(ctx, fullAuth0Id, &management.UserIdentityLink{
		LinkWith: &targetIdToken,
	})

	if err != nil {
		log.Info("failed to link user", zap.String("auth0_id", fullAuth0Id), zap.Any("link", link))
		return err
	}

	log.Debug("successfully linked user", zap.Any("link", link))
	return nil
}

func (a *Auth0Repository) UnlinkAccount(ctx context.Context, fullAuth0Id, provider, userId string) {
	user, err := a.auth0Management.User.Unlink(ctx, fullAuth0Id, provider, userId)
	log.PanicIfError("failed to unlink user", err)
	log.Info("unlinked user succesfully", zap.Any("links", user))
}

func ConvertToFullId(auth0Provider, auth0Id string) string {
	return fmt.Sprintf("%s|%s", auth0Provider, auth0Id)
}

func ConvertFullIdToIdProvider(auth0FullId string) (auth0Provider string, auth0Id string, err error) {
	auth0Array := strings.Split(auth0FullId, "|")

	if len(auth0Array) == 2 && auth0Array[0] != "" && auth0Array[1] != "" {
		return auth0Array[0], auth0Array[1], nil
	} else if len(auth0Array) == 3 && auth0Array[0] == "oauth2" && auth0Array[1] != "" && auth0Array[2] != "" {
		return fmt.Sprintf("%s|%s", auth0Array[0], auth0Array[1]), auth0Array[2], nil
	}

	return "", "", fmt.Errorf("failed to parse auth0 full id, invalid format must be of provider|id or oauth2|provider|id, instead got '%s'", auth0FullId)
}
