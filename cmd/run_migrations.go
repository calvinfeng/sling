package cmd

import (
	"fmt"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	
	_ "github.com/lib/pq" // Driver
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zpl0310/database/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	host         = "localhost"
	port         = "5432"
	user         = "pli"
	password     = "pheebia"
	database     = "sling"
	ssl          = "sslmode=disable"
	migrationDir = "file://./migrations/"
)

var log = logrus.WithFields(logrus.Fields{
	"pkg": "cmd",
})

var pgAddr = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, database, ssl)

// RunMigrationsCmd is a command to run migration.
var RunMigrationsCmd = &cobra.Command{
	Use:   "runmigrations",
	Short: "run migration on database",
	RunE:  runMigrations,
}

func createUserList(db *gorm.DB, users []*model.User) error {
	for i, u := range users {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return err
		}

		u.PasswordDigest = hashBytes
		u.JWTToken = fmt.Sprintf("admin-%d", i)
		if err := db.Create(u).Error; err != nil {
			return err
		}
		logrus.Infof("user %v is created", u.ID)
	}
	return nil
}

func createRoomList(db *gorm.DB, rooms []*model.Room) error {
	for _, r := range rooms {
		if err := db.Create(r).Error; err != nil {
			return err
		}
		logrus.Infof("room %v is created", r.ID)
	}
	return nil
}

func createMessageList(db *gorm.DB, messages []*model.Message) error {
	for _, m := range messages {
		if err := db.Create(m).Error; err != nil {
			return err
		}
		logrus.Infof("message %v is created", m.ID)
	}
	return nil
}

func createUserroomList(db *gorm.DB, userrooms []*model.Usersrooms) error {
	for _, ur := range userrooms {
		if err := db.Create(ur).Error; err != nil {
			return err
		}
		logrus.Infof("usersrooms %v is created", ur.UserID)
	}
	return nil
}

func runMigrations(cmd *cobra.Command, args []string) error {

	migration, err := migrate.New(migrationDir, pgAddr)
	if err != nil {
		return err
	}

	log.Info("performing reset on database")
	if err = migration.Drop(); err != nil {
		return err
	}

	if err := migration.Up(); err != nil {
		return err
	}

	log.Info("migration has been performed successfully")

	db, err := gorm.Open("postgres", pgAddr)
	if err != nil {
		return err
	}

	createUserList(db, []*model.User{
		&model.User{
			Name:     "Alice",
			Email:    "alice@goacademy.com",
			Password: "alice",
		},
		&model.User{
			Name:     "Bob",
			Email:    "bob@goacademy.com",
			Password: "bob",
		},
		&model.User{
			Name:     "Calvin",
			Email:    "calvin@goacademy.com",
			Password: "calvin",
		},
		&model.User{
			Name:     "Ria",
			Email:    "ria@goacademy.com",
			Password: "ria",
		},
	})

	createRoomList(db, []*model.Room{
		&model.Room{
			RoomName: "Room1",
			Type:     0,
		},
		&model.Room{
			RoomName: "Bob~Alice",
			Type:     1,
		},
		&model.Room{
			RoomName: "Room2",
			Type:     0,
		},
	})

	createMessageList(db, []*model.Message{
		&model.Message{
			CreateTime: time.Now(),
			Body:       "Hi~~",
			SenderID:   1,
			RoomID:     1,
		},
		&model.Message{
			CreateTime: time.Now(),
			Body:       "You clearly don’t know who you’re talking to, 
			so let me clue you in. I am not in danger, Skyler. I am the danger. 
			A guy opens his door and gets shot, and you think that of me? No! I am the one who knocks!",
			SenderID:   2,
			RoomID:     1,
		},
		&model.Message{
			CreateTime: time.Now(),
			Body:       "Hi~~",
			SenderID:   3,
			RoomID:     1,
		},
		&model.Message{
			CreateTime: time.Now(),
			Body:       "Hi~~",
			SenderID:   1,
			RoomID:     2,
		},
		&model.Message{
			CreateTime: time.Now(),
			Body:       "Hi~~",
			SenderID:   2,
			RoomID:     2,
		},
		&model.Message{
			CreateTime: time.Now(),
			Body:       "Hi~~",
			SenderID:   1,
			RoomID:     3,
		},
		&model.Message{
			CreateTime: time.Now(),
			Body:       "Hi~~",
			SenderID:   3,
			RoomID:     3,
		},
		&model.Message{
			CreateTime: time.Now(),
			Body:       "Hi~~",
			SenderID:   4,
			RoomID:     3,
		},
	})

	createUserroomList(db, []*model.Usersrooms{
		&model.Usersrooms{
			UserID:    1,
			RoomID:    1,
			HasUnread: true,
		},
		&model.Usersrooms{
			UserID:    2,
			RoomID:    1,
			HasUnread: false,
		},
		&model.Usersrooms{
			UserID:    3,
			RoomID:    1,
			HasUnread: true,
		},
		&model.Usersrooms{
			UserID:    3,
			RoomID:    3,
			HasUnread: false,
		},
		&model.Usersrooms{
			UserID:    1,
			RoomID:    3,
			HasUnread: false,
		},
		&model.Usersrooms{
			UserID:    4,
			RoomID:    3,
			HasUnread: true,
		},
		&model.Usersrooms{
			UserID:    2,
			RoomID:    2,
			HasUnread: false,
		},
		&model.Usersrooms{
			UserID:    1,
			RoomID:    2,
			HasUnread: true,
		},
	})
	return nil
}
