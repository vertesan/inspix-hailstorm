#!/bin/bash
set -euo pipefail
REPO_NAME=repo
REPO_URI='git@github.com:vertesan/privtest.git'
SSH_KEY_PATH='/app/secrets/key'
UPDATE_FLAG='cache/updated'

if [ ! -f $SSH_KEY_PATH ]; then
  echo "git ssh key file does not exists, will be stopping process..."
  exit -1
fi

./hailstorm --dbonly

if [ -f $UPDATE_FLAG ]; then 
  git config --global user.name vts-server
  git config --global user.email vts-server@e.mail
  git config --global core.sshCommand "ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i $SSH_KEY_PATH -F /dev/null"
  git clone $REPO_URI $REPO_NAME

  cp masterdata/*.yaml $REPO_NAME/
  git -C $REPO_NAME add .
  git -C $REPO_NAME commit -m 'update'
  git -C $REPO_NAME push 
fi

echo "All processes completed."
