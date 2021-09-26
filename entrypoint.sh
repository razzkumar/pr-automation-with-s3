#!/bin/sh

set -eu

if [ -z "$AWS_ACCESS_KEY_ID" ]; then
  echo "AWS_ACCESS_KEY_ID is not set. Quitting."
  exit 1
fi

if [ -z "$AWS_SECRET_ACCESS_KEY" ]; then
  echo "AWS_SECRET_ACCESS_KEY is not set. Quitting."
  exit 1
fi

# Check checkoperation

if [ -z "$ACTION" ]; then
  ACTION="create"
fi

#listing all available eviroment variables

case "$ACTION" in
# Create action will create new s3 static site and  deploy on it
create)
    if [[ "$GITHUB_EVENT_NAME" == "pull_request" ]];then
     # check GH_ACCESS_TOKEN is set or not for the commit_
      if [ -z "$GH_ACCESS_TOKEN" ]; then
        echo "GH_ACCESS_TOKEN is not set. Quitting."
        exit 1
      fi

     # Running prcomment command to commit
      /s3 -action $ACTION

      ## Fail action if /s3 command throw error
      if [ $? -ne 0 ];then
        echo "::error::Failed to deploy to the s3"
        exit 1
      fi

    else
      echo "::error::Unable to build "
      exit 1
    fi
  ;;
  #Simpy delet the static site from s3
  delete)
    /s3 -action $ACTION

    if [ $? -ne 0 ]; then
      echo "::error::Unable to Unable to delete "
      exit 1
    fi
  ;;
  deploy)
    if [ -z "$AWS_S3_BUCKET" ]; then
      echo "$AWS_S3_BUCKET is not set. Quitting."
      exit 1
    fi

    /s3 -action $ACTION

    if [ $? -ne 0 ]; then
      echo "::error::Unable to Build static file "
      exit 1
    fi
  ;;
  *)
  echo "::error:: Can't perform any task"
  exit 1
  ;;
esac
