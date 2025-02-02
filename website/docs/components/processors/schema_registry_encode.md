---
title: schema_registry_encode
type: processor
status: experimental
categories: ["Parsing","Integration"]
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the contents of:
     lib/processor/schema_registry_encode.go
-->

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

:::caution EXPERIMENTAL
This component is experimental and therefore subject to change or removal outside of major version releases.
:::
Automatically encodes and validates messages with schemas from a Confluent Schema Registry service.

Introduced in version 3.58.0.


<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

<TabItem value="common">

```yaml
# Common config fields, showing default values
label: ""
schema_registry_encode:
  url: ""
  subject: ""
```

</TabItem>
<TabItem value="advanced">

```yaml
# All config fields, showing default values
label: ""
schema_registry_encode:
  url: ""
  subject: ""
  tls:
    skip_cert_verify: false
    enable_renegotiation: false
    root_cas: ""
    root_cas_file: ""
    client_certs: []
```

</TabItem>
</Tabs>

Encodes messages automatically from schemas obtains from a [Confluent Schema Registry service](https://docs.confluent.io/platform/current/schema-registry/index.html) by polling the service for the latest schema version for target subjects.

If a message fails to encode under the schema then it will remain unchanged and the error can be caught using error handling methods outlined [here](/docs/configuration/error_handling).

Currently only Avro schemas are supported.

## Fields

### `url`

The base URL of the schema registry service.


Type: `string`  

### `subject`

The schema subject to derive schemas from.
This field supports [interpolation functions](/docs/configuration/interpolation#bloblang-queries).


Type: `string`  

```yaml
# Examples

subject: foo

subject: ${! meta("kafka_topic") }
```

### `tls`

Custom TLS settings can be used to override system defaults.


Type: `object`  

### `tls.skip_cert_verify`

Whether to skip server side certificate verification.


Type: `bool`  
Default: `false`  

### `tls.enable_renegotiation`

Whether to allow the remote server to repeatedly request renegotiation. Enable this option if you're seeing the error message `local error: tls: no renegotiation`.


Type: `bool`  
Default: `false`  
Requires version 3.45.0 or newer  

### `tls.root_cas`

An optional root certificate authority to use. This is a string, representing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.


Type: `string`  
Default: `""`  

```yaml
# Examples

root_cas: |-
  -----BEGIN CERTIFICATE-----
  ...
  -----END CERTIFICATE-----
```

### `tls.root_cas_file`

An optional path of a root certificate authority file to use. This is a file, often with a .pem extension, containing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.


Type: `string`  
Default: `""`  

```yaml
# Examples

root_cas_file: ./root_cas.pem
```

### `tls.client_certs`

A list of client certificates to use. For each certificate either the fields `cert` and `key`, or `cert_file` and `key_file` should be specified, but not both.


Type: `array`  

```yaml
# Examples

client_certs:
  - cert: foo
    key: bar

client_certs:
  - cert_file: ./example.pem
    key_file: ./example.key
```

### `tls.client_certs[].cert`

A plain text certificate to use.


Type: `string`  
Default: `""`  

### `tls.client_certs[].key`

A plain text certificate key to use.


Type: `string`  
Default: `""`  

### `tls.client_certs[].cert_file`

The path to a certificate to use.


Type: `string`  
Default: `""`  

### `tls.client_certs[].key_file`

The path of a certificate key to use.


Type: `string`  
Default: `""`  


