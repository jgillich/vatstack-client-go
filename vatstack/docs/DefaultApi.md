# \DefaultApi

All URIs are relative to *https://api.vatstack.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatch**](DefaultApi.md#CreateBatch) | **Post** /batches | Creates a batch object.
[**CreateEvidence**](DefaultApi.md#CreateEvidence) | **Post** /evidences | Creates an evidence object.
[**CreateQuote**](DefaultApi.md#CreateQuote) | **Post** /quotes | Creates a quote object.
[**CreateSupply**](DefaultApi.md#CreateSupply) | **Post** /supplies | Creates a supply object.
[**CreateValidation**](DefaultApi.md#CreateValidation) | **Post** /validations | Creates a validation object.
[**DeleteBatchById**](DefaultApi.md#DeleteBatchById) | **Delete** /batches/{id} | Deletes a batch object by the {id} path parameter.
[**DeleteSupplyById**](DefaultApi.md#DeleteSupplyById) | **Delete** /supplies/{id} | Deletes a supply object by the {id} path parameter.
[**GetAllBatches**](DefaultApi.md#GetAllBatches) | **Get** /batches | Retrieves all batch objects in order of creation, with the latest appearing highest.
[**GetAllEvidences**](DefaultApi.md#GetAllEvidences) | **Get** /evidences | Retrieves all evidence objects in order of creation, with the latest appearing highest.
[**GetAllHits**](DefaultApi.md#GetAllHits) | **Get** /hits | This endpoint can be useful if you want to perform a quick check against your hits count before initiating an API request.
[**GetAllQuotes**](DefaultApi.md#GetAllQuotes) | **Get** /quotes | Retrieves all quote objects in order of creation, with the most recent appearing highest.
[**GetAllRates**](DefaultApi.md#GetAllRates) | **Get** /rates | Retrieves all VAT rate objects, including standard VAT rates and VAT rates for digital products.
[**GetAllSupplies**](DefaultApi.md#GetAllSupplies) | **Get** /supplies | Retrieves all supply objects in order of creation, with the most recent appearing highest.
[**GetAllValidations**](DefaultApi.md#GetAllValidations) | **Get** /validations | Retrieves all validation objects in order of creation, with the most recent appearing highest.
[**GetBatchById**](DefaultApi.md#GetBatchById) | **Get** /batches/{id} | Retrieves a batch object by the {id} path parameter.
[**GetEvidenceById**](DefaultApi.md#GetEvidenceById) | **Get** /evidences/{id} | Retrieves an evidence object by the {id} path parameter.
[**GetQuoteById**](DefaultApi.md#GetQuoteById) | **Get** /quotes/{id} | Retrieves a quote object.
[**GetRateByIPAddress**](DefaultApi.md#GetRateByIPAddress) | **Get** /rates/geolocate | Retrieves a geolocated VAT rate by IP address.
[**GetSupplyById**](DefaultApi.md#GetSupplyById) | **Get** /supplies/{id} | Retrieves a supply object.
[**GetValidationById**](DefaultApi.md#GetValidationById) | **Get** /validations/{id} | Retrieves a validation object.
[**UpdateBatchById**](DefaultApi.md#UpdateBatchById) | **Put** /batches/{id} | Updates a batch object by the {id} path parameter.
[**UpdateEvidenceById**](DefaultApi.md#UpdateEvidenceById) | **Put** /evidences/{id} | Updates an evidence object by the {id} path parameter.
[**UpdateSupplyById**](DefaultApi.md#UpdateSupplyById) | **Put** /supplies/{id} | Updates a supply object by the {id} path parameter.



## CreateBatch

> Batch CreateBatch(ctx).Name(name).Queries(queries).Execute()

Creates a batch object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    name := "name_example" // string | Descriptive name of the batch object. Give it any name you like to serve for your reference.
    queries := []string{"Inner_example"} // []string | Array of VAT numbers to be queried. Should be an array of strings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateBatch(context.Background()).Name(name).Queries(queries).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBatch`: Batch
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Descriptive name of the batch object. Give it any name you like to serve for your reference. | 
 **queries** | **[]string** | Array of VAT numbers to be queried. Should be an array of strings. | 

### Return type

[**Batch**](Batch.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEvidence

> Evidence CreateEvidence(ctx).BankAddress(bankAddress).BillingAddress(billingAddress).IpAddress(ipAddress).Execute()

Creates an evidence object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    bankAddress := TODO // CreateEvidenceRequestBankAddress |  (optional)
    billingAddress := TODO // CreateEvidenceRequestBillingAddress |  (optional)
    ipAddress := TODO // CreateEvidenceRequestIpAddress |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEvidence(context.Background()).BankAddress(bankAddress).BillingAddress(billingAddress).IpAddress(ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEvidence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvidence`: Evidence
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEvidence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankAddress** | [**CreateEvidenceRequestBankAddress**](CreateEvidenceRequestBankAddress.md) |  | 
 **billingAddress** | [**CreateEvidenceRequestBillingAddress**](CreateEvidenceRequestBillingAddress.md) |  | 
 **ipAddress** | [**CreateEvidenceRequestIpAddress**](CreateEvidenceRequestIpAddress.md) |  | 

### Return type

[**Evidence**](Evidence.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuote

> Quote CreateQuote(ctx).Amount(amount).CountryCode(countryCode).Category(category).IpAddress(ipAddress).Validation(validation).Vat(vat).Execute()

Creates a quote object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    amount := int32(56) // int32 | Amount in **cents** (e.g. 100.50 must be expressed as 10050). This common common workaround prevents unexpected rounding issues.
    countryCode := "countryCode_example" // string | 2-letter ISO country code. If provided, the `ip_address` parameter will be ignored. (optional)
    category := "category_example" // string | Digital products category used for calculation. Supports `audiobook`, `broadcasting`, `ebook`, `eperiodical`, `eservice` and `telecommunication`. (optional)
    ipAddress := "ipAddress_example" // string | IP address to geolocate the VAT rate for. If neither IP address nor `country_code` is provided, it will be automatically determined from the request. (optional)
    validation := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of a [validation object](https://vatstack.com/docs/validations). This is useful if you let your customer enter a VAT number beforehand. Its valid value can affect `vat.amount`, `vat.rate` and `amount_total` when zero-rating. (optional)
    vat := TODO // CreateQuoteRequestVat |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateQuote(context.Background()).Amount(amount).CountryCode(countryCode).Category(category).IpAddress(ipAddress).Validation(validation).Vat(vat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuote`: Quote
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **int32** | Amount in **cents** (e.g. 100.50 must be expressed as 10050). This common common workaround prevents unexpected rounding issues. | 
 **countryCode** | **string** | 2-letter ISO country code. If provided, the &#x60;ip_address&#x60; parameter will be ignored. | 
 **category** | **string** | Digital products category used for calculation. Supports &#x60;audiobook&#x60;, &#x60;broadcasting&#x60;, &#x60;ebook&#x60;, &#x60;eperiodical&#x60;, &#x60;eservice&#x60; and &#x60;telecommunication&#x60;. | 
 **ipAddress** | **string** | IP address to geolocate the VAT rate for. If neither IP address nor &#x60;country_code&#x60; is provided, it will be automatically determined from the request. | 
 **validation** | **string** | Unique identifier of a [validation object](https://vatstack.com/docs/validations). This is useful if you let your customer enter a VAT number beforehand. Its valid value can affect &#x60;vat.amount&#x60;, &#x60;vat.rate&#x60; and &#x60;amount_total&#x60; when zero-rating. | 
 **vat** | [**CreateQuoteRequestVat**](CreateQuoteRequestVat.md) |  | 

### Return type

[**Quote**](Quote.md)

### Authorization

[secretKey](../README.md#secretKey), [publicKey](../README.md#publicKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSupply

> Supply CreateSupply(ctx).Amount(amount).CountryCode(countryCode).InvoiceNumber(invoiceNumber).Issued(issued).AmountRefunded(amountRefunded).Currency(currency).Description(description).Evidence(evidence).Name(name).Notes(notes).Validation(validation).Vat(vat).Execute()

Creates a supply object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    amount := int32(56) // int32 | Amount in **cents** (e.g. 100.50 must be expressed as 10050). This common common workaround prevents unexpected rounding issues.
    countryCode := "countryCode_example" // string | 2-letter ISO country code of the place of supply that is relevant for the `vat.rate`.
    invoiceNumber := "invoiceNumber_example" // string | A custom string for the invoice number issued to the customer. It’s advisable to follow sequential numbering.
    issued := "issued_example" // string | ISO date at which the invoice was issued to the customer.
    amountRefunded := int32(56) // int32 | Amount in cents refunded back to the customer. (optional)
    currency := "currency_example" // string | 3-letter ISO 4217 currency code used to charge the `amount`. The currency is used to correctly convert to your reporting currency. (optional)
    description := "description_example" // string | A custom string to describe the supplied item. (optional)
    evidence := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an [evidence object](https://vatstack.com/docs/evidences). The pieces of non-contradictory evidence contained therein will affect the `evidence_status`. (optional)
    name := "name_example" // string | A custom string for the name of the customer. (optional)
    notes := "notes_example" // string | A custom string for additional notes. (optional)
    validation := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of a [validation object](https://vatstack.com/docs/validations). This is useful if you let your customer enter a VAT number beforehand. Its valid value can affect `vat.amount`, `vat.rate` and `amount_total` when zero-rating. (optional)
    vat := TODO // CreateSupplyRequestVat |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSupply(context.Background()).Amount(amount).CountryCode(countryCode).InvoiceNumber(invoiceNumber).Issued(issued).AmountRefunded(amountRefunded).Currency(currency).Description(description).Evidence(evidence).Name(name).Notes(notes).Validation(validation).Vat(vat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSupply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSupply`: Supply
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSupply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **int32** | Amount in **cents** (e.g. 100.50 must be expressed as 10050). This common common workaround prevents unexpected rounding issues. | 
 **countryCode** | **string** | 2-letter ISO country code of the place of supply that is relevant for the &#x60;vat.rate&#x60;. | 
 **invoiceNumber** | **string** | A custom string for the invoice number issued to the customer. It’s advisable to follow sequential numbering. | 
 **issued** | **string** | ISO date at which the invoice was issued to the customer. | 
 **amountRefunded** | **int32** | Amount in cents refunded back to the customer. | 
 **currency** | **string** | 3-letter ISO 4217 currency code used to charge the &#x60;amount&#x60;. The currency is used to correctly convert to your reporting currency. | 
 **description** | **string** | A custom string to describe the supplied item. | 
 **evidence** | **string** | Unique identifier of an [evidence object](https://vatstack.com/docs/evidences). The pieces of non-contradictory evidence contained therein will affect the &#x60;evidence_status&#x60;. | 
 **name** | **string** | A custom string for the name of the customer. | 
 **notes** | **string** | A custom string for additional notes. | 
 **validation** | **string** | Unique identifier of a [validation object](https://vatstack.com/docs/validations). This is useful if you let your customer enter a VAT number beforehand. Its valid value can affect &#x60;vat.amount&#x60;, &#x60;vat.rate&#x60; and &#x60;amount_total&#x60; when zero-rating. | 
 **vat** | [**CreateSupplyRequestVat**](CreateSupplyRequestVat.md) |  | 

### Return type

[**Supply**](Supply.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateValidation

> Validation CreateValidation(ctx).Query(query).Type_(type_).Execute()

Creates a validation object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    query := "query_example" // string | VAT number that you want to validate.
    type_ := "type__example" // string | Restrict validation to a region. If not provided, the type is automatically determined based on the VAT number given. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateValidation(context.Background()).Query(query).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateValidation`: Validation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | VAT number that you want to validate. | 
 **type_** | **string** | Restrict validation to a region. If not provided, the type is automatically determined based on the VAT number given. | 

### Return type

[**Validation**](Validation.md)

### Authorization

[secretKey](../README.md#secretKey), [publicKey](../README.md#publicKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBatchById

> DeleteBatchById(ctx, id).Execute()

Deletes a batch object by the {id} path parameter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteBatchById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBatchById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSupplyById

> DeleteSupplyById(ctx, id).Execute()

Deletes a supply object by the {id} path parameter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteSupplyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSupplyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupplyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBatches

> GetAllBatches200Response GetAllBatches(ctx).Limit(limit).Page(page).Status(status).Execute()

Retrieves all batch objects in order of creation, with the latest appearing highest.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    limit := int32(56) // int32 | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. (optional)
    page := int32(56) // int32 | Integer for the current page. (optional)
    status := "status_example" // string | Show only objects with the given `status`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllBatches(context.Background()).Limit(limit).Page(page).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBatches`: GetAllBatches200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. | 
 **page** | **int32** | Integer for the current page. | 
 **status** | **string** | Show only objects with the given &#x60;status&#x60;. | 

### Return type

[**GetAllBatches200Response**](GetAllBatches200Response.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllEvidences

> GetAllEvidences200Response GetAllEvidences(ctx).Limit(limit).Page(page).Execute()

Retrieves all evidence objects in order of creation, with the latest appearing highest.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    limit := int32(56) // int32 | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. (optional)
    page := int32(56) // int32 | Integer for the current page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllEvidences(context.Background()).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllEvidences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllEvidences`: GetAllEvidences200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllEvidences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllEvidencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. | 
 **page** | **int32** | Integer for the current page. | 

### Return type

[**GetAllEvidences200Response**](GetAllEvidences200Response.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllHits

> Hit GetAllHits(ctx).Execute()

This endpoint can be useful if you want to perform a quick check against your hits count before initiating an API request.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllHits(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllHits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllHits`: Hit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllHits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllHitsRequest struct via the builder pattern


### Return type

[**Hit**](Hit.md)

### Authorization

[secretKey](../README.md#secretKey), [publicKey](../README.md#publicKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllQuotes

> GetAllQuotes200Response GetAllQuotes(ctx).Limit(limit).Page(page).Execute()

Retrieves all quote objects in order of creation, with the most recent appearing highest.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    limit := int32(56) // int32 | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. (optional)
    page := int32(56) // int32 | Integer for the current page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllQuotes(context.Background()).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllQuotes`: GetAllQuotes200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. | 
 **page** | **int32** | Integer for the current page. | 

### Return type

[**GetAllQuotes200Response**](GetAllQuotes200Response.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRates

> GetAllRates200Response GetAllRates(ctx).CountryCode(countryCode).Limit(limit).MemberState(memberState).Page(page).Execute()

Retrieves all VAT rate objects, including standard VAT rates and VAT rates for digital products.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    countryCode := "countryCode_example" // string | Filter results by a 2-letter ISO country code. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. (optional)
    memberState := true // bool | Boolean to filter results by EU Member State. (optional)
    page := int32(56) // int32 | Integer for the current page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllRates(context.Background()).CountryCode(countryCode).Limit(limit).MemberState(memberState).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllRates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRates`: GetAllRates200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | Filter results by a 2-letter ISO country code. | 
 **limit** | **int32** | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. | 
 **memberState** | **bool** | Boolean to filter results by EU Member State. | 
 **page** | **int32** | Integer for the current page. | 

### Return type

[**GetAllRates200Response**](GetAllRates200Response.md)

### Authorization

[secretKey](../README.md#secretKey), [publicKey](../README.md#publicKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSupplies

> GetAllSupplies200Response GetAllSupplies(ctx).CountryCode(countryCode).IssuedSince(issuedSince).IssuedUntil(issuedUntil).Limit(limit).Name(name).Page(page).Execute()

Retrieves all supply objects in order of creation, with the most recent appearing highest.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    countryCode := "countryCode_example" // string | The 2-letter ISO country code by which you want to filter your records. (optional)
    issuedSince := time.Now() // string | Show only objects where the `issued` date is this date or later. Format `YYYY-MM-DD`. (optional)
    issuedUntil := time.Now() // string | Show only objects where the `issued` date is this date or earlier. Format `YYYY-MM-DD`. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. (optional)
    name := "name_example" // string | Show only objects with that customer name. (optional)
    page := int32(56) // int32 | Integer for the current page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllSupplies(context.Background()).CountryCode(countryCode).IssuedSince(issuedSince).IssuedUntil(issuedUntil).Limit(limit).Name(name).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllSupplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSupplies`: GetAllSupplies200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllSupplies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSuppliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | The 2-letter ISO country code by which you want to filter your records. | 
 **issuedSince** | **string** | Show only objects where the &#x60;issued&#x60; date is this date or later. Format &#x60;YYYY-MM-DD&#x60;. | 
 **issuedUntil** | **string** | Show only objects where the &#x60;issued&#x60; date is this date or earlier. Format &#x60;YYYY-MM-DD&#x60;. | 
 **limit** | **int32** | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. | 
 **name** | **string** | Show only objects with that customer name. | 
 **page** | **int32** | Integer for the current page. | 

### Return type

[**GetAllSupplies200Response**](GetAllSupplies200Response.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllValidations

> GetAllValidations200Response GetAllValidations(ctx).Batch(batch).Limit(limit).Page(page).Query(query).RequestedSince(requestedSince).RequestedUntil(requestedUntil).Type_(type_).Execute()

Retrieves all validation objects in order of creation, with the most recent appearing highest.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    batch := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Show only objects that belong to a batch with the given identifier. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. (optional)
    page := int32(56) // int32 | Integer for the current page. (optional)
    query := "query_example" // string | The VAT number you want to search in the `query` field of your records. (optional)
    requestedSince := time.Now() // string | Show only objects where the `requested` date is this date or later. Format `YYYY-MM-DD`. (optional)
    requestedUntil := time.Now() // string | Show only objects where the `requested` date is this date or earlier. Format `YYYY-MM-DD`. (optional)
    type_ := "type__example" // string | Show only objects of specified type field. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllValidations(context.Background()).Batch(batch).Limit(limit).Page(page).Query(query).RequestedSince(requestedSince).RequestedUntil(requestedUntil).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllValidations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllValidations`: GetAllValidations200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllValidations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllValidationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batch** | **string** | Show only objects that belong to a batch with the given identifier. | 
 **limit** | **int32** | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20. | 
 **page** | **int32** | Integer for the current page. | 
 **query** | **string** | The VAT number you want to search in the &#x60;query&#x60; field of your records. | 
 **requestedSince** | **string** | Show only objects where the &#x60;requested&#x60; date is this date or later. Format &#x60;YYYY-MM-DD&#x60;. | 
 **requestedUntil** | **string** | Show only objects where the &#x60;requested&#x60; date is this date or earlier. Format &#x60;YYYY-MM-DD&#x60;. | 
 **type_** | **string** | Show only objects of specified type field. | 

### Return type

[**GetAllValidations200Response**](GetAllValidations200Response.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchById

> Batch GetBatchById(ctx, id).Execute()

Retrieves a batch object by the {id} path parameter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetBatchById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBatchById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBatchById`: Batch
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBatchById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Batch**](Batch.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvidenceById

> Evidence GetEvidenceById(ctx, id).Execute()

Retrieves an evidence object by the {id} path parameter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEvidenceById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEvidenceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvidenceById`: Evidence
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEvidenceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvidenceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Evidence**](Evidence.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteById

> Quote GetQuoteById(ctx, id).Execute()

Retrieves a quote object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetQuoteById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetQuoteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuoteById`: Quote
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetQuoteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Quote**](Quote.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateByIPAddress

> Rate GetRateByIPAddress(ctx).IpAddress(ipAddress).Execute()

Retrieves a geolocated VAT rate by IP address.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    ipAddress := "ipAddress_example" // string | IPv4 or IPv6 address to geolocate. If none is provided, the IP address of the request is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRateByIPAddress(context.Background()).IpAddress(ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRateByIPAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateByIPAddress`: Rate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRateByIPAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRateByIPAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipAddress** | **string** | IPv4 or IPv6 address to geolocate. If none is provided, the IP address of the request is used. | 

### Return type

[**Rate**](Rate.md)

### Authorization

[secretKey](../README.md#secretKey), [publicKey](../README.md#publicKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupplyById

> Supply GetSupplyById(ctx, id).Execute()

Retrieves a supply object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSupplyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSupplyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupplyById`: Supply
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSupplyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Supply**](Supply.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidationById

> Validation GetValidationById(ctx, id).Execute()

Retrieves a validation object.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetValidationById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetValidationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidationById`: Validation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetValidationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Validation**](Validation.md)

### Authorization

[secretKey](../README.md#secretKey), [publicKey](../README.md#publicKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBatchById

> Batch UpdateBatchById(ctx, id).Mode(mode).Name(name).Queries(queries).Status(status).Execute()

Updates a batch object by the {id} path parameter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    mode := "mode_example" // string | Either `append` to append an array of queries to the existing array of queries, or `replace` to empty the existing array of queries prior to adding. Defaults to `append`. (optional)
    name := "name_example" // string | Descriptive name of the batch object. Give it any name you like to serve for your reference. (optional)
    queries := []string{"Inner_example"} // []string | Array of VAT numbers to be added while considering the `mode`. (optional)
    status := "status_example" // string | Set it to `scheduled` once you want to schedule the validation process with the queries provided. It’s possible to add queries and change the status in one go. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateBatchById(context.Background(), id).Mode(mode).Name(name).Queries(queries).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBatchById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBatchById`: Batch
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBatchById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBatchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | Either &#x60;append&#x60; to append an array of queries to the existing array of queries, or &#x60;replace&#x60; to empty the existing array of queries prior to adding. Defaults to &#x60;append&#x60;. | 
 **name** | **string** | Descriptive name of the batch object. Give it any name you like to serve for your reference. | 
 **queries** | **[]string** | Array of VAT numbers to be added while considering the &#x60;mode&#x60;. | 
 **status** | **string** | Set it to &#x60;scheduled&#x60; once you want to schedule the validation process with the queries provided. It’s possible to add queries and change the status in one go. | 

### Return type

[**Batch**](Batch.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEvidenceById

> Evidence UpdateEvidenceById(ctx, id).BankAddress(bankAddress).BillingAddress(billingAddress).IpAddress(ipAddress).Execute()

Updates an evidence object by the {id} path parameter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    bankAddress := TODO // CreateEvidenceRequestBankAddress |  (optional)
    billingAddress := TODO // CreateEvidenceRequestBillingAddress |  (optional)
    ipAddress := TODO // CreateEvidenceRequestIpAddress |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEvidenceById(context.Background(), id).BankAddress(bankAddress).BillingAddress(billingAddress).IpAddress(ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEvidenceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEvidenceById`: Evidence
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEvidenceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEvidenceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bankAddress** | [**CreateEvidenceRequestBankAddress**](CreateEvidenceRequestBankAddress.md) |  | 
 **billingAddress** | [**CreateEvidenceRequestBillingAddress**](CreateEvidenceRequestBillingAddress.md) |  | 
 **ipAddress** | [**CreateEvidenceRequestIpAddress**](CreateEvidenceRequestIpAddress.md) |  | 

### Return type

[**Evidence**](Evidence.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupplyById

> Supply UpdateSupplyById(ctx, id).Amount(amount).AmountRefunded(amountRefunded).CountryCode(countryCode).Currency(currency).Description(description).Evidence(evidence).InvoiceNumber(invoiceNumber).Issued(issued).Name(name).Notes(notes).Validation(validation).Vat(vat).Execute()

Updates a supply object by the {id} path parameter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/vatstack-client-go"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    amount := int32(56) // int32 | Amount in **cents** (e.g. 100.50 must be expressed as 10050). This common common workaround prevents unexpected rounding issues. (optional)
    amountRefunded := int32(56) // int32 | Amount in cents refunded back to the customer. (optional)
    countryCode := "countryCode_example" // string | 2-letter ISO country code of the place of supply that is relevant for the `vat.rate`. (optional)
    currency := "currency_example" // string | 3-letter ISO 4217 currency code used to charge the `amount`. The currency is used to correctly convert to your reporting currency. (optional)
    description := "description_example" // string | A custom string to describe the supplied item. (optional)
    evidence := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an [evidence object](https://vatstack.com/docs/evidences). The pieces of non-contradictory evidence contained therein will affect the `evidence_status`. (optional)
    invoiceNumber := "invoiceNumber_example" // string | A custom string for the invoice number issued to the customer. It’s advisable to follow sequential numbering. (optional)
    issued := "issued_example" // string | ISO date at which the invoice was issued to the customer. (optional)
    name := "name_example" // string | A custom string for the name of the customer. (optional)
    notes := "notes_example" // string | A custom string for additional notes. (optional)
    validation := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of a [validation object](https://vatstack.com/docs/validations). This is useful if you let your customer enter a VAT number beforehand. Its valid value can affect `vat.amount`, `vat.rate` and `amount_total` when zero-rating. (optional)
    vat := TODO // CreateSupplyRequestVat |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSupplyById(context.Background(), id).Amount(amount).AmountRefunded(amountRefunded).CountryCode(countryCode).Currency(currency).Description(description).Evidence(evidence).InvoiceNumber(invoiceNumber).Issued(issued).Name(name).Notes(notes).Validation(validation).Vat(vat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSupplyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSupplyById`: Supply
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSupplyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupplyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **int32** | Amount in **cents** (e.g. 100.50 must be expressed as 10050). This common common workaround prevents unexpected rounding issues. | 
 **amountRefunded** | **int32** | Amount in cents refunded back to the customer. | 
 **countryCode** | **string** | 2-letter ISO country code of the place of supply that is relevant for the &#x60;vat.rate&#x60;. | 
 **currency** | **string** | 3-letter ISO 4217 currency code used to charge the &#x60;amount&#x60;. The currency is used to correctly convert to your reporting currency. | 
 **description** | **string** | A custom string to describe the supplied item. | 
 **evidence** | **string** | Unique identifier of an [evidence object](https://vatstack.com/docs/evidences). The pieces of non-contradictory evidence contained therein will affect the &#x60;evidence_status&#x60;. | 
 **invoiceNumber** | **string** | A custom string for the invoice number issued to the customer. It’s advisable to follow sequential numbering. | 
 **issued** | **string** | ISO date at which the invoice was issued to the customer. | 
 **name** | **string** | A custom string for the name of the customer. | 
 **notes** | **string** | A custom string for additional notes. | 
 **validation** | **string** | Unique identifier of a [validation object](https://vatstack.com/docs/validations). This is useful if you let your customer enter a VAT number beforehand. Its valid value can affect &#x60;vat.amount&#x60;, &#x60;vat.rate&#x60; and &#x60;amount_total&#x60; when zero-rating. | 
 **vat** | [**CreateSupplyRequestVat**](CreateSupplyRequestVat.md) |  | 

### Return type

[**Supply**](Supply.md)

### Authorization

[secretKey](../README.md#secretKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

