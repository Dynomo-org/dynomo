package usecase

import (
	"context"
	"dynapgen/util/log"
	"time"

	"github.com/rs/xid"
)

func (uc *Usecase) GetAllTemplates(ctx context.Context) ([]Template, error) {
	templates, err := uc.db.GetAllTemplates(ctx)
	if err != nil {
		log.Error(err, "uc.db.GetAllTemplates() got error")
		return nil, err
	}

	result := make([]Template, len(templates))
	for index, template := range templates {
		result[index] = convertTemplateFromDB(template)
	}

	return result, nil
}

func (uc *Usecase) GetTemplateByID(ctx context.Context, id string) (Template, error) {
	template, err := uc.db.GetTemplateByID(ctx, id)
	if err != nil {
		log.Error(err, "uc.db.GetTemplateByID() got error", id)
		return Template{}, err
	}

	return convertTemplateFromDB(template), nil
}

func (uc *Usecase) AddTemplate(ctx context.Context, template Template) error {
	template.ID = xid.New().String()
	if err := uc.db.InsertTemplate(ctx, template.convertToDB()); err != nil {
		log.Error(err, "uc.db.InsertTemplate() got error", template)
		return nil
	}

	return nil
}

func (uc *Usecase) UpdateTemplate(ctx context.Context, template Template) error {
	template.UpdatedAt = time.Now()
	if err := uc.db.UpdateTemplate(ctx, template.convertToDB()); err != nil {
		log.Error(err, "uc.db.UpdateTemplate() got error", template)
		return err
	}

	return nil
}
