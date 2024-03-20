#!/bin/bash

#kind create cluster --config=kind.yml --name=soatcluster

kubectl apply -f mongo-configmap.yml
kubectl apply -f mongo-deploy.yml
kubectl apply -f mongo.yml

kubectl apply -f rabbitmq-configmap.yml
kubectl apply -f rabbitmq-deploy.yml
kubectl apply -f rabbitmq.yml

kubectl apply -f soatdb-configmap.yml
kubectl apply -f soatdb-deploy.yml
kubectl apply -f soatdb.yml

kubectl apply -f soatmspayment-configmap.yml
kubectl apply -f soatmspayment-deploy.yml
kubectl apply -f soatmspayment.yml

kubectl apply -f soatmsorder-deploy.yml
kubectl apply -f soatmsorder.yml
