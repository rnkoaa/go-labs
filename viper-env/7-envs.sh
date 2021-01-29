#!/bin/sh

export CI=vela
export VELA=true
export VELA_BUILD_EVENT=deployment
export VELA_BUILD_TARGET=stage
export VELA_DEPLOYMENT=stage
export VELA_BUILD_REF=heads/tags/641
export VELA_DESCRIPTION=dry=true
export APPLICATION=batchconsumer

