package embed

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"
)

// Purge purge data before time
func Purge(ctx context.Context, time time.Time) error {
	wid, err := selectWidBeforeTime(time)
	if err != nil {
		return err
	}

	for _, id := range wid {
		name := fmt.Sprintf("%016X", id)
		path := filepath.Join(Directory, index, name[:4], name[8:12], name)
		if err = os.Remove(path); err != nil {
			if !os.IsNotExist(err) {
				log.Info().Msgf("purging %v has err: %v", path, err)
			}
		}
	}

	return purgeBeforeTime(ctx, time)
}
