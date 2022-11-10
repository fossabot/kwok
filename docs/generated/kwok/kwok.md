## kwok

kwok is a tool for simulate thousands of fake kubelets

### Synopsis

kwok is a tool for simulate thousands of fake kubelets

```
kwok [command] [flags]
```

### Options

```
      --cidr string                                        CIDR of the pod ip (default "10.0.0.1/24")
      --disregard-status-with-annotation-selector string   All node/pod status excluding the ones that match the annotation selector will be watched and managed.
      --disregard-status-with-label-selector string        All node/pod status excluding the ones that match the label selector will be watched and managed.
  -h, --help                                               help for kwok
      --kubeconfig string                                  Path to the kubeconfig file to use
      --manage-all-nodes                                   All nodes will be watched and managed. It's conflicted with manage-nodes-with-annotation-selector and manage-nodes-with-label-selector.
      --manage-nodes-with-annotation-selector string       Nodes that match the annotation selector will be watched and managed. It's conflicted with manage-all-nodes.
      --manage-nodes-with-label-selector string            Nodes that match the label selector will be watched and managed. It's conflicted with manage-all-nodes.
      --master string                                      Server is the address of the kubernetes cluster
      --node-ip ip                                         IP of the node (default 196.168.0.1)
      --server-address string                              Address to expose health and metrics on
```

