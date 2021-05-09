package main
 
import (
        "log"
        "github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
		"github.com/aws/aws-sdk-go/aws"
		// "github.com/aws/aws-sdk-go/aws/awserr"
		"bytes"


)
//  GOOS=linux go build -o main goStringManipulation.go; zip main.zip main; c:Users/danny/go/bin/build-lambda-zip.exe -output main.zip main
var invokeCount = 0
// var myObjects []*s3.Object
var myObjects *s3.Object
func init() {

		// svc := s3.New(session.New())
		// input := &s3.ListBucketsInput{}

		// result, err := svc.ListBuckets(input)
		// if err != nil {
		// 	if aerr, ok := err.(awserr.Error); ok {
		// 		switch aerr.Code() {
		// 		default:
		// 			log.Print(aerr.Error())
		// 		}
		// 	} else {
		// 		// Print the error, cast err to awserr.Error to get the Code and
		// 		// Message from an error.
		// 		log.Print(err.Error())
		// 	}
		// 	return
		// }
		// log.Print(result)
        // log.Print("starting new session")
        // svc := s3.New(session.New())
        // log.Print("svc made ")
        // // input := &s3.ListObjectsV2Input{
        // //         Bucket: aws.String("gostorage"),
		// // }
		// input := &s3.GetObjectInput{
		// 	Bucket: aws.String("gostorage"),
		// 	Key: aws.String("hello_world.txt"),
		// }
		// log.Print("input is made",input)
		// // result, _ := svc.ListObjectsV2(input)
		// result, _ := svc.GetObject(input)
		// log.Print("result",result)
        // myObjects = result
}
 
// function, which takes a string as
// argument and return the reverse of string.
func reverse(s string) string {
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    // return the reversed string.
    return string(rns)
}

func LambdaHandler() (string, error) {
        invokeCount = invokeCount + 1
		// log.Print(result)
        log.Print("starting new session")
        svc := s3.New(session.New())
        log.Print("svc made ")
        // input := &s3.ListObjectsV2Input{
        //         Bucket: aws.String("gostorage"),
		// }
		input := &s3.GetObjectInput{
			Bucket: aws.String("gostorage"),
			Key: aws.String("hello_world.txt"),
		}
		log.Print("input is made",input)
		// result, _ := svc.ListObjectsV2(input)
		result, _ := svc.GetObject(input)

		// log.print(string(result.Body))

		buf := new(bytes.Buffer)
		buf.ReadFrom(result.Body)
		oldStr := buf.String()
		log.Print("old string is ",oldStr)
		//flip the string over 
		newStr := reverse(oldStr)
		log.Print("new string is ",newStr)
		// buf := new(bytes.Buffer)
		// buf.ReadFrom(result)
		// myFileContentAsString := buf.String()

		return newStr, nil
}
 
func main() {
        log.Print("starting")
        lambda.Start(LambdaHandler)
}