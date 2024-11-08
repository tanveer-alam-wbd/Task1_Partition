package servics

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"kafkaexperiment/models"
// 	"kafkaexperiment/utils"
// 	"time"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	"github.com/docker/docker/integration-cli/environment"
// )

// type dataAccess interface {
// 	WriteMessages(message []utils.models) error
// }

// type s3DataAccess struct {
// 	s3client *s3.S3
// 	bucket string
// }

// func NewS3dataAccess(environment *utils.Environment) *s3DataAccess {
// 	s3client := utils.NewS3Client
// }