package utils

// import (
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/credentials"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/service/s3"
//     "fmt"
// )

// func NewS3Client(environment *Environment) *s3.S3 {

//     sess, err := session.NewSession(&aws.Config{
//         Region:      aws.String(environment.AwsRegion),
//         Credentials: credentials.NewStaticCredentials(environment.AwsAccessKey, environment.AwsSecretKey, ""),
//     })
//     if err != nil {
//         fmt.Println("Cannot initiate S3 session")
//         panic(err)
//     }

//     // Create and return the S3 service client
//     return s3.New(sess)
// }

