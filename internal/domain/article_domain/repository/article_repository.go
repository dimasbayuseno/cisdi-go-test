package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/entity"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/constant"
	dbpostgres "github.com/dimasbayuseno/cisdi-go-test/pkg/db/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

type Repository struct {
	db dbpostgres.Queryer
}

func New(db dbpostgres.Queryer) *Repository {
	return &Repository{db: db}
}

func (r Repository) BeginTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r Repository) WithTX(tx pgx.Tx) *Repository {
	return &Repository{db: tx}
}

func (r Repository) Create(ctx context.Context, data entity.Article) (*entity.Article, error) {
	var id uuid.UUID
	err := r.db.QueryRow(ctx, `
        INSERT INTO articles (author_id, title, slug, status) 
        VALUES ($1, $2, $3, $4)
        RETURNING id`,
		data.AuthorID, data.Title, data.Slug, data.Status).Scan(&id)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID || pgxError.Code == constant.ErrSQLFKViolation {
				err = constant.ErrInvalidUUID
			}
		}
		err = fmt.Errorf("article.repository.Create: failed to create article: %w", err)
		return nil, err
	}

	data.ID = id
	return &data, nil
}

func (r Repository) GetArticles(ctx context.Context, role string, currentUserID uuid.UUID, params model.GetArticlesRequest) ([]model.ArticleResponse, error) {

	baseQuery := `
		SELECT
			a.id,
			a.author_id,
			a.title,
			a.slug,
			a.status,
			a.created_at,
			a.updated_at,
			a.published_at,
			av.version_number,
			av.article_tag_relationship_score,
			COALESCE(
				(
					SELECT jsonb_agg(jsonb_build_object('id', t.id, 'name', t.name))
					FROM tags t
					JOIN article_version_tags avt_tags ON t.id = avt_tags.tag_id
					WHERE avt_tags.article_version_id = av.id
				),
				'[]'::jsonb
			) AS tags
		FROM
			articles a
		JOIN
			article_versions av ON a.id = av.article_id
	`
	args := []interface{}{}
	whereClauses := []string{}
	argCount := 1

	if params.TagID != uuid.Nil {
		baseQuery += " JOIN article_version_tags avt_filter ON av.id = avt_filter.article_version_id "
		whereClauses = append(whereClauses, fmt.Sprintf("avt_filter.tag_id = $%d", argCount))
		args = append(args, params.TagID)
		argCount++
	}

	switch role {
	case "editor":
		clause := fmt.Sprintf("(a.status IN ('published', 'archived') OR (a.status = 'draft' AND a.author_id = $%d))", argCount)
		whereClauses = append(whereClauses, clause)
		args = append(args, currentUserID)
		argCount++
	case "admin":
		if params.Status != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("a.status = $%d", argCount))
			args = append(args, params.Status)
			argCount++
		}
	default:
		whereClauses = append(whereClauses, fmt.Sprintf("a.status = '%s'", entity.ArticleStatusPublished))
	}

	if params.AuthorID != uuid.Nil {
		whereClauses = append(whereClauses, fmt.Sprintf("a.author_id = $%d", argCount))
		args = append(args, params.AuthorID)
		argCount++
	}
	if len(whereClauses) > 0 {
		baseQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	distinctQuery := fmt.Sprintf("SELECT DISTINCT ON (id) * FROM (%s) AS all_versions ORDER BY id, version_number DESC", baseQuery)
	finalQuery := fmt.Sprintf("SELECT * FROM (%s) AS latest_articles", distinctQuery)

	validSortColumns := map[string]string{
		"created_at":   "created_at",
		"updated_at":   "updated_at",
		"published_at": "published_at",
		"score":        "article_tag_relationship_score",
	}
	if col, ok := validSortColumns[params.SortBy]; ok {
		sortOrder := "DESC"
		if strings.ToUpper(params.SortOrder) == "ASC" {
			sortOrder = "ASC"
		}
		finalQuery += fmt.Sprintf(" ORDER BY %s %s", col, sortOrder)
	} else {
		finalQuery += " ORDER BY published_at DESC"
	}

	if params.Limit > 0 {
		finalQuery += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, params.Limit)
		argCount++
	}
	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		finalQuery += fmt.Sprintf(" OFFSET $%d", argCount)
		args = append(args, offset)
		argCount++
	}

	rows, err := r.db.Query(ctx, finalQuery, args...)
	if err != nil {
		err = fmt.Errorf("article.repository.GetArticles: failed to query articles: %w", err)
		return nil, err
	}
	defer rows.Close()

	var articles []model.ArticleResponse
	for rows.Next() {
		var article model.ArticleResponse
		var tagsJSON []byte

		err := rows.Scan(
			&article.ID,
			&article.AuthorID,
			&article.Title,
			&article.Slug,
			&article.Status,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.PublishedAt,
			&article.VersionNumber,
			&article.ArticleTagRelationshipScore,
			&tagsJSON,
		)
		if err != nil {
			err = fmt.Errorf("article.repository.GetArticles: failed to scan article row: %w", err)
			return nil, err
		}

		if err := json.Unmarshal(tagsJSON, &article.Tags); err != nil {
			err = fmt.Errorf("article.repository.GetArticles: failed to unmarshal tags: %w", err)
			return nil, err
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		err = fmt.Errorf("article.repository.GetArticles: error iterating rows: %w", err)
		return nil, err
	}

	return articles, nil
}

