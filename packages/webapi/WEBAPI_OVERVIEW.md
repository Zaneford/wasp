This page contains requirements to create a proper API documentation and client generation.

# Setup / Workflow

The webapi uses Swagger to document all routes, http codes and request/response models.

This gives an overview of the whole API on a separate website, which is to be found on `http://localhost:9090/doc`.

Furthermore, the Swagger standard allows the generation of client libraries. 

It's not required to create an API client manually. 

Inside the Wasp repository, the generated client can be found inside `clients/apiclient`.

It is mainly used by the cluster tests and the `wasp-cli`.

## Ensuring consistency

The API uses `echoswagger` which wraps around the Echo routing system. It automatically adds a new documentation for each (`.GET`|`.POST`...) definition.

Request and response models, returned status codes, query and path variables _need_ to be defined manually.

If this definition is missing, it will result in an incomplete api client and the new feature is unusable from generated clients.

Adding the documentation is simple though. See: `controllers/chains/controller.go`: `RegisterAdmin`: `adminAPI.GET("chains", c.getChainList,`

### Model declaration

Whenever a new request or response model is required, make sure to follow this layout:

```go
type ContractInfoResponse struct {
	Description string `json:"description" swagger:"desc(The description of the contract.),required"`
	HName       string `json:"hName" swagger:"desc(The id (HName as Hex)) of the contract.),required"`
	Name        string `json:"name" swagger:"desc(The name of the contract.),required"`
	ProgramHash string `json:"programHash" swagger:"desc(The hash of the contract. (Hex encoded)),required"`
}
```
Adding the `json` tag is mandatory. Always define the name in camelCase.

The `required` tag is needed when the returned property is not nullable, which is mostly the case in the API.


If possible, add a description (`desc`) to the property.

`uint` types have a special tag requirement which, if missed, will be interpreted as `int`. See [#Uints](#Uints))

`int64`, `uint64` and types above are unsupported, as JavaScript does not offer the precision required to parse it from JSON. 

They need to be sent as `string`, and the documentation should point that out:

```go
type AssetsResponse struct {
	BaseTokens   string         `json:"baseTokens" swagger:"required,desc(The base tokens (uint64 as string))"`
	NativeTokens []*NativeToken `json:"nativeTokens" swagger:"required"`
}
```


## Regenerating the client

If the webapi was changed, a regeneration of the client libraries is required.

To be able to generate the client, `openapi-generator` is required to be installed on your system.

Temporarily this version of the generator is mandatory: https://github.com/lmoe/openapi-generator/tree/feature/go_gen_unsigned_int

After installing, the api client can be regenerated by calling `clients/apiclient/generate_client.sh`.

# Swagger documentation pitfalls

To ensure a properly generated api client, a few common pitfalls are documented here and need to be considered when changing the API.

## Uints

UInts are unsupported by the Swagger standard. It only knows signed integers and decimals.

Usually openapi api generators will treat swagger integers as uint types, if the documented property contains a min value of at least 0. 
This allows generation of clients with proper uint typing. 

Therefore, all Uints (8-32) need to have a min(0) or min(1) documented. See: `models/chain.go` => `ChainInfoResponse`: `MaxBlobSize`

Example model:

```go
type StateTransaction struct {
	StateIndex    uint32 `json:"stateIndex" swagger:"desc(The state index),required,min(1)"`
	TransactionID string `json:"txId" swagger:"desc(The transaction ID),required"`
}
```

### uints in path parameters 

Paths like `/accounts/account/:id` are mostly documented with `.AddParamPath`. It automatically gets the proper type and documents it.

The only exception are uints. If the query path requires uints, it is mandatory to use `.AddParamPathNested` instead. 

Those properties also require a `min(0)` or `min(1)`.

See: `controllers/corecontracts/controller.go` at route `chains/:chainID/core/blocklog/blocks/:blockIndex` => `getBlockInfo`: `blockIndex`. 

Those properties need to be named the same way as the parameters in the route. The linter will complain about unused properties. 
Therefore, a `//nolint:unused` is required in these cases.

### (u)int 64 / big.Int

All number types above 58 bit are unsupported in JavaScript when consumed via JSON. Therefore, those types need to be sent as strings by the server. 

Unfortunately, this means the client has to decode these string numbers as a proper integer type. 

The documentation should point out that these strings are actually numbers and should be treated as such. 

See: `models/core_accounts.go` => `AccountNonceResponse`: `Nonce`

## Pointers and 'required'

By default, all properties are generated as pointers. Even if the webapi models are only using references.

This can be changed by supplying a 'required' tag in the swagger section.

See: `models/vm.go` => `ReceiptResponse`: `Request` (string)

Sometimes having an explicit nullable type is required though. 

In this example, the receipt response sometimes has an error, sometimes not.

See: `models/vm.go` => `ReceiptResponse`: `Error` (*ReceiptError)

By omitting the `required` tag, the property is marked as nullable and the error can be properly checked by `Error == nil` by the client.