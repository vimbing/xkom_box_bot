package webhook

func (e *Embed) AddField(field Field) *Embed {
	e.Fields = append(e.Fields, field)
	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetImage(img string) *Embed {
	e.Image.Url = img
	return e
}
