#!/bin/bash

#!/bin/sh
set -e
set -o pipefail

WORKING_DIRECTORY="$PWD"

# Set some defaults

[ -z "$GITHUB_PAGES_BRANCH" ]   && GITHUB_PAGES_BRANCH=gh-pages
[ -z "$HELM_CHARTS_SOURCE" ]    && HELM_CHARTS_SOURCE="$WORKING_DIRECTORY/repos"
[ -z "$HELM_VERSION" ]          && HELM_VERSION=3.2.0
[ -z "$KUBEVAL_VERSION" ]       && KUBEVAL_VERSION=0.15.0
[ -z "$KUBERNETES_VERSION" ]    && KUBERNETES_VERSION=1.18.0
[ -z "$GITHUB_PAGES_REPO" ]     && GITHUB_PAGES_REPO=funkypenguins-geek-cookbook/charts-dev

echo "GITHUB_PAGES_REPO=$GITHUB_PAGES_REPO"
echo "GITHUB_PAGES_BRANCH=$GITHUB_PAGES_BRANCH"
echo "HELM_CHARTS_SOURCE=$HELM_CHARTS_SOURCE"
echo "HELM_VERSION=$HELM_VERSION"
echo "KUBERNETES_VERSION=$KUBERNETES_VERSION"
echo "KUBEVAL_VERSION=$KUBEVAL_VERSION"
echo "PATH=$PATH"

echo '>> Prepare...'
# mkdir -p /tmp/helm/bin
# mkdir -p /tmp/helm/publish
# mkdir -p /tmp/kubeval/bin
# mkdir -p /tmp/kubeval/manifests

# echo '>> Installing kubeval...'
# wget https://github.com/garethr/kubeval/releases/download/${KUBEVAL_VERSION}/kubeval-linux-amd64.tar.gz 
# tar xzvf kubeval-linux-amd64.tar.gz
# chmod u+x kubeval
# mv kubeval /usr/local/bin


echo ">> Checking out $GITHUB_PAGES_BRANCH branch from $GITHUB_PAGES_REPO"
rm -rf to_publish
git clone -b "$GITHUB_PAGES_BRANCH" "https://github.com/$GITHUB_PAGES_REPO.git" to_publish

echo '>> Building charts...'
find "$HELM_CHARTS_SOURCE" -mindepth 4 -maxdepth 4 -type d | grep charts/ | while read chart; do
  echo ">>> helm lint $chart"
  helm lint "$chart"

#   echo ">>> kubeval $chart"
#   /root/project/.circleci/prep-kubeval.sh
#   mkdir -p "/tmp/kubeval/manifests/$chart_name"
#   helm dep update $chart
#   helm template $chart --output-dir "/tmp/kubeval/manifests/$chart_name"
#   kubeval -d "/tmp/kubeval/manifests/$chart_name"
 
#   #echo ">>> unittest $chart"
#   #/root/project/.circleci/prep-unit-tests.sh  
#   #helm unittest $chart 

#   chart_name="`basename "$chart"`"
#   echo ">>> helm package -d $chart_name $chart"
#   mkdir -p "$chart_name"
#   helm package -d "$chart_name" "$chart"
done
# echo '>>> helm repo index'
# helm repo index .

# if [ "$CIRCLE_BRANCH" != "master" ]; then
#   echo "Current branch is not master and do not publish"
#   exit 0
# fi

# echo ">> Publishing to $GITHUB_PAGES_BRANCH branch of $GITHUB_PAGES_REPO"
# git config user.email "$CIRCLE_USERNAME@users.noreply.github.com"
# git config user.name "Circle CI"
# git add .
# git status
# git commit -m "Published by Circle CI $CIRCLE_BUILD_URL"
# git push origin "$GITHUB_PAGES_BRANCH"