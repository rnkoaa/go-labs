#!/bin/sh

export CI=vela
export VELA=true
export VELA_BUILD_EVENT=deployment
export VELA_BUILD_TARGET=stage
export VELA_DEPLOYMENT=stage
export VELA_BUILD_REF=heads/branches/main
export VELA_DESCRIPTION=dry=true,cluster=batchconsumer-stage
export APPLICATION=batchconsumer

