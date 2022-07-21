# Pie-Dictionary

Pie-Dictionary is a simple, lightweight in-memory caching micreoservice which uses Redis as backend & has 
inbuilt Trie algorithm built-in for fast key search based on prefix/postfix & is written in GoLang

**Note**: this is not a production ready application and we do not recommend its in production

## Background details
* This application uses Redis core to store all the relevant data
* It can run anywhere, but we have provided support for only Kubernetes in this repo

## API details
* Application exposes 3 APIs -
  * GET  - `/get/<key>` -> gets value of key specified
  * GET  - `/search?<prefix>/<postfix>=value` -> gets gets all keys having provided prefix/postfix
  * POST - `/set` -> adds new key-value to data-store
    * Required Body while making this API call: 
    ```
    {
        "key": "key",
        "value": "value"
    }
    ```
      
## Prerequisites:
* GoLang 1.16+ required
* Kubernetes -having LoadBalancer support (latest supported versions recommended)

**NOTE:** application requires LoadBalancer k8s svc

## Try it yourself
* kubectl apply -f /home/niranjan/Desktop/lummo/k8s-configs/ -n nir
* get public IP of LB svc deployed
* Have fun!!! Start accessing above mentioned APIs