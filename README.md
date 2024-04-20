# Ben's golangPlayground
Repo containing a variety of testing and exploration in Go.

## Database Stuff

### Components of an Object
1. Schema defined in SQL
    A SQL file containing the actual DB semantics / schema for the migrations
2. DB Model & Functions
    For the specific DB, the model should be created, so that you can marshall to and from.

    Alongside the model, the common functions associate with that data should be created too.
    e.g:
        - GetByID
        - Create
        - DeleteByID
        - ModifyByID
        - AddXorYByID
3. Project Model (optional)
    The project may contain a second object type that is more generalised. The DB model should be able to marhshall to and from this type.

    This provides a common place for multiple different elements to interact with data. For example, the RestAPI and the DB.


### ByPrimaryKey or ByName (unique constraint)?
When I create functions for my data, how should I interact with it? Should I do it via Primary Key, such as ID, or should I do it by a unique constraint such as name?

- If you are using a unique constraint, you must index it to imrove performance.
- Primary Column can never be Null, where unique constraint can be.
- One primary key per table.

Primary keys should generally not be exposed, as they somewhat indicate implementation. e.g:
```
/api/app/1
/api/app/2
/api/app/3
```
Then it makes clear how to attempt ANY app access.

### Examples
A good example of PSQL schema:
https://github.com/mattermost/mattermost/tree/master/server/channels/db/migrations/postgres

```sql
CREATE INDEX IF NOT EXISTS idx_teams_name ON teams (name);

ALTER TABLE teams DROP COLUMN IF EXISTS allowopeninvite;
ALTER TABLE teams DROP COLUMN IF EXISTS lastteamiconupdate;
ALTER TABLE teams DROP COLUMN IF EXISTS description;
ALTER TABLE teams DROP COLUMN IF EXISTS groupconstrained;

DROP INDEX IF EXISTS idx_teams_invite_id;
DROP INDEX IF EXISTS idx_teams_update_at;
DROP INDEX IF EXISTS idx_teams_create_at;
DROP INDEX IF EXISTS idx_teams_delete_at;
DROP INDEX IF EXISTS idx_teams_scheme_id;

DROP TABLE IF EXISTS teams;
```

The above snippet shows the use of `IF NOT EXISTS` which applies the schema if it is not already.

It also demonstrates how you can deal with legacy schemas - via `ALTER TABLE` and `DROP INDEX IF EXISTS`.


SQL Objects and Functions
https://github.com/mattermost/mattermost/tree/master/server/channels/store/sqlstore
https://github.com/mattermost/mattermost/blob/master/server/channels/store/sqlstore/team_store.go

The above is a link to the `team` object code, for managing inside the SQL store. 




General Models for the Application
https://github.com/mattermost/mattermost/tree/master/server/public/model

Above contains what looks like a big list of models that is used by any part of the application. This indicates that many parts of the application will marshall to and from these models.

It also contains many helper functions, such as validations, to ensure the requirements are met.


https://github.com/mattermost/mattermost/blob/master/server/public/model/utils.go

utils.go has some utility functions, as well as the error types that can be used across the application.

### GORM or no GORM
When should I use GORM?
1. Rapid development
2. Applications with complex relationships
3. CRUD heavy applications - this makes the application simpler and more maintanable.
4. Non-performant - GORM adds some overhead due to the abstraction layer.

When should I not?
1. High-performance
2. Learning purposes - GORM obscures the underlying principles of DBA.
3. When you need full SQL control
4. Simple Applications - GORM may be overkill 

The general consensus for ORM and/or GORM, is bad. It adds complexity, and obscurity. It is difficult to manage after your application passes a certain scale, and also reduces application performance.

If possible, learning and utilising SQL directly is a better option.

```golang
type AppError struct {
	Id              string `json:"id"`
	Message         string `json:"message"`               // Message to be display to the end user without debugging information
	DetailedError   string `json:"detailed_error"`        // Internal error string to help the developer
	RequestId       string `json:"request_id,omitempty"`  // The RequestId that's also set in the header
	StatusCode      int    `json:"status_code,omitempty"` // The http status code
	Where           string `json:"-"`                     // The function where it happened in the form of Struct.Func
	IsOAuth         bool   `json:"is_oauth,omitempty"`    // Whether the error is OAuth specific
	SkipTranslation bool   `json:"-"`                     // Whether translation for the error should be skipped.
	params          map[string]any
	wrapped         error
}
```

