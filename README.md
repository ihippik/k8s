# K8S

K8S deployment examples.

The application can be launched in two modes of operation:

* Job
* Server (API)

### Kubectl

| Command                                | Description                                               |
|----------------------------------------|-----------------------------------------------------------|
| alias k=kubectl                        | set alias                                                 |
| k delete all --all                     | delete all resource in default ns                         |
| k apply -f entity.yml                  | create/update resource from YAML file                     |
| k apply -f ./dir                       | create resources from all manifest files in the directory |
| k delete `entity_type` `entity_name`   | delete resource                                           |
| k describe `entity_type` `entity_name` | get a detailed description of the resource                |