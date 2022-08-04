package event


func Presentation(ID uint) (*PresentationDTO, error){
	return nil,nil
}

func Presentations()([]*PresentationDTO, error){
	return nil,nil
}

func PresentationsByEvent(eventID uint)([]*PresentationDTO, error){
	return nil,nil
}

func AddPresentation(pr *PresentationDTO) error{
	return nil
}

func EditPresentation(pr *PresentationDTO) error{
	return nil
}

func DeletePresentation(ID uint) error{
	return nil
}

func getPresenters(prID uint)([]*PresenterDTO,[]*SponsorDTO, error){
	return nil,nil,nil
} 