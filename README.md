# faasflow-minio-datastore
A **[faasflow](https://github.com/s8sg/faasflow)** datastore implementation that uses minio DB to store data  
which can also be used with s3 bucket

## Minio Data-Store Configuration
Minio data-store needs the below configuration in the `flow.yml`
```yaml
s3_url
s3_region
s3_tls
s3_secret_key_name
s3_access_key_name
```

## Use Minio dataStore in `faasflow`
* Set the `stack.yml` with the necessary environments
```yaml
      s3_url: "minio.faasflow:9000"
      s3_tls: false
      s3_secret_key_name: "s3-secret-key"
      s3_access_key_name: "s3-access-key"
    secrets:
      - s3-secret-key
      - s3-access-key
```
* Use the `faasflowMinioDataStore` as a DataStore on `handler.go`
```go
minioDataStore "github.com/s8sg/faas-flow-minio-datastore"

func DefineDataStore() (faasflow.DataStore, error) {

       // initialize minio DataStore
       miniods, err := minioDataStore.InitFromEnv()
       if err != nil {
               return nil, err
       }

       return miniods, nil
}
```
