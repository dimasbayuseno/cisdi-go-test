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
			if pgxError.Code == constant.ErrSQLUniqueViolation {
				err = constant.ErrArticleSlugAlreadyExist
			}
		}
		err = fmt.Errorf("article.repository.Create: failed to create article: %w", err)
		return nil, err
	}

	data.ID = id
	return &data, nil
}

func (r Repository) GetArticles(ctx context.Context, role string, currentUserID uuid.UUID, params model.GetArticlesRequest) ([]model.ArticleResponse, error) {
	query := `
	WITH latest_articles AS (
		SELECT DISTINCT ON (a.id)
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
		FROM articles a
		JOIN article_versions av ON a.id = av.article_id`

	var joins []string
	var whereClauses []string
	var args []interface{}
	argCount := 1

	if params.TagID != uuid.Nil {
		joins = append(joins, "JOIN article_version_tags avt_filter ON av.id = avt_filter.article_version_id")
		whereClauses = append(whereClauses, fmt.Sprintf("avt_filter.tag_id = $%d", argCount))
		args = append(args, params.TagID)
		argCount++
	}

	switch role {
	case "admin":
		if params.Status != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("a.status = $%d", argCount))
			args = append(args, params.Status)
			argCount++
		}
	case "editor":
		whereClauses = append(whereClauses, fmt.Sprintf("(a.status IN ('published', 'archived') OR (a.status = 'draft' AND a.author_id = $%d))", argCount))
		args = append(args, currentUserID)
		argCount++
	default:
		whereClauses = append(whereClauses, fmt.Sprintf("a.status = $%d", argCount))
		args = append(args, "published")
		argCount++
	}

	if params.AuthorID != uuid.Nil {
		whereClauses = append(whereClauses, fmt.Sprintf("a.author_id = $%d", argCount))
		args = append(args, params.AuthorID)
		argCount++
	}

	if len(joins) > 0 {
		query += " " + strings.Join(joins, " ")
	}
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " ORDER BY a.id, av.version_number DESC"
	query += ") SELECT * FROM latest_articles"

	validSorts := map[string]string{
		"created_at":   "created_at",
		"updated_at":   "updated_at",
		"published_at": "published_at",
		"score":        "article_tag_relationship_score",
	}

	if sortCol, ok := validSorts[params.SortBy]; ok {
		sortOrder := "DESC"
		if strings.ToUpper(params.SortOrder) == "ASC" {
			sortOrder = "ASC"
		}
		query += fmt.Sprintf(" ORDER BY %s %s", sortCol, sortOrder)
	} else {
		query += " ORDER BY published_at DESC"
	}

	if params.Limit > 0 {

		if params.Limit > 100 {
			params.Limit = 100
		}
		query += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, params.Limit)
		argCount++
	}

	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		query += fmt.Sprintf(" OFFSET $%d", argCount)
		args = append(args, offset)
		argCount++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("article.repository.GetArticles: failed to query: %w", err)
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
			return nil, fmt.Errorf("article.repository.GetArticles: failed to scan row: %w", err)
		}

		if err := json.Unmarshal(tagsJSON, &article.Tags); err != nil {
			return nil, fmt.Errorf("article.repository.GetArticles: failed to parse tags: %w", err)
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("article.repository.GetArticles: iteration error: %w", err)
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

func (r Repository) GetArticleBySlug(ctx context.Context, slug string) (data entity.Article, err error) {
	err = r.db.QueryRow(ctx, `SELECT id, author_id, title, slug, status, published_at FROM articles WHERE slug = $1 AND status = $2`, slug, entity.ArticleStatusPublished).Scan(&data.ID, &data.AuthorID, &data.Title, &data.Slug, &data.Status, &data.PublishedAt)
	if err != nil {
		return data, fmt.Errorf("repository.GetArticleBySlug: failed to get article by slug: %w", err)
	}

	return data, nil
}

func (r Repository) GetArticleByID(ctx context.Context, id string) (data entity.Article, err error) {
	err = r.db.QueryRow(ctx, `SELECT id, author_id, title, slug, status, published_at FROM articles WHERE id = $1`, id).Scan(&data.ID, &data.AuthorID, &data.Title, &data.Slug, &data.Status, &data.PublishedAt)
	if err != nil {
		return data, fmt.Errorf("repository.GetArticleBySlug: failed to get article by id: %w", err)
	}

	return data, nil
}

func (r Repository) GetArticlesCount(ctx context.Context, role string, currentUserID uuid.UUID, params model.GetArticlesRequest) (int, error) {
	query := `
	WITH latest_articles AS (
		SELECT DISTINCT ON (a.id)
			a.id,
			a.author_id,
			a.status,
			av.version_number,
			av.article_tag_relationship_score
		FROM articles a
		JOIN article_versions av ON a.id = av.article_id`

	var joins []string
	var whereClauses []string
	var args []interface{}
	argCount := 1

	if params.TagID != uuid.Nil {
		joins = append(joins, "JOIN article_version_tags avt_filter ON av.id = avt_filter.article_version_id")
		whereClauses = append(whereClauses, fmt.Sprintf("avt_filter.tag_id = $%d", argCount))
		args = append(args, params.TagID)
		argCount++
	}

	switch role {
	case "admin":
		if params.Status != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("a.status = $%d", argCount))
			args = append(args, params.Status)
			argCount++
		}
	case "editor":
		whereClauses = append(whereClauses,
			fmt.Sprintf("(a.status IN ('published', 'archived') OR (a.status = 'draft' AND a.author_id = $%d))", argCount),
		)
		args = append(args, currentUserID)
		argCount++
	default:
		whereClauses = append(whereClauses, fmt.Sprintf("a.status = $%d", argCount))
		args = append(args, "published")
		argCount++
	}

	if params.AuthorID != uuid.Nil {
		whereClauses = append(whereClauses, fmt.Sprintf("a.author_id = $%d", argCount))
		args = append(args, params.AuthorID)
		argCount++
	}

	if len(joins) > 0 {
		query += " " + strings.Join(joins, " ")
	}
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " ORDER BY a.id, av.version_number DESC"
	query += ") SELECT COUNT(*) FROM latest_articles"

	var count int
	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("article.repository.GetArticlesCount: failed to query: %w", err)
	}

	return count, nil
}

