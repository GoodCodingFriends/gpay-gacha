package gcs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GoodCodingFriends/gpay-gacha/source"
	"github.com/morikuni/failure"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type gcsSource struct {
	c           *storage.Client
	bucketNames []string
}

func New(ctx context.Context, bucketNames []string) (source.Source, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	for _, name := range bucketNames {
		if v := os.Getenv(fmt.Sprintf("GCS_BUCKET_%s", strings.ToUpper(name))); v == "" {
			return nil, failure.Wrap(errors.New("GCS env missing"), failure.Context{"bucket": name})
		}
	}

	return &gcsSource{c: client, bucketNames: bucketNames}, nil
}

func (s *gcsSource) Random(ctx context.Context) (io.ReadCloser, error) {
	bktName := s.bucketNames[rand.Int31n(len(s.bucketNames))]
	bkt := s.c.Bucket(bktName)
	n, err := strconv.Atoi(os.Getenv(fmt.Sprintf("GCS_BUCKET_%s", strings.ToUpper(bktName))))
	if err != nil {
		return nil, failure.Wrap(err)
	}
	bkt.Object()
}

func (s *gcsSource) cacheObjects(ctx context.Context) error {

}
