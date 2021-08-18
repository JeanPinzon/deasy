# <<.Name>>

A REST API scaffolded using Deasy


## How to run

> It requires Docker

```
make run
```


## How to Deploy

Github CI do it for us.

You need to set some secrets on your Github repository settings:

| Secret key            | Secret value description                                                      |
|-----------------------|-------------------------------------------------------------------------------|
| DOCKERHUB_USERNAME    | Your docker hub user name                                                     |
| DOCKERHUB_PASSWORD    | Your docker hub password                                                      |
| AWS_ACCESS_KEY_ID     | Your AWS access id                                                            |
| AWS_SECRET_ACCESS_KEY | Your AWS access secret                                                        |
| KUBE_CONFIG_DATA      | A base 64 of your `$HOME/.kube/config`. Try `cat $HOME/.kube/config | base64` |

> [Take a look here](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html) how to manage AWS access keys