The above is the error struct that they work with. This comes with a `ToJSON` function, which would be useful for API error messages.

### This is a LOT of code to write - SQLBoiler

```
go get -u github.com/volatiletech/sqlboiler
go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
```

SQLBoiler is basically ORM, without GORM. It generates code that queries a live DB, so it is SQL FIRST, where other ORM's are Model first.

Should I use it? I don't know.
PROS:
- Speed of Development
- Generation that can be defined in native SQL code 

CONS:
- Complex generated code that still requires integration with REST
- Layer of abstraction
- For interaction, I'm not writing SQL, I'm writing SQLBoiler code 
- An added step to the pipeline
- ORM complexities

For now, I don't think it's worth using.

## OAuth // OIDC

### What they both do 
#### OAuth 2.0
- An authorization framework that enables applications to obtain limited acess to user accounts on an HTTP service.
- It delegates user authentication to the service hosting the user account and authorzing third=party applications to access the user account.

Key components:
- Resource owner - the end user who owns the data
- Client - the application requesting access to the users account
- Authorization server - THe server that authenticates the resource owner and issues access tokens after getting proper authorization
- Resource server - the server hosting the protected resources. This is the repostiory im working on here in this case. It will accept server requests based on access tokens.

There are certain types of OAuth 2.0 flows that can offer authentication in different ways:
- Authorization - Used by the clients that can maintain confidentiality of the client secret
- Implicit - Designed for clients that cannot securely store the client secret
- Resource owner password credentials - Directly exchange username and password for an access token
- Client Credentials - Client can request an access token using only its client credentials


##### Examples
An example of web routes with OAuth can be found here:
https://github.com/mattermost/mattermost/blob/master/server/channels/web/oauth.go

It uses the stdlib. 

#### OIDC
OIDC is built **on top of oauth** and adds an identity layer. It enables clients (our service in this case) to verify the identity of the end-user based on the authentication performed by an authorization server, and also obtain profile information.

Key features of OIDC
- ID Token - a JWT that contains the users profile infromation returned from the server
- UserInfo Endpoint - standard way to obtain user data 
- Discovery Document - JSON document at a well known URL contains the OIDC provider's config information.


#### The differences / Similarities
OAuth2 provides the foundational authentication mechanism that OIDC builds upon. OIDC introduces an ID token that gives the client a way to authenticate the user.

OIDC is an authentication layer on top of OAuth 2.0.

### Implementation
This typically invovles redirecting a user to an OIDC provider and then handling the callback from the OIDC provider after authentication is complete.

```go
import (
    "golang.org/x/oauth2"
    "github.com/coreos/go-oidc"
)
```


## REST
### Code Generation
Similarly to SQL, there is a LOT of boiler plate code for REST implementations. The goal with this research was to find a way to alleviate some of the heavy lifting without losing control of the REST implementation.

I wanted to try to ensure that we could use OpenAPI, alongside Swagger for user experience, so I attempted to generate the code from the OpenAPI spec.

1. openapi-generator-cli
This generates an entire golang program which becomes cumbersome to edit and maintain as soon as you wish to make changes - doesn't seem like the right option.

```makefile
	@sudo rm -rf ${PWD}/api/gen/
	@docker run --rm -v ${PWD}/api:/local openapitools/openapi-generator-cli generate \
    -i /local/openapi.yaml \
    -g go-server \
    -o /local/gen/ \
    -c /local/config.yaml
	@sudo rm ${PWD}/api/gen/main.go
	@sudo rm ${PWD}/api/gen/go.mod
	@sudo rm ${PWD}/api/gen/README.md
	@sudo rm ${PWD}/api/gen/Dockerfile
```
2. oapi-codegen
https://github.com/deepmap/oapi-codegen
A specific piece of tooling for golang.
It offers more customisation and seems a lot closer to what I am looking for.
```bash
oapi-codegen -package gen -generate models api/oapi.yaml > api/model-gen.go
```
The above is a great way to only generate model code - which saves a large chunk of time while also retaining full control of the app.



### How does Generated Code fit into my manually written Code?

- At some point it needs to interact with the rest of my app; where does it do this?



- If I need to make changes to generated code, how can I achieve this without changing the generated files?



- Can I generate code that also generates swagger comments?
OAPI yamls have all information required. 

## Running HTTPS


