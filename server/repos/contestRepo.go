package repos

import (
	"context"
	"log"
	"time"

	//"errors"

	"encoding/json"

	"github.com/thebogie/smacktalkgaming/config"
	"github.com/thebogie/smacktalkgaming/types"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ContestRepo interface
type ContestRepo interface {
	AddContest(*types.Contest)
	GetContestsAfterOrOnDateRange(time.Time) []types.Contest
}

type contestRepo struct {
	dbconn       *mongo.Database
	dbCollection string
}

// NewContestRepo instantiates Contest Controller
func NewContestRepo(dbconn *mongo.Database, dbCollection string) ContestRepo {
	return &contestRepo{
		dbconn:       dbconn,
		dbCollection: dbCollection,
	}
}

// CreateNewContest is func adapter save record under database
func (c *contestRepo) AddContest(in *types.Contest) {

	//var createdDocument session.MongoDbDocument

	ctx := context.Background()
	collection := c.dbconn.Collection(c.dbCollection)

	if in.Contestid == primitive.NilObjectID {
		in.Contestid = primitive.NewObjectID()
		_, err := collection.InsertOne(ctx, in)

		if err != nil {
			config.Apex.Errorf("Failed to insert new contest with error: %v", err)
			return
		}
		config.Apex.Infof("AddContest: %+v", in)

		data, err := json.Marshal(in)
		if err != nil {
			config.Apex.Fatalf("Unable to marshal contest %v", err)
		}

		config.ContestLogger.Printf("%s", data)

	}

	return
}

// GetRangeContests: gathers a chronological array of contest objects that have the player in outcomes and ended within time range
func (c *contestRepo) GetContestsAfterOrOnDateRange(daterange time.Time) (contestlist []types.Contest) {

	//var createdDocument session.MongoDbDocument
	ctx := context.Background()
	collection := c.dbconn.Collection(c.dbCollection)

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		config.Apex.Infof("FAILED TO FIND %s", err)
		//log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result types.Contest

		err := cur.Decode(&result)
		if err != nil {
			config.Apex.Infof("RESULT ERROR %s", err)
			//log.Fatal(err)
		}
		config.Apex.Infof("%v", result)

		config.Apex.Infof("START: %+v", result.Start)
		config.Apex.Infof("DIFF: %+v", result.Start.After(daterange))

		if result.Start.After(daterange) {
			contestlist = append(contestlist, result)
		}

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return contestlist
}
