# Helm chart for `messages`

This is the Helm chart for the `messages` application.

## TL;DR

```bash
$ helm install .
```

## Prerequisites

- Kubernetes 1.8+

## Installing the Chart

To install the chart with the release name `messages`:

```bash
$ helm install --name messages .
```

The command deploys this application on the Kubernetes cluster in the default
configuration. The [configuration](#configuration) section lists the parameters
that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `messages` deployment:

```bash
$ helm delete messages
```

The command removes all the Kubernetes components associated with the chart and
deletes the release. You can also completely purge this release with:

```bash
$ helm delete --purge messages
```

## Configuration

The following table lists the configurable parameters of the `messages` chart
and their default values.

| Parameter                         | Description                                  | Default              |
|-----------------------------------|----------------------------------------------|----------------------|
| `replicaCount`                    | Number of replicas for this application      | `1`                  |
| `image.repository`                | `messages` Image name                        | `mssola/messages`    |
| `image.tag`                       | `messages` Image tag                         | `{VERSION}`          |
| `image.pullPolicy`                | The pull policy for the image                | `Always`             |
| `service.type`                    | The type of service to be exposed            | `NodePort`           |
| `service.port`                    | The port where this service will be bound    | `80`                 |
| `tolerations`                     | Kubernetes tolerations for the created pods  | `[]`                 |
| `affinity`                        | Kubernetes affinity for the created pods     | `{}`                 |
| `persistence`                     | The host path where the data will be located | `/opt/messages-data` |
| `redis.image.tag`                 | The tag to be used for redis                 | `5.0.3`              |
| `redis.persistence.existingClaim` | The PVC to be used for redis                 | `redis-storage`      |

### Persistence

As you could guess from the `persistence` configuration, we are using `hostPath`
to write the data. Of course this is not recommended for production
environments, but it's simple enough for trainings, as this application is
intended for. Therefore, there is no reason to touch the `persistence`
configuration.

Last but not least, as you can see we are overwriting the default value from
redis. Again, don't touch this, since we instruct redis to use the PVC created
by this Helm chart.
