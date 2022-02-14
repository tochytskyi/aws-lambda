# aws-lambda
Convert JPEG images from S3 bucket to BMP, GIF, PNG into another bucket

### Setup two buckets
- `jpeg-images` for source jpeg images
- `converted-jpeg-images` for converted jpeg images

### Setup lambda
- Create Go 1.x lambda function with S3 trigger (`jpeg-images`)
- Build go executable
```shell
cd ./function
GOOS=linux go build -o converter-function main.go
```
- Make zip from `converter-function` executable and upload to AWS lambda
- Setup Runtime settings to invoke go executable
  ![img_5.png](docs/img_5.png)
- Add env variables to use S3 credentials for you IAM role: `AccessKeyID`, `AccessKeySecret`

  ![img_4.png](docs/img_4.png)

Result lambda scheme:

![img_2.png](docs/img_2.png)

### Test
Upload jpg image to `jpeg-images` bucket

`jpeg-images`

![img.png](docs/img.png)

`converted-jpeg-images` after couple of seconds

![img_1.png](docs/img_1.png)

and check CloudWatch logs

![img_3.png](docs/img_3.png)