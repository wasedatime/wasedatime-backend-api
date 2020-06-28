#!/bin/bash
set -x

if [ $TRAVIS_BRANCH == 'staging' ] ; then
    git checkout staging
    git remote add deploy 'deploy@139.59.216.161:wasedatime-backend.git'
    git add .
    git status
    git commit -m "Deploy from Travis CI"
    git status
    git push deploy staging --force
else
    echo "Not deploying, since this branch isn't staging."
fi
