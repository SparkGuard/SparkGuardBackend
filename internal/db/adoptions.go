package db

import (
	"github.com/huandu/go-sqlbuilder"
)

func GetAdoption(id uint) (*Adoption, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.
		Select("id", "work_id", "path", "part_offset", "part_size", "refers_to", "similarity_score", "is_ai_generated", "verdict", "description").
		From("adoptions").
		Where(sb.Equal("id", id))

	query, args := sb.Build()
	row := db.QueryRow(query, args...)

	var adoption Adoption

	if err := row.Scan(&adoption.ID, &adoption.WorkID, &adoption.Path, &adoption.PartOffset, &adoption.PartSize, &adoption.RefersTo,
		&adoption.SimilarityScore, &adoption.IsAIGenerated, &adoption.Verdict, &adoption.Description); err != nil {
		return nil, err
	}

	return &adoption, nil
}

func CreateAdoption(adoption *Adoption) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	sb.
		InsertInto("adoptions").
		Cols("work_id", "path", "part_offset", "part_size", "refers_to", "similarity_score", "is_ai_generated", "verdict", "description").
		Values(adoption.WorkID, adoption.Path, adoption.PartOffset, adoption.PartSize, adoption.RefersTo,
			adoption.SimilarityScore, adoption.IsAIGenerated, adoption.Verdict, adoption.Description).
		SQL("RETURNING id")

	query, args := sb.Build()

	return db.QueryRow(query, args...).Scan(&adoption.ID)
}

func EditAdoption(adoption *Adoption) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()

	sb.Update("adoptions").Set(
		sb.Assign("work_id", adoption.WorkID),
		sb.Assign("path", adoption.Path),
		sb.Assign("part_offset", adoption.PartOffset),
		sb.Assign("part_size", adoption.PartSize),
		sb.Assign("refers_to", adoption.RefersTo),
		sb.Assign("similarity_score", adoption.SimilarityScore),
		sb.Assign("is_ai_generated", adoption.IsAIGenerated),
		sb.Assign("verdict", adoption.Verdict),
		sb.Assign("description", adoption.Description),
	).Where(sb.Equal("id", adoption.ID))

	query, args := sb.Build()
	_, err := db.Exec(query, args...)

	return err
}

// GetAdoptionsByWork retrieves all adoptions for a specific work
func GetAdoptionsByWork(workID uint) ([]*Adoption, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "work_id", "path", "part_offset", "part_size", "refers_to", "similarity_score", "is_ai_generated", "verdict", "description").
		From("adoptions").
		Where(sb.Equal("work_id", workID))

	query, args := sb.Build()
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	adoptions := make([]*Adoption, 0)
	for rows.Next() {
		adoption := Adoption{}
		if err = rows.Scan(&adoption.ID, &adoption.WorkID, &adoption.Path, &adoption.PartOffset, &adoption.PartSize,
			&adoption.RefersTo, &adoption.SimilarityScore, &adoption.IsAIGenerated, &adoption.Verdict, &adoption.Description); err != nil {
			return nil, err
		}
		adoptions = append(adoptions, &adoption)
	}

	return adoptions, nil
}

func GetRelatedAdoptions(workID uint) ([]*Adoption, error) {
	query := `
	WITH RECURSIVE related_adoptions AS (
	    SELECT id, work_id, path, part_offset, part_size, refers_to, similarity_score, is_ai_generated, verdict, description
	    FROM adoptions
	    WHERE work_id = $1

	    UNION ALL

	    SELECT t.id, t.work_id, t.path, t.part_offset, t.part_size, t.refers_to, t.similarity_score, t.is_ai_generated, t.verdict, t.description
	    FROM adoptions t
	    INNER JOIN related_adoptions ra
	        ON t.work_id = ra.refers_to OR t.refers_to = ra.id
	)
	SELECT id, work_id, path, part_offset, part_size, refers_to, similarity_score, is_ai_generated, verdict, description
	FROM related_adoptions;
	`

	rows, err := db.Query(query, workID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	adoptions := make([]*Adoption, 0)
	for rows.Next() {
		adoption := Adoption{}
		if err = rows.Scan(&adoption.ID, &adoption.WorkID, &adoption.Path, &adoption.PartOffset, &adoption.PartSize,
			&adoption.RefersTo, &adoption.SimilarityScore, &adoption.IsAIGenerated, &adoption.Verdict, &adoption.Description); err != nil {
			return nil, err
		}
		adoptions = append(adoptions, &adoption)
	}

	return adoptions, nil
}
