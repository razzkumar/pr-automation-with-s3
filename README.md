# Static Site Automation

### What it is?
This Gihub action that uses the [golang aws sdk](https://aws.amazon.com/sdk-for-go/) to build s3 bucket and attach policy
for static website deploy the static file to that newly created s3 bucket and
comment the url to the PR. To deploy static file it uses either from your
repository or build during your workflow. There is [self hosted](https://github.com/razzkumar/frontend-PR-automation) tool, if Github
action is not feasible.


### Best for?
   - Immediate feedback visually to developers or anyone interested in changes.
   - Reduce burden of having to build application for QA and verify the changes.
   - Faster iterations.

### How to use?

Add `.yml` file/s such as given examples in your `.github/workflows` folder. [Refer to the documentation on workflow YAML syntax here.](https://help.github.com/en/articles/workflow-syntax-for-github-actions)

##### The following example will:
   - Create s3 bucket and attach policy for static site
   - Build the javascript/typescript frontend application with the help of
     given command (ex: `BUILD_COMMAND="yarn build"`)
   - Upload build file (static site) to s3
   - Comment the URL of the static site to the Pull Request
   - Delete the aws S3 bucket after PR is merged

##### Config file: `.github/workflows/deploy-existing.yml`

```yaml
name: Next js frontend dev

on:
  push:
    branches:
    - dev

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Build and deploy Studio app
      uses: razzkumar/pr-automation-with-s3@v1.0.2
      env:
        AWS_S3_BUCKET: ${{ secrets.AWS_S3_BUCKET }} 
        AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
        AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        AWS_REGION: "us-east-1"
        SRC_FOLDER: "out"
        ACTION: 'deploy'
        BUILD_COMMAND: "yarn build && yarn export"
        CLOUDFRONT_ID: ${{ secrets.CLOUDFRONT_ID }}
        SECRETS_MANAGER: ${{ secrets.SECRETS_MANAGER }} // name of secrets on secret manager
```


##### Config file: `.github/workflows/deploy-on-pr.yml`

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
      uses: razzkumar/pr-automation-with-s3@v1.0.2
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


##### Config file: `.github/workflows/cleanup-on-pr-merge.yml`

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
      uses: razzkumar/pr-automation-with-s3@v1.0.2
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

| Key                     | Suggested Type | Value                                                                                                                                                                                                                       | Required                                                                  | Default                                                                                                                  |
| -------------           | -------------  | -------------                                                                                                                                                                                                               | -------------                                                             | -------------                                                                                                            |
| `GH_ACCESS_TOKEN`       | `secrect env`  | Your Github access token used while commenting PR                                                                                                                                                                           | **YES/NO** If `ACTION: create` then it's required,otherwise it's optional | NA                                                                                                                       |
| `AWS_ACCESS_KEY_ID`     | `secret env`   | Your AWS Access Key. [More info here.](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html)                                                                                                         | **Yes**                                                                   | N/A                                                                                                                      |
| `AWS_SECRET_ACCESS_KEY` | `secret env`   | Your AWS Secret Access Key. [More info here.](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html)                                                                                                  | **Yes**                                                                   | N/A                                                                                                                      |
| `AWS_S3_BUCKET`         | `secret env`   | The name of the bucket you're syncing to. For example, `jarv.is` or `my-app-releases`.                                                                                                                                      | **YES/NO**                                                                | - If running on PR it will genereat by tool `PR-Branch`.pr`PR-number`.auto-deploy - In the case of depoyment it required |
| `AWS_REGION`            | `env`          | The region where you created your bucket. Set to `us-east-2` by default. [Full list of regions here.](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions) | No                                                                        | `us-east-2`                                                                                                              |
| `SRC_FOLDER`            | `env`          | The local directory (or file) you wish to deploy to S3. For example, `public`. Defaults to `build`.                                                                                                                         | No                                                                        | `build` (based on react app)                                                                                             |
| `IS_BUILD`              | `env`          | This is the flag that indicate that build a project or not                                                                                                                                                                  | No                                                                        | `true` (It will run `yarn && yarn build` by default)                                                                     |
| `ACTION`                | `env`          | This is also a flag that indicate what to do (`create`:-create s3 (if not exist) bucket,build react and comment on PR,`deploy`:helps to deploy to s3,`delete`: delete the s3 bucket)                                        | No                                                                        | `create` (It will create s3 (if not exist),built the app, deploy to s3 and comment URL to PR`)                           |
| `BUILD_COMMAND`         | `env`          | How to build the react app if its `npm run build` then it will run `npm install && npm run build`                                                                                                                           | No                                                                        | `yarn build` (It will run `yarn && yarn build` by default)                                                               |
| `CLOUDFRONT_ID`         | `secret env`   | id of cloudfront for invalidation                                                                                                                                                                                           | No                                                                        |                                                                                                                          |
| `SECRETS_MANAGER`       | `env`          | name of the aws secres manager key                                                                                                                                                                                          | No                                                                        |                                                                                                                          |


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

[![HitCount](http://hits.dwyl.com/razzkumar/pr-automation-with-s3.svg)](http://hits.dwyl.com/razzkumar/pr-automation-with-s3)
