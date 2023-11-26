package web

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/mraron/njudge/internal/njudge/db/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *Server) StartBackgroundJobs() {
	go s.runUpdateProblems()
	//Just a temporary solution
	if s.DB != nil {
		go s.runStatisticsUpdate()
	}
}

func (s *Server) runUpdateProblems() {
	for {
		if err := s.ProblemStore.Update(); err != nil {
			log.Print(err)
		}

		time.Sleep(20 * time.Second)
	}
}

func (s *Server) runStatisticsUpdate() {
	for {
		probs, err := models.ProblemRels().All(context.Background(), s.DB)
		if err != nil {
			log.Print(err)
			continue
		}

		userPoints := make(map[int]float64)
		for _, p := range probs {
			solvedBy, err := models.Submissions(
				qm.Distinct("user_id"),
				qm.Where("verdict = 0"),
				qm.Where("problem_id = ?", p.ID),
			).All(context.Background(), s.DB)

			if err != nil {
				log.Print(err)
				continue
			}

			if len(solvedBy) > 0 {
				points := math.Sqrt(1.0 / float64(len(solvedBy)))

				for _, uid := range solvedBy {
					userPoints[uid.UserID] += points
				}

				if _, err = s.DB.Exec("UPDATE problem_rels SET solver_count = (SELECT COUNT(distinct user_id) from submissions where problemset = problem_rels.problemset and problem = problem_rels.problem and verdict = 0) WHERE problem_id=$1", p.ID); err != nil {
					log.Print(err)
					continue
				}
			}
		}

		for uid, pts := range userPoints {
			var user models.User
			user.ID = uid
			user.Points = null.Float32From(float32(pts))
			if _, err := user.Update(context.Background(), s.DB, boil.Whitelist("points")); err != nil {
				log.Print(err)
				continue
			}
		}

		time.Sleep(5 * time.Minute)
	}
}
