# K8S

K8S deployment examples.

#### Description
The application can be launched in two modes of operation:

* Job `./app job -sleep 10`
* HTTP Server `./app server`

The `server` is a task that never ends and a `job`, on the contrary, 
ends after a time specified by the flag `sleep`.

You can see a description of this by simply running the binary.


#### Let's prepare the Docker image
* `docker login`
* `docker build -t ihippik/k8s-scum:main .`
* `docker push ihippik/k8s-scum:main`

### Kubectl

| Command                                | Description                                               |
|----------------------------------------|-----------------------------------------------------------|
| alias k=kubectl                        | set alias                                                 |
| k delete all --all                     | delete all resource in default ns                         |
| k apply -f entity.yml                  | create/update resource from YAML file                     |
| k apply -f ./dir                       | create resources from all manifest files in the directory |
| k delete `entity_type` `entity_name`   | delete resource                                           |
| k describe `entity_type` `entity_name` | get a detailed description of the resource                |