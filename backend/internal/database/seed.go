package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"
)

func Seed(db *gorm.DB) error {
	// Проверяем, есть ли уже пользователи.
	// Если есть — считаем, что seed уже выполнялся.
	var userCount int64

	if err := db.Model(&pm_entity.User{}).Count(&userCount).Error; err != nil {
		return err
	}

	if userCount > 0 {
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {

		// =====================================================
		// USERS
		// =====================================================

		alice := pm_entity.User{
			ID:           uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:         "Alice",
			Email:        "alice@example.com",
			HashPassword: "$2a$10$7EqJtq98hPqEX7fNZaFWoOhi4qN9qN6z7f1Qx3V6KqLJ7L8vZ8vG6",
		}

		bob := pm_entity.User{
			ID:           uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:         "Bob",
			Email:        "bob@example.com",
			HashPassword: "$2a$10$7EqJtq98hPqEX7fNZaFWoOhi4qN9qN6z7f1Qx3V6KqLJ7L8vZ8vG6",
		}

		if err := tx.Create(&alice).Error; err != nil {
			return err
		}

		if err := tx.Create(&bob).Error; err != nil {
			return err
		}

		// =====================================================
		// PROJECTS
		// =====================================================

		pmProject := pm_entity.Project{
			ID:          uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			Name:        "PM with Essence",
			IsPublic:    true,
			Description: "Main project management system",
		}

		testProject := pm_entity.Project{
			ID:          uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			Name:        "Test Project",
			IsPublic:    false,
			Description: "Private project for testing",
		}

		if err := tx.Create(&pmProject).Error; err != nil {
			return err
		}

		if err := tx.Create(&testProject).Error; err != nil {
			return err
		}

		// =====================================================
		// BOARDS
		// =====================================================

		developmentBoard := pm_entity.Board{
			ID:          uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaaaaaa"),
			Name:        "Development Board",
			Description: "Main development Kanban board",
			IsPublic:    true,
			ProjectID:   pmProject.ID,
		}

		researchBoard := pm_entity.Board{
			ID:          uuid.MustParse("aaaaaaaa-2222-2222-2222-aaaaaaaaaaaa"),
			Name:        "Research Board",
			Description: "Research and analysis tasks",
			IsPublic:    true,
			ProjectID:   pmProject.ID,
		}

		privateBoard := pm_entity.Board{
			ID:          uuid.MustParse("bbbbbbbb-1111-1111-1111-bbbbbbbbbbbb"),
			Name:        "Private Board",
			Description: "Private testing board",
			IsPublic:    false,
			ProjectID:   testProject.ID,
		}

		if err := tx.Create(&developmentBoard).Error; err != nil {
			return err
		}

		if err := tx.Create(&researchBoard).Error; err != nil {
			return err
		}

		if err := tx.Create(&privateBoard).Error; err != nil {
			return err
		}

		// =====================================================
		// COLUMNS
		// =====================================================

		backlog := pm_entity.Column{
			ID:      uuid.MustParse("aaaaaaaa-aaaa-1111-1111-aaaaaaaaaaaa"),
			Name:    "Backlog",
			IsFinal: false,
			Order:   0,
			BoardID: developmentBoard.ID,
		}

		inProgress := pm_entity.Column{
			ID:      uuid.MustParse("aaaaaaaa-aaaa-2222-2222-aaaaaaaaaaaa"),
			Name:    "In Progress",
			IsFinal: false,
			Order:   1,
			BoardID: developmentBoard.ID,
		}

		review := pm_entity.Column{
			ID:      uuid.MustParse("aaaaaaaa-aaaa-3333-3333-aaaaaaaaaaaa"),
			Name:    "Review",
			IsFinal: false,
			Order:   2,
			BoardID: developmentBoard.ID,
		}

		done := pm_entity.Column{
			ID:      uuid.MustParse("aaaaaaaa-aaaa-4444-4444-aaaaaaaaaaaa"),
			Name:    "Done",
			IsFinal: true,
			Order:   3,
			BoardID: developmentBoard.ID,
		}

		ideas := pm_entity.Column{
			ID:      uuid.MustParse("aaaaaaaa-bbbb-1111-1111-aaaaaaaaaaaa"),
			Name:    "Ideas",
			IsFinal: false,
			Order:   0,
			BoardID: researchBoard.ID,
		}

		researching := pm_entity.Column{
			ID:      uuid.MustParse("aaaaaaaa-bbbb-2222-2222-aaaaaaaaaaaa"),
			Name:    "Researching",
			IsFinal: false,
			Order:   1,
			BoardID: researchBoard.ID,
		}

		researchDone := pm_entity.Column{
			ID:      uuid.MustParse("aaaaaaaa-bbbb-3333-3333-aaaaaaaaaaaa"),
			Name:    "Completed",
			IsFinal: true,
			Order:   2,
			BoardID: researchBoard.ID,
		}

		todo := pm_entity.Column{
			ID:      uuid.MustParse("bbbbbbbb-aaaa-1111-1111-bbbbbbbbbbbb"),
			Name:    "TODO",
			IsFinal: false,
			Order:   0,
			BoardID: privateBoard.ID,
		}

		privateDone := pm_entity.Column{
			ID:      uuid.MustParse("bbbbbbbb-aaaa-2222-2222-bbbbbbbbbbbb"),
			Name:    "Done",
			IsFinal: true,
			Order:   1,
			BoardID: privateBoard.ID,
		}

		columns := []pm_entity.Column{
			backlog,
			inProgress,
			review,
			done,
			ideas,
			researching,
			researchDone,
			todo,
			privateDone,
		}

		if err := tx.Create(&columns).Error; err != nil {
			return err
		}

		// =====================================================
		// ISSUES
		// =====================================================

		now := time.Now()

		issues := []pm_entity.Issue{
			{
				ID:          uuid.MustParse("10000000-0000-0000-0000-000000000001"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Implement JWT authentication",
				Description: "Implement registration, login and JWT authentication.",
				IsComplete:  false,
				Owner:       alice.ID,
				ColumnID:    backlog.ID,
			},
			{
				ID:          uuid.MustParse("10000000-0000-0000-0000-000000000002"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Create project API",
				Description: "Implement CRUD operations for projects.",
				IsComplete:  false,
				Owner:       bob.ID,
				ColumnID:    backlog.ID,
			},
			{
				ID:          uuid.MustParse("10000000-0000-0000-0000-000000000003"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Implement Kanban board",
				Description: "Create board, column and issue management.",
				IsComplete:  false,
				Owner:       alice.ID,
				ColumnID:    inProgress.ID,
			},
			{
				ID:          uuid.MustParse("10000000-0000-0000-0000-000000000004"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Connect frontend to backend",
				Description: "Add Axios API requests to the React frontend.",
				IsComplete:  false,
				Owner:       bob.ID,
				ColumnID:    inProgress.ID,
			},
			{
				ID:          uuid.MustParse("10000000-0000-0000-0000-000000000005"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Review database architecture",
				Description: "Check relationships between projects, boards, columns and issues.",
				IsComplete:  false,
				Owner:       alice.ID,
				ColumnID:    review.ID,
			},
			{
				ID:          uuid.MustParse("10000000-0000-0000-0000-000000000006"),
				CreatedAt:   now.Add(-48 * time.Hour),
				UpdatedAt:   now.Add(-24 * time.Hour),
				CompletedAt: now.Add(-24 * time.Hour),
				Name:        "Configure Docker",
				Description: "Create Docker Compose configuration.",
				IsComplete:  true,
				Owner:       bob.ID,
				ColumnID:    done.ID,
			},

			// Research board

			{
				ID:          uuid.MustParse("20000000-0000-0000-0000-000000000001"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Study Essence framework",
				Description: "Research Essence kernel and alpha concepts.",
				IsComplete:  false,
				Owner:       alice.ID,
				ColumnID:    ideas.ID,
			},
			{
				ID:          uuid.MustParse("20000000-0000-0000-0000-000000000002"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Analyze existing PM tools",
				Description: "Compare Trello, Jira and other project management systems.",
				IsComplete:  false,
				Owner:       bob.ID,
				ColumnID:    researching.ID,
			},
			{
				ID:          uuid.MustParse("20000000-0000-0000-0000-000000000003"),
				CreatedAt:   now.Add(-72 * time.Hour),
				UpdatedAt:   now.Add(-48 * time.Hour),
				CompletedAt: now.Add(-48 * time.Hour),
				Name:        "Define domain model",
				Description: "Create domain entities and relationships.",
				IsComplete:  true,
				Owner:       alice.ID,
				ColumnID:    researchDone.ID,
			},

			// Private board

			{
				ID:          uuid.MustParse("30000000-0000-0000-0000-000000000001"),
				CreatedAt:   now,
				UpdatedAt:   now,
				Name:        "Test private project",
				Description: "Check access restrictions for private projects.",
				IsComplete:  false,
				Owner:       alice.ID,
				ColumnID:    todo.ID,
			},
			{
				ID:          uuid.MustParse("30000000-0000-0000-0000-000000000002"),
				CreatedAt:   now.Add(-24 * time.Hour),
				UpdatedAt:   now.Add(-24 * time.Hour),
				CompletedAt: now.Add(-24 * time.Hour),
				Name:        "Test completed issue",
				Description: "Check completed issue behaviour.",
				IsComplete:  true,
				Owner:       alice.ID,
				ColumnID:    privateDone.ID,
			},
		}

		if err := tx.Create(&issues).Error; err != nil {
			return err
		}

		return nil
	})
}
