# [Workshop](../README.md) &raquo; [Helm](./README.md) &raquo; Inspect and Upgrade Release

## Task objective

Every release has its own release history. For every release with the same
name, helm creates a new release version and deploys it. If the deployment
failes, rollback to the previous version is performed.

We can inspect the previously created releases by using `helm history` command.

```console
# helm history my-deploy
REVISION	UPDATED                 	STATUS  	CHART                 	APP VERSION	DESCRIPTION
1       	Tue May 26 22:39:49 2020	deployed	workshop-example-0.1.0	1.16.0     	Install complete
```

---

In previous tasks, we have deployed `workshop-example` chart with the default
values. If we would like to update the configuration of the release, we can
upgrade the release and specify the updated values by any of:

* creating a new file with configuration overrides:
  * copy default `values.yaml` file to `values-overrides.yaml`
  * update values in `values-overrides.yaml`
  * update the release by executing

    ```bash
    helm upgrade my-deploy workshop-example --values values-overrides.yaml
    ```

In this task we will learn how to adapt the configuration values and deploy the updated resources to kubernetes cluster.

## Task

Upgrade and deploy the release with all modifications, we have created in previous hands-on parts:

* update annotations and labels
  * add `workshop` and `code` annotations
* update environment variables
  * add `CODE: $CODE` and `WORKSHOP_USERNAME: k8s-workshop` environment variables
* update health checks
  * add readiness check to `/ready`
  * add liveness check to `/live`
* update resource limits
  * configure requested resources to `cpu: 10m` and `memory: 32Mi`
  * configure limits to `cpu: 100m` and `memory: 128Mi`
* enable `ConfigMap`
* configure Secret
  * `password`: `your-random-password`
* enable PersistentVolume

Additionally, configure `Ingress` to expose the application to outside world.
k8s.3fs.si`.

<details>
    <summary>Show solution</summary>

Copy `values.yaml` file to `values-overrides.yaml` and edit the file with

```diff
--- values.yaml	2020-05-26 23:00:36.000000000 +0200
+++ values-overrides.yaml	2020-05-26 23:03:36.000000000 +0200
@@ -15,20 +15,20 @@ imagePullSecrets: []
 nameOverride: ""
 fullnameOverride: ""

-environment: {}
-  # CODE: cranky-hippo
-  # WORKSHOP_USERNAME: k8s-workshop
+environment:
+  CODE: cranky-hippo
+  WORKSHOP_USERNAME: k8s-workshop

-secret: {}
-  # password: k8s-workshop-password
+secret:
+  password: k8s-workshop-password

 healthchecks:
   livenessPath: '/live'
   readinessPath: '/ready'

-podAnnotations: {}
-  # workshop: k8s-workshop-episode-2
-  # code: cranky-hippo
+podAnnotations:
+  workshop: k8s-workshop-episode-2
+  code: cranky-hippo

 podSecurityContext: {}
   # fsGroup: 2000
@@ -46,21 +46,21 @@ service:
   port: 80

 ingress:
-  enabled: false
+  enabled: true
   annotations:
     kubernetes.io/ingress.class: nginx
   hosts:
-    - host: $CODE.k8s.3fs.si
+    - host: cranky-hippo.k8s.3fs.si
       paths:
         - /
   tls:
     - secretName: k8s.3fs.si-certificate
       hosts:
-        - $CODE.k8s.3fs.si
+        - cranky-hippo.k8s.3fs.si

 persistence:
   ## If persistence is enabled, uploaded files will be persisted.
-  enabled: false
+  enabled: true
   ## The path the volume will be mounted at.
   mountPath: /uploadfiles

@@ -85,7 +85,7 @@ resources:
     memory: 64Mi
   requests:
     cpu: 10m
-    memory: 64Mi
+    memory: 32Mi

 nodeSelector: {}
```

Upgrade the release by executing `helm upgrade my-deploy workshop-example --values values-overrides.yaml`

```console
# helm upgrade my-deploy . -f values-overrides.yaml
Release "my-deploy" has been upgraded. Happy Helming!
NAME: my-deploy
LAST DEPLOYED: Tue May 26 23:03:02 2020
NAMESPACE: default
STATUS: deployed
REVISION: 2
NOTES:
1. Get the application URL by running these commands:
  https://cranky-hippo.k8s.3fs.si/
```

Update `values-overrides.yaml` file is available [here](./solution/values-overrides.yaml)

</details>

## Inspect the release

After a successful deployment we can check history once again.

```console
# helm history my-deploy
REVISION	UPDATED                 	STATUS    	CHART                 	APP VERSION	DESCRIPTION
1       	Tue May 26 22:39:49 2020	superseded	workshop-example-0.1.0	1.16.0     	Install complete
2       	Tue May 26 23:03:02 2020	deployed  	workshop-example-0.1.0	1.16.0     	Upgrade complete
```

There is a way to check which user provided values were changed by executing the following command:

```console
# helm get values my-deploy
USER-SUPPLIED VALUES:
affinity: {}
environment:
  CODE: cranky-hippo
  WORKSHOP_USERNAME: k8s-workshop
fullnameOverride: ""
healthchecks:
  livenessPath: /live
  readinessPath: /ready
...
```

## Next: [Rollback release](./03_rollback_release.md)
