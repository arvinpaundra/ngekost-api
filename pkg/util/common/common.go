package common

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed)
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func GracefulShutdown(ctx context.Context, timeout time.Duration, ops func(context.Context) error) <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		s := make(chan os.Signal, 1)

		defer close(wait)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		defer close(s)

		log.Println("shutting down application")

		// set timeout for ops to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit\n", timeout)
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		err := ops(ctx)
		if err != nil {
			log.Printf("cleaning up failed: %s", err.Error())
		}
	}()

	return wait
}
