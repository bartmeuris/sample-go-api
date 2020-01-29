# Helm deploy tool

## Goal

* Allow self-servicing of development environments by developers:
  * Select which microservice versions are deployed to which environment
  * Allow manipulating parameters per microservice for development environments
* Allow easy deployments to production and preproduction

## Data model

### Roles

Different roles exist:

* guest (default)
* developer
* deployer
* admin

### Cluster

Combination of:

* OpenShift/K8s Cluster
* Credentials
* Default namespace
* role

### Project

A project defines which microservices can be deployed to this project.

* List of microservices

### Microservice

A definition of a microservice

* name
* Git URL
* (Chart name?)

### Environment

An environment is linked to a project, and defines:

* which version of which microservice is running in this environment
* the role of the environment (dev/qa/preprod/prod)
* Cluster to deploy to
* Namespace to deploy to (if empty -> see cluster)
