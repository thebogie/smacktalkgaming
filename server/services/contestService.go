package services

import (
	"time"

	"github.com/thebogie/smacktalkgaming/repos"
	"github.com/thebogie/smacktalkgaming/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ContestService interface
type ContestService interface {
	UpdateContest(*types.Contest) (string, error)
	GetUserContestsByDateRange(primitive.ObjectID, time.Time) []types.Contest
}

type contestService struct {
	Repo repos.ContestRepo
}

// NewContestService will instantiate User Service
func NewContestService(
	repo repos.ContestRepo) ContestService {

	return &contestService{
		Repo: repo,
	}
}

func (cs *contestService) UpdateContest(contest *types.Contest) (string, error) {
	var retVal string

	contestname, err := cs.inventcontestname()

	if err != nil {
		return retVal, err
	}
	contest.Contestname = contestname
	cs.Repo.AddContest(contest)

	return contestname, nil
}

/* private */
func (cs *contestService) inventcontestname() (string, error) {

	noun := FetchWordnikWord("noun")
	adj := FetchWordnikWord("adjective")

	//check for error
	return "The " + adj + " " + noun, nil
}

func (cs *contestService) GetUserContestsByDateRange(userid primitive.ObjectID, daterange time.Time) (contestlist []types.Contest) {
	//if id == 0 {
	//	return nil, errors.New("id param is required")
	//}
	contestlist = cs.Repo.GetContestsAfterOrOnDateRange(daterange)
	return contestlist
}
