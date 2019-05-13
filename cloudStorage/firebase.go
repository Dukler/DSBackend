package cloudStorage

import (
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"path/filepath"
)


type store struct{
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle
	ctx context.Context
}

var (
	Firebase store
)


func (f *store) InitBucket (bucketName string){
	f.bucketName = bucketName
	s := fmt.Sprintf("%s.appspot.com",bucketName)
	config := &firebase.Config{
		StorageBucket: s,
	}
	path,_:=filepath.Abs("./tokens/firebase.json")
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	f.bucket = bucket
}
func (f *store)ReadFile (file string) []byte{
	rc, err:= f.bucket.Object(file).NewReader(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatalln(err)

	}
	return slurp
}
