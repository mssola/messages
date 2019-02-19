## The `manifests` directory

This directory contains manifests to be used during trainings. In particular,
there is an introductory [pod.yml](./manifests/pod.yml) file that allows
trainees to get acquainted with the most basic block of Kubernetes.

Once they are done, they will jump towards the [full.yml](./manifests/full.yml)
file, which contains everything to get a working deployment of the `messages`
application. One can perform the following commands:

```bash
$ kubectl create -f manifests/full.yml
$ export NODE_IP=$(kubectl get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")
$ export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services messages)
$ echo http://$NODE_IP:$NODE_PORT
# Paste the URL you got from the above command into your web browser.
```

Once you are done playing with this deployment, you can simply execute:

    $ kubectl delete -f manifests/full.yml

## The `messages` directory

This directory contains the Helm chart. The instructions on how to use are
located [here](./messages/README.md).
