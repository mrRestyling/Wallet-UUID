package storage

import "errors"

//
var (
	Clone    = "Уже существует"
	ErrClone = errors.New("clone")

	ErrInternal = errors.New("internal")
	Internal    = "внутренняя ошибка"
	ErrNotFound = errors.New("not found")
	AddOK       = "Успешно добавлено"
	RegOK       = "Успешно зарегистрирован"

	//

	ErrSongNotFound = errors.New("empty")
	SongNotFound    = "песня не найдена"

	ErrGroupNotFound = errors.New("empty")
	GroupNotFound    = "группа не найдена"

	UpdateOK = "песня обновлена"
	NoUpdate = "песня не обновлена"

	DeleteOK = "песня удалена"
	NoDelete = "песня не удалена"

	AddSongOK = "песня добавлена в базу данных"

	AlreadySong  = "песня существует в базе данных"
	AlreadyGroup = "группа в базе данных"
)
