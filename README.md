


## GitHub Action to Deploy static website to S3 Bucket and comment the Url to PR☁

This action uses the golang aws sdk to build aws static website bucket and
deploy the static file to that newly created s3 bucket and comment the url to
the PR. To deploy static file it uses either from your repository or build
during your workflow. There is [self
hosed](https://github.com/razzkumar/frontend-PR-automation) tool, if Github
action is not feasible.


## Usage

### `deploy-on-pr.yml` and `cleanup-on-pr-merge.yml` Example

Place in a `.yml` file such as this one in your `.github/workflows` folder. [Refer to the documentation on workflow YAML syntax here.](https://help.github.com/en/articles/workflow-syntax-for-github-actions)
#### The following example includes optimal defaults for a any client side javascript static website:

- It build the javascript/typescript frontend application with the help of given
  command (ex: `BUILD_COMMAND="yarn build"`)
- If If it't PR then it will comment the URL of the static site to the PR
- It also delete the Build s3 bucket on PR merge or close


#### depoy-on-pr.yml
  - What is does:
    - Create s3 bucket and attach policy for static site
    - Build the application
    - Upload build file (static file) to s3
    - Comment the URL of the static site to the Pull Request

```yaml
name: Deploy site to S3 And add comment to PR and delete after merge

on:
  pull_request:
    branches:
    - master

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Static site deploy to s3 and comment on PR
      uses: razzkumar/pr-automation-with-s3@v1.0.0
      env:
        AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
        AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        GH_ACCESS_TOKEN: ${{ secrets.GH_ACCESS_TOKEN}}
        AWS_REGION: 'us-east-2'     # optional: defaults to us-east-2
        SRC_FOLDER: 'build'         # optional: defaults to build (react app)
        IS_BUILD: 'true'            # optional: defaults to true
        ACTION: "create"            # optional: defaults to create (option:create,delete and deploy)
        BUILD_COMMAND: "yarn build" # optional: defaults to `yarn build`
```


#### cleanup-on-pr-merge.yml
  - It should be inclued with the `deploy-on-pr.yml`
    - What it does:
      - Delete the aws S3 bucket after PR is merged

```yaml
name: Delete S3 bucket after PR merge

on:
  pull_request:
    types: [closed]

jobs:
  delete:
    runs-on: ubuntu-latest
    steps:
    - name: Clean up temperory bucket
      if: github.event.pull_request.merged == true
      uses: razzkumar/pr-automation-with-s3@v1.0.0
      env:
        AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
        AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        AWS_REGION: 'us-east-2'     # optional: defaults to us-east-2
        ACTION: "delete"            # Action must be delete to delete

```

- It can be used for many other purpose also,some of the [examples](https://github.com/razzkumar/pr-automation-with-s3/blob/master/examples/deploy-pre-build-site.yml) are:
  - Deploy prebuild app [config link](https://github.com/razzkumar/pr-automation-with-s3/blob/master/examples)
  - Build react app and deploy it. [config link](https://github.com/razzkumar/pr-automation-with-s3/blob/master/examples/build-and-deploy-react-app.yml)

### Configuration

The following settings must be passed as environment variables as shown in the example. Sensitive information, especially `GH_ACCESS_TOKEN`,`AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`, should be [set as encrypted secrets](https://help.github.com/en/articles/virtual-environments-for-github-actions#creating-and-using-secrets-encrypted-variables) — otherwise, they'll be public to anyone browsing your repository's source code and CI logs.

| Key | Value | Suggested Type | Required | Default |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `GH_ACCESS_TOKEN` | Your Github access token used while commenting PR | `secrect env` | **YES/NO** If `ACTION: create` then it's required,otherwise if optional | NA  |
| `AWS_ACCESS_KEY_ID` | Your AWS Access Key. [More info here.](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html) | `secret env` | **Yes** | N/A |
| `AWS_SECRET_ACCESS_KEY` | Your AWS Secret Access Key. [More info here.](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html) | `secret env` | **Yes** | N/A |
| `AWS_S3_BUCKET` | The name of the bucket you're syncing to. For example, `jarv.is` or `my-app-releases`. | `secret env` | **YES/NO** | - If running on PR it will genereat by tool `PR-Branch`.pr`PR-number`.auto-deploy - In the case of depoyment it required |
| `AWS_REGION` | The region where you created your bucket. Set to `us-east-2` by default. [Full list of regions here.](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions) | `env` | No | `us-east-2` |
| `SRC_FOLDER` | The local directory (or file) you wish to deploy to S3. For example, `public`. Defaults to `build`. | `env` | No | `build` (based on react app) |
| `IS_BUILD` | This is the flag that indicate that build a project or not | `env` | No | `true` (It will run `yarn && yarn build` by default) |
| `ACTION` | This is also a flag that indicate what to do (`create`:-create s3 (if not exist) bucket,build react and comment on PR,`deploy`:helps to deploy to s3,`delete`: delete the s3 bucket) | `env` | No | `create` (It will create s3 (if not exist),built the app, deploy to s3 and comment URL to PR`) |
| `BUILD_COMMAND` | How to build the react app if its `npm run build` then it will run `npm install && npm run build` | `env` | No | `yarn build` (It will run `yarn && yarn build` by default) |


#### Note for S3 Bucket creation
 - It only create a s3 bucket if not `exist`
 - While Creating bucket for the pull_request S3 bucket name will be: `PR-Branch`.pr`PR-number`.auto-deploy
    - For Eg.:
      - if  base branch is `SIG-1000` and PR number is `23` the the bucket name will be `sig-100.pr23.auto-deploy`
 - If we deploy app on push or (not on pull requst) like prebuild app deployment, app build and deploy then the bucket name will be `$AWS_S3_BUCKET.auto-deploy`
    - For Eg.
      - if `AWS_S3_BUCKET=dev-test-deployment` then bucket will be `dev-test-deployment.auto-deploy`

## TODO
 - [ ] Add tests
 - [ ] Add option to deploy on aws cloudfront
 - [ ] Design PR comment done by tool
 - [ ] Maintain code quality

## Contributing
  Feel free to send pull requests

## License
This project is distributed under the [MIT license](LICENSE.md)
