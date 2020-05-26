[1mdiff --git a/pkg/adaptor/infrastructure/repository/repository_suite_test.go b/pkg/adaptor/infrastructure/repository/repository_suite_test.go[m
[1mindex 19b78ae3..a9d91615 100644[m
[1m--- a/pkg/adaptor/infrastructure/repository/repository_suite_test.go[m
[1m+++ b/pkg/adaptor/infrastructure/repository/repository_suite_test.go[m
[36m@@ -6,7 +6,6 @@[m [mimport ([m
 	"net/http/httptest"[m
 	"testing"[m
 [m
[31m-	"github.com/aws/aws-sdk-go/aws/awserr"[m
 	"github.com/aws/aws-sdk-go/aws/session"[m
 	"github.com/aws/aws-sdk-go/service/s3"[m
 	_ "github.com/golang-migrate/migrate/v4/database/mysql"[m
[36m@@ -77,27 +76,12 @@[m [mfunc truncate(db *gorm.DB) {[m
 	Expect(db.Exec("SET FOREIGN_KEY_CHECKS=1").Error).To(Succeed())[m
 }[m
 [m
[31m-func prepareBucket(sess *session.Session, bucket string) error {[m
[32m+[m[32mfunc clear(sess *session.Session, bucket string) error {[m
 	s3c := s3.New(sess)[m
 [m
[31m-	_, err := s3c.CreateBucket(&s3.CreateBucketInput{[m
[31m-		Bucket: &bucket,[m
[31m-	})[m
[31m-[m
[31m-	if err == nil {[m
[31m-		return nil[m
[31m-	}[m
[31m-[m
[31m-	awsErr, ok := err.(awserr.Error)[m
[31m-	if !(ok && awsErr.Code() == s3.ErrCodeBucketAlreadyExists) {[m
[31m-		return err[m
[31m-	}[m
[31m-[m
[31m-	// Bucketが既に存在している場合[m
[31m-[m
 	var errDelete error[m
 	listInput := &s3.ListObjectsV2Input{Bucket: &bucket}[m
[31m-	err = s3c.ListObjectsV2Pages(listInput, func(page *s3.ListObjectsV2Output, lastPage bool) bool {[m
[32m+[m	[32merr := s3c.ListObjectsV2Pages(listInput, func(page *s3.ListObjectsV2Output, lastPage bool) bool {[m
 		for _, obj := range page.Contents {[m
 			_, err := s3c.DeleteObject(&s3.DeleteObjectInput{[m
 				Bucket: &bucket,[m
[1mdiff --git a/pkg/adaptor/infrastructure/repository/user_test.go b/pkg/adaptor/infrastructure/repository/user_test.go[m
[1mindex fe03bdf6..968f264c 100644[m
[1m--- a/pkg/adaptor/infrastructure/repository/user_test.go[m
[1m+++ b/pkg/adaptor/infrastructure/repository/user_test.go[m
[36m@@ -29,7 +29,7 @@[m [mvar _ = Describe("UserRepositoryImpl", func() {[m
 [m
 		truncate(db)[m
 [m
[31m-		Expect(prepareBucket(tests.AWS, tests.Config.AWS.FilesBucket)).To(Succeed())[m
[32m+[m		[32mExpect(clearBucket(tests.AWS, tests.Config.AWS.FilesBucket)).To(Succeed())[m
 	})[m
 [m
 	DescribeTable("Storeは引数のuserを作成するか、その状態になるように更新する",[m
