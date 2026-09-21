# Opt-in administrative service authentication

`Authenticate` continues to accept user identities only. Services may explicitly
install `AuthenticateWithServiceClients(extractUser, config.AdminServiceClients)`
on their administrative route group, followed by the existing `CheckPermissions`
handler. Other route groups must keep the existing user-only middleware.

Configure an Auth0 client grant for the API's `admin` scope. Tokens must pass the
configured RS256/JWKS signature, issuer, audience and time validation. Machine
access additionally requires `gty: client-credentials`, a future `exp`, exactly
`<allowed-client-id>@clients` as the subject, a matching `azp` or `client_id`
claim, and the `admin` scope. If both client claims exist both must match. A
`permissions` claim alone does not replace the scope check.

```yaml
admin_service_clients:
  - client_id: ExactCaseSensitiveClientID
    principal_guid: server-configured-service-principal-guid
```

An omitted/empty list disables machine access. Invalid entries or duplicate
client IDs fail closed for the entire list. Wildcards are not accepted. Entries
are a list because Viper lowercases map keys, while Auth0 client IDs are case
sensitive. The list is copied when the handler is constructed; reload/restart to
apply a changed allowlist. Removing a client from this configuration revokes its
access even if its token has not expired.

The principal GUID is taken exclusively from this server-side mapping; token
custom claims cannot choose it. With `extractUser=true`, provide a `UserService`
that rechecks the mapped account on every request and rejects missing, banned or
deleted accounts and accounts whose stored Auth0 subject does not match the
machine subject. The returned user must have a positive local ID and the exact
mapped GUID. This keeps database authorization and audit foreign keys valid.
With `extractUser=false`, use a stable nonempty service identifier as the GUID;
no database lookup is performed.

Before the lookup, context contains the original Auth0 subject, provider
`client-credentials`, application ID, mapped GUID and `admin` permission.
`ServiceClientIDContextKey` identifies the application separately from any human
operator asserted by a trusted proxy. Never treat such operator metadata as
independently authenticated identity. Preserve the service identity in audit
records, and record human attribution separately.

Authentication errors never include the token or untrusted claim contents.
Error/recovery request logging includes the HTTP method and matched route
pattern only: no headers, cookies, query parameters or concrete path parameters.
Issuer and key-ID validation is strict for user tokens as well, and JWKS network
requests have a timeout, refuse redirects and cap response bytes.

Validation: `go test -race ./middleware` covers real RSA signatures, denied
identities/scopes/claims, expiration, service-account lookup failures, unchanged
user authentication, secret-free error/recovery logs and concurrent key reads.
