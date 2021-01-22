# Messages

This is a very simple web application that allows you to post messages and show
them on an index page. This is the classic
[guestbook](https://github.com/kubernetes/examples/tree/master/guestbook)
application, but way simpler and created so it could be showcased during
trainings. In particular, this repository also holds Kubernetes manifests to
train people with Kubernetes, and a Helm chart to do the same with Helm. This
application uses [redis](https://redis.io/) as the backend storage.

## Building and running

This application has been written in [go](https://golang.org/), so in order to
build this application you need to perform:

    $ go build

That being said, in order to run this application we recommend using Docker. You
can build the Docker image like this:

    $ docker build -t mssola/messages:0.1 .

I have already pushed some tags on the
[DockerHub](https://cloud.docker.com/u/mssola/repository/docker/mssola/messages). In
order to run this application you need to take into account some environment
variables:

- `MESSAGES_FILE_PATH`: the base path were the index page is located. You
  usually don't have to set this one, because it only matters if you are
  planning to have the index page on another path (which is something that
  happens on the Docker image but that's already handled for you).
- `MESSAGES_REDIS_HOST`: the host were redis is located (default: `localhost`).
- `MESSAGES_REDIS_PORT`: the port were redis is bound (default: `6379`).
- `MESSAGES_REDIS_PASSWORD`: the password to be used when connecting to the
  redis instance.

## Kubernetes and Helm

You can find more information about how to deploy this application with plain
Kubernetes manifests or with Helm in the [kubernetes](./kubernetes) directory.

## License

```
Copyright (C) 2019-2021 Miquel Sabaté Solà <mikisabate@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```
