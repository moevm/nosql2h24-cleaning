package main

import (
	"context"
	"fmt"
	"hello_world/pkg/logger"
	"log/slog"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type World struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
}

func main() {
	log := logger.SetupLogger("local")

	// hardcoded mongodb URI
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Error("mongo-driver connection failed", slog.Any("err", err))
		os.Exit(1)
	}

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Error("mongo-driver disconnect failed:", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Error("mongo-driver ping failed", slog.Any("err", err))
	}

	worlds := client.Database("hello_world").Collection("worlds")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertRes, err := worlds.InsertMany(ctx, []interface{}{
		World{
			Name: "Willy Wonka's World",
			Description: `Willy Wonka is shown as a kind,
			quirky man who wants to invite many different kids from different backgrounds to his chocolate factory.`,
		},
		World{
			Name: "Ringworld",
			Description: `1970’s Ringworld is a load of fun for people enamored with technology and what the future of humanity might include in terms of tech.
			While it doesn’t include even the Internet as we know it, it features human beings with near-infinite lifespans, x-ray lasers, inertialess drive,
			variable swords (like a light saber, but better), super materials, and most importantly, a world in the shape of a ring,
			that produces scientifically realistic artificial gravity by spinning.`,
		},
		World{
			Name: "Foundation World",
			Description: `In 12,067 G.E. ("Galactic Era"), mathematician and psychologist Hari Seldon has developed psychohistory,
			a new field of science and psychology that allows for the probabilistic prediction of future events. By means of psychohistory,
			Seldon has discovered the decline and eventual fall of the Galactic Empire, angering its rulers, the Commission of Public Safety.
			Seldon defends his beliefs, and the Commission, not wanting to make Seldon a martyr, offers him exile to a remote world, Terminus.`,
		},
	})
	if err != nil {
		log.Error("worlds.InsertMany failed", slog.Any("err", err))
	}
	log.Info(fmt.Sprintf("worlds.InsertMany %v", insertRes.InsertedIDs))

	id := insertRes.InsertedIDs[2].(bson.ObjectID)
	log.Info(fmt.Sprintf("deleting document with _id: %s", id))

	delRes, err := worlds.DeleteOne(
		ctx,
		bson.D{
			{Key: "_id", Value: id},
		},
	)
	if err != nil {
		log.Error("worlds.DeleteOne failed", slog.Any("err", err))
	}

	log.Info(fmt.Sprintf("deleted %v documents", delRes.DeletedCount))

	cursor, err := worlds.Find(ctx, bson.M{})
	if err != nil {
		log.Error("worlds.Find failed", slog.Any("err", err))
	}

	w := []World{}
	err = cursor.All(ctx, &w)
	if err != nil {
		log.Error("cursor.All failed", slog.Any("err", err))
	}

	log.Info("readed documents:")
	for _, world := range w {
		log.Info(fmt.Sprintf("%v", world))
	}
}
