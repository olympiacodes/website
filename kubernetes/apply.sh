#!/bin/sh

NAMESACE=bellinghamcodes

kubectl apply -f namespace.yaml
kubectl --namespace=${NAMESACE} apply -f ./secrets.yaml
kubectl --namespace=${NAMESACE} apply -f ./service.yaml
kubectl --namespace=${NAMESACE} apply -f ./deployment.yaml