func (r Repository) UpdateArticleStatusWithPublishDate(ctx context.Context, id uuid.UUID, status string) error {

	validStatuses := map[string]bool{
		string(entity.ArticleStatusDraft):     true,
		string(entity.ArticleStatusPublished): true,
		string(entity.ArticleStatusArchived):  true,
	}

	if !validStatuses[status] {
		return fmt.Errorf("repository.UpdateArticleStatusWithPublishDate: invalid status: %s", status)
	}

	var query string
	var args []interface{}

	if status == string(entity.ArticleStatusPublished) {

		query = `UPDATE articles SET status = $1, published_at = NOW(), updated_at = NOW() WHERE id = $2`
		args = []interface{}{status, id}
	} else {

		query = `UPDATE articles SET status = $1, updated_at = NOW() WHERE id = $2`
		args = []interface{}{status, id}
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("repository.UpdateArticleStatusWithPublishDate: failed to update article status: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("repository.UpdateArticleStatusWithPublishDate: article with id %s not found", id)
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, id uuid.UUID) error {
	cmd, err := r.db.Exec(ctx, `
		DELETE FROM articles
		WHERE id = $1
	`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrUserNotFound
		}

		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLInvalidUUID {
				err = constant.ErrArticleNotFound
			}
		}
		err = fmt.Errorf("article.repository.Delete: failed to delete article: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = constant.ErrUserNotFound
		err = fmt.Errorf("article.repository.Update: failed to update article: %w", err)
		return err
	}

	return nil
}
