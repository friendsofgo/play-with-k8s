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

![giphyneitor](https://res.cloudinary.com/fogo/image/upload/c_scale,w_500/v1589561912/giphyneitor_run.png)