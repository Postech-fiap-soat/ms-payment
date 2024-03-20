#!/bin/bash

kubectl apply -f soatdb-configmap.yml

kubectl apply -f soatdb.yml

kubectl apply -f soatdb-svc.yml

kubectl apply -f soatapp-configmap.yml

kubectl apply -f secrets.yml

kubectl apply -f soatapp.yml

kubectl apply -f soatapp-svc.yml