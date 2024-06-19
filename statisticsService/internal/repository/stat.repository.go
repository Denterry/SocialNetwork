package repository

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/domain/models"
)

type statRepositoryClickhouse struct {
	db clickhouse.Conn
}

func NewStatRepositoryClickhouse(db clickhouse.Conn) *statRepositoryClickhouse {
	return &statRepositoryClickhouse{db: db}
}

func (repository *statRepositoryClickhouse) GetLikesViewsOnPost(ctx context.Context, postId int64) (totalViews, totalLikes uint64, err error) {
	query := `SELECT
		countIf(event = 'view') AS totalViews,
		countIf(event = 'like') AS totalLikes
	FROM post_events
	WHERE postID = ?`
	err = repository.db.QueryRow(ctx, query, postId).Scan(&totalViews, &totalLikes)
	if err != nil {
		return
	}

	return
}

func (repository *statRepositoryClickhouse) GetTopNPosts(ctx context.Context, n int64, event string) ([]*models.PostTop, error) {
	query := `SELECT 
		postID,
		COUNT(*) AS number
	FROM 
		post_events 
	WHERE 
		event = ? 
	GROUP BY 
		postID
	ORDER BY 
		number DESC 
	LIMIT ?`
	/*	query := `SELECT
			postID,
			COUNT(DISTINCT userID) AS number
		FROM
			post_events
		WHERE
			event = ?
		GROUP BY
			postID
		ORDER BY
			number DESC
		LIMIT ?`*/
	rows, err := repository.db.Query(ctx, query, event, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topNPosts []*models.PostTop
	for rows.Next() {
		var top models.PostTop
		if err := rows.Scan(&top.PostId, &top.Number); err != nil {
			return nil, err
		}
		topNPosts = append(topNPosts, &top)
	}

	return topNPosts, nil
}

func (repository *statRepositoryClickhouse) GetAllPostsWithLikes(ctx context.Context) ([]*models.PostTop, error) {
	query := `SELECT 
		postID,
		COUNT(*) AS number
	FROM 
		post_events 
	WHERE 
		event = ? 
	GROUP BY 
		postID
	ORDER BY 
		number DESC`
	rows, err := repository.db.Query(ctx, query, "like")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topNPosts []*models.PostTop
	for rows.Next() {
		var top models.PostTop
		if err := rows.Scan(&top.PostId, &top.Number); err != nil {
			return nil, err
		}
		topNPosts = append(topNPosts, &top)
	}

	return topNPosts, nil
}
