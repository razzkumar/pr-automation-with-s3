# GitHub Action to Deploy static website to S3 Bucket and comment the Url to PR☁

This action uses the golang aws sdk to build aws static website bucket and
deploy the static file to that newly created s3 bucket and comment the url to
the PR. To deploy static file it uses either from your repository or build
during your workflow. Dockerfile and s3 utils can be found on
[Link](https://github.com/razzkumar/pr-automation-s3-utils).


## Usage

### `deploy-pr.yml` and `cleanup-pr.yml` Example

Place in a `.yml` file such as this one in your `.github/workflows` folder. [Refer to the documentation on workflow YAML syntax here.](https://help.github.com/en/articles/workflow-syntax-for-github-actions)
#### The following example includes optimal defaults for a any client side javascript static website:

- It build the javascript/typescript frontend application with the help of given
  command (ex: `BUILD_COMMAND="yarn build"`)
- If If it't PR then it will comment the URL of the static site to the PR
- It also delete the Build s3 bucket on PR merge or close

```yaml
name: Deploy site to S3 And add comment to PR

on:
  pull_request:
    branches:
    - master

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - uses: razz
      with:
        args: --acl public-read --follow-symlinks --delete
      env:
        AWS_S3_BUCKET: ${{ secrets.AWS_S3_BUCKET }}
        AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
        AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        AWS_REGION: 'us-west-1'   # optional: defaults to us-east-1
        SOURCE_DIR: 'public'      # optional: defaults to entire repository
```