func (r Repository) GetArticleDetails(ctx context.Context, articleID uuid.UUID, role string, currentUserID uuid.UUID) (*model.ArticleDetailResponse, error) {
	query := `
		SELECT
			a.id,
			a.author_id,
			a.title,
			a.slug,
			a.status,
			a.created_at,
			a.updated_at,
			a.published_at,
			av.version_number,
			av.content,
			av.article_tag_relationship_score,
			COALESCE(
				(
					SELECT jsonb_agg(jsonb_build_object('id', t.id, 'name', t.name))
					FROM tags t
					JOIN article_version_tags avt ON t.id = avt.tag_id
					WHERE avt.article_version_id = av.id
				),
				'[]'::jsonb
			) AS tags
		FROM
			articles a
		JOIN
			article_versions av ON a.id = av.article_id
		WHERE
			a.id = $1
			AND (
				a.status = 'published'
				OR $2 = 'admin'
				OR ($2 = 'editor' AND (a.status IN ('published', 'archived') OR a.author_id = $3))
			)
		ORDER BY
			av.version_number DESC
		LIMIT 1;
	`

	var article model.ArticleDetailResponse
	var tagsJSON []byte

	err := r.db.QueryRow(ctx, query, articleID, role, currentUserID).Scan(
		&article.ID,
		&article.AuthorID,
		&article.Title,
		&article.Slug,
		&article.Status,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.PublishedAt,
		&article.VersionNumber,
		&article.Content,
		&article.ArticleTagRelationshipScore,
		&tagsJSON,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, constant.ErrArticleNotFound
		}
		err = fmt.Errorf("article.repository.GetArticleDetails: failed to query or scan article: %w", err)
		return nil, err
	}

	if err := json.Unmarshal(tagsJSON, &article.Tags); err != nil {
		err = fmt.Errorf("article.repository.GetArticleDetails: failed to unmarshal tags: %w", err)
		return nil, err
	}

	return &article, nil
}

func (r Repository) GetLastArticleVersionNumber(ctx context.Context, articleID uuid.UUID) (int64, error) {
	var lastVersionNumber int64
	err := r.db.QueryRow(ctx, `SELECT COALESCE(MAX(version_number), 0) FROM article_versions WHERE article_id = $1`, articleID).Scan(&lastVersionNumber)
	if err != nil {
		return 0, fmt.Errorf("repository.GetLastArticleVersionNumber: failed to get last version number: %w", err)
	}

	return lastVersionNumber, nil
}
