FROM ubuntu:bionic as base

# Install curl.
RUN apt-get update \
  && apt-get install -y curl git bash-completion vim nano \
  && echo 'source /etc/bash_completion' >> /etc/bash.bashrc \
  && rm -rf /var/lib/apt/lists/*

# Install kubectl.
ENV K8S_RELEASE=v1.16.10 K8S_SHA512=455d8bd8881aa996dd35969b2448fa54ab15650ce2d58c19bf89ff8e11f287beb136ba64ca62d64c958a1227dfee1e94408941258f0b9421ef06bb48a09bbcd2

RUN curl -L -o k8s.tgz https://dl.k8s.io/${K8S_RELEASE}/kubernetes-client-linux-amd64.tar.gz \
  && echo -n "${K8S_SHA512} k8s.tgz" | sha512sum -c --quiet \
  && tar xzf k8s.tgz \
  && mv kubernetes/client/bin/kubectl /usr/local/bin/kubectl \
  && rm -r kubernetes/ k8s.tgz \
  && echo 'source <(kubectl completion bash)' >> /etc/bash.bashrc

# Install helm.
ENV HELM_RELEASE=v3.1.2 HELM_SHA256=e6be589df85076108c33e12e60cfb85dcd82c5d756a6f6ebc8de0ee505c9fd4c

RUN curl -L -o helm.tgz https://get.helm.sh/helm-${HELM_RELEASE}-linux-amd64.tar.gz \
  && echo -n "${HELM_SHA256} helm.tgz" | sha256sum -c --quiet \
  && tar xzf helm.tgz \
  && mv linux-amd64/helm /usr/local/bin/helm \
  && rm -r linux-amd64/ helm.tgz \
  && echo 'source <(helm completion bash)' >> /etc/bash.bashrc \
  && helm repo add stable https://kubernetes-charts.storage.googleapis.com/

COPY motd /etc/motd
RUN cat /etc/motd >> /etc/bash.bashrc

# Make repository available.
RUN git clone https://github.com/3fs/k8s-workshop-episode-2.git /repo

FROM base as console

# Insert k8s.3fs.si Secret.
COPY k8s.3fs.si-cert.yaml /k8s.3fs.si-cert.yaml

# Set up kubectl config. and workshop env variable.
ARG CLUSTER_ENDPOINT
ARG USER_TOKEN
ARG USER_NAMESPACE
ENV CLUSTER_ENDPOINT=${CLUSTER_ENDPOINT} USER_TOKEN=${USER_TOKEN} USER_NAMESPACE=${USER_NAMESPACE} CODE=${USER_NAMESPACE}

COPY ca.crt /ca.crt

RUN kubectl config set-cluster default \
    --embed-certs=true \
    --server=${CLUSTER_ENDPOINT} \
    --certificate-authority=/ca.crt \
  && kubectl config set-credentials default --token=${USER_TOKEN} \
  && kubectl config set-context default-context \
    --cluster=default \
    --user=default \
    --namespace=${USER_NAMESPACE} \
  && kubectl config use-context default-context
