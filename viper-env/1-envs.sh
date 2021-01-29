#!/bin/sh

export CI=vela
export VELA=true
export VELA_BUILD_EVENT=deployment
export VELA_BUILD_TARGET=dev
export VELA_DEPLOYMENT=dev
export VELA_BUILD_REF=heads/branches/main
export APPLICATION=batchconsumer

