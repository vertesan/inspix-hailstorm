#!/bin/bash
set -euo pipefail
REPO_NAME=repo
SSH_KEY_PATH='/app/secrets/key'
UPDATE_FLAG='cache/updated'

if [ ! -f $SSH_KEY_PATH ]; then
  echo "git ssh key file does not exists, will be stopping process..."
  exit 166
fi

if [ ! -v REPO_NAME ]; then
  echo "REPO_NAME is not set, will be stopping process..."
  exit 167
fi

./hailstorm --dbonly

if [ -f $UPDATE_FLAG ]; then 
  git config --global user.name vts-server
  git config --global user.email 169537433+vts-server@users.noreply.github.com
  git config --global core.sshCommand "ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i $SSH_KEY_PATH -F /dev/null"
  echo "Cloning from remote repository..."
  git clone $REPO_URI $REPO_NAME

  cp masterdata/*.yaml $REPO_NAME/
  git -C $REPO_NAME add .
  cur_ver=`cat cache/currentVersion.txt`
  git -C $REPO_NAME commit -m "$cur_ver"
  echo "Pushing to remote repository..."
  git -C $REPO_NAME push 
fi

echo "All processes completed."
