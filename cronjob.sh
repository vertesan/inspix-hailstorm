#!/bin/bash

# optional ENVs:
# INS_DO_PUSH: push master database files to a git repository after decrypting if there is a new version
# INS_DIFF_REPO_URI: git repository URI, takes no effect if INS_DO_PUSH is not set
# INS_SSH_KEY_PATH: ssh key used to push files to INS_DIFF_REPO_URI, takes no effect if INS_DO_PUSH is not set
# INS_REPO_USER_NAME: user name used for push action, takes no effect if INS_DO_PUSH is not set
# INS_REPO_USER_EMAIL: user email used for push action, takes no effect if INS_DO_PUSH is not set

set -euo pipefail
cd "$(dirname "$0")"

REPO_NAME=diff_repo
UPDATE_FLAG='cache/updated'

. .env.local
logfile="cache/`date '+%Y-%m-%d'`.log"
exec > >(tee -a "$logfile") 2>&1

echo "======================== cronjob run at $(date '+%Y-%m-%d %H:%M:%S') ==========================="

if [ ! -v REPO_NAME ]; then
  echo "REPO_NAME is not set, will be stopping process..."
  exit 167
fi

./hailstorm --dbonly

if [ ! -v INS_DO_PUSH ]; then
  echo "All processes completed."
  exit 0
fi

if [ ! -f "$INS_SSH_KEY_PATH" ]; then
  echo "git ssh key file does not exists, will be stopping process..."
  exit 166
fi

if [ -f "$UPDATE_FLAG" ]; then 
  if [ ! -d "$REPO_NAME" ]; then
    echo "Cloning from remote repository..."
    GIT_SSH_COMMAND="ssh -o IdentitiesOnly=yes -i ${INS_SSH_KEY_PATH}" git clone $INS_DIFF_REPO_URI $REPO_NAME
  fi

  git -C $REPO_NAME config user.name $INS_REPO_USER_NAME
  git -C $REPO_NAME config user.email $INS_REPO_USER_EMAIL
  git -C $REPO_NAME config core.sshCommand "ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i $INS_SSH_KEY_PATH -F /dev/null"
  git -C $REPO_NAME pull

  cp masterdata/*.yaml $REPO_NAME/
  git -C $REPO_NAME add .
  cur_ver=`cat cache/currentVersion.txt`
  git -C $REPO_NAME commit -m "$cur_ver"
  echo "Pushing to remote repository..."
  git -C $REPO_NAME push 
fi

echo "All processes completed."
