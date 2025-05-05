package mapper

import "github.com/jinzhu/copier"

func MapModelToDTO(model any, dto any) error {
	return copier.Copy(dto, model)
}
