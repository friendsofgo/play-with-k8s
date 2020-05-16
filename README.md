# Play with Kubernetes
This repository was created to demostrate how to deploy a simple application using Kubernetes in first step and
in second step deploy a Kubernetes application using [Helm](https://helm.sh/).

We will use the [minikube tool](https://kubernetes.io/docs/tasks/tools/install-minikube/) to deploy our application.

## Build the image

First of all we will need an application to deploy, for that reason we've created a simple application on Go, using
[Giphy API](https://developers.giphy.com/) to print a random gif in our browser.

*You would need your own Giphy API key to use the application properly*

To build the image execute:

```
$ docker build friendsofgo/giphyneitor .
```

And you could run it:

```
$ docker run -p 8080:8080 -e GIPHY_API_KEY=your_api_key friendsofgo/giphyneitor
```

And then you can navigate to `http://localhost:8080` and see something similar to:

![giphyneitor](https://res.cloudinary.com/fogo/image/upload/c_scale,w_500/v1589566547/fogo/blog/giphyneitor_mood.png)

The image is hosted on docker hub: [https://hub.docker.com/r/friendsofgo/giphyneitor](https://hub.docker.com/r/friendsofgo/giphyneitor)

## Install the application on Kubernetes

Firs of all you will need a Kubernete cluster, we recommend to use [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) to execute on local, but you could use
any other valid tool or even an cloud cluster.

To run our application we will need the files inside of `deploy/k8s`, we have a file for each Kubernetes resource:

### Configmap

As we saw in the docker image section, we will need a environment variable to use our application, for that reason we will create
a [configmap](https://kubernetes.io/docs/concepts/configuration/configmap/) to maintain our configuration part isolate from application code.

The `configmap is in `deploy/k8s/configmap.yaml` and you can install on your cluster using the next command:

```
$ kubectl apply -f deploy/k8s/configmap.yaml
```

### Service

As you know the pods are mortal; when they die never resurrected again, for that you can use the deployment to ensure that
when a pod die, then another will be create, but each pod has is own ip address so in some cases for example a frontend that needs a backend
will need a way to find this IP each time.

For that reason exists the [services](https://kubernetes.io/docs/concepts/services-networking/service/), a Service is an abstraction which defines a logical set of Pods and a policy by which to access them.

Our `service` file is in `deploy/k8s/service` and you can install on your cluster using:

```
$ kubectl apply -f deploy/k8s/service.yaml
```

### Deployment

Until now, we've saw how to create our configuration, how to create our service to communicate our pods in a same way, but 
we need a way to put in common all this resources.

In a few word a [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) is where you descibre the desire state of your Kubernete application, and
the Deployment controller changes the actual state to the desired state at a controlled rate.

You can find our `deployment` on `deploy/k8s/deployment.yaml` and install it on your cluster using:

```
$ kubectl apply -f deploy/k8s/deployment.yaml
```

## Check that all it's working

Now if we can if all our resources are created using `kubectl get`.

```
$ kubectl get pods

NAME                           READY   STATUS    RESTARTS   AGE
giphyneitor-6cd9bd877d-ns6ps   1/1     Running   0          13h
```

```
$ kubectl get deployments

NAME          READY   UP-TO-DATE   AVAILABLE   AGE
giphyneitor   1/1     1            1           13h
```

```
$ kubectl get services

NAME          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
giphyneitor   ClusterIP   10.101.111.215   <none>        8080/TCP   13h
```

### Connect with our pod

If you want mapping your pod with your local server, to see if the application is running well
we can use the command `kubectl port-forward {pod_name} {ports}`.

```
$ kubectl port-forward giphyneitor-6cd9bd877d-ns6ps 8080:8080

Frwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
```

And that is, now lunch your favorite browser on `http://localhost:8080` and enjoy with your gif!

## Helm

Coming soon!