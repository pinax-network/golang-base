# Golang base package

This is the base package containing useful modules for all kind of applications.

## List of modules

### Audit

This module provides a way to do audit logs on resources. It can log resources being created, deleted or updated. Out
of the box it comes with sinks to log to file, console or MySQL.

#### Example

```go
// create a file sink writing to /var/log/audits/audit.log
fileSink, _ := audit.NewFileSink("/var/log/audits/audit.log")

// initialize the global audit singleton to use that file sink
audit.InitializeAuditLog(fileSink)

// log the creating of a resource at the current time
audit.LogCreateResource(userId, resourceId, myResource, time.Now())
```

#### Prometheus Metrics

```
# HELP audit_created_resources Counter for created resources.
# TYPE audit_created_resources counter
audit_created_resources 27
# HELP audit_deleted_resources Counter for deleted resources.
# TYPE audit_deleted_resources counter
audit_deleted_resources 0
# HELP audit_updated_resources Counter for updated resources.
# TYPE audit_updated_resources counter
audit_updated_resources 13
```

### Config

Provides some helpers to easily load and validate config files from file.

### Consul

This modules wraps the Consul KV Api and provides some easy ways to load and store different data types without
having to handle any type conversions.

### Database

Contains a handler for MySQL connection pools that will connect to multiple MySQL instances simultaneously and keeps
track which instances are healthy by checking the cluster or database status regularly.

#### Prometheus Metrics

```
# HELP database_connections_healthy Number of currently healthy connections in the MySQL connection pool.
# TYPE database_connections_healthy gauge
database_connections_healthy 3
# HELP database_connections_unhealthy Number of currently unhealthy connections in the MySQL connection pool.
# TYPE database_connections_unhealthy gauge
database_connections_unhealthy 0
# HELP database_no_healthy_connection_errors Error counter for 'no healthy connections available'
# TYPE database_no_healthy_connection_errors counter
database_no_healthy_connection_errors 0
```

### dfuse

A wrapper for the dfuse client. Deprecated, use Substreams instead. 

### Elastic

This module contains a wrapper for the Elastic client to easily bulk index documents. 

### Log

Provides a logging library which utilizes zap under the hood. It comes with helpers and sane defaults to easily
instantiate loggers.

### Middleware

This module provides multiple middlewares for the Gin framework. Middlewares contained are:

* `ApiKeyMiddleware` adding api key authentication to endpoints
* `AuthMiddleware` checks if a JWT contains a certain permission
* `Errors` deals with unhandled errors within a request, logs them and translates them into a proper error response 
* `JwksMiddleware` checks json web tokens issued by Auth0. It will also keep track of the current keys that are used by Auth0 and updates on key rotations automatically.
* `LanguageMiddleware` parses the custom `X-Accept-Language` header and injects the proper language into the context
* `Recovery` catches panics within the application and translates them into proper responses
* `ReverseProxyMiddleware` proxies requests to another API

### Response

Provides a common HTTP response format for successful responses as well as errors.

### Sanitizer

This module contains a struct field sanitizer. 

### Sendgrid

Provides a Sendgrid client.

### Shufti

Provides a Shufti client.

### Twilio 

Provides a Twilio client.

### Validate

Provides some custom validation rules.
