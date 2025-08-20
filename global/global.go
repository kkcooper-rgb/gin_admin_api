package global

import (
	"context"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
	Ctx = context.Background()
)
