package entities

import (
	"net/http"
	"time"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

type LogSearch struct {
	OnlyEntries    *bool
	DateFrom       *time.Time
	DateTo         *time.Time
	Level          *string
	Application    *string
	Endpoint       *string
	Method         *string
	Status         *int
	EntriesMessage *string
	EntriesLevel   *string
}

func (s *LogSearch) PopulateFromRequest(r *http.Request) (search LogSearch, err error) {
	search.OnlyEntries, err = router.GetQueryParam[bool](r, "onlyentries")
	if err != nil {
		return search, errors.NewBadRequestError("invalid onlyentries")
	}

	search.DateFrom, err = router.GetQueryParam[time.Time](r, "datefrom")
	if err != nil {
		return search, errors.NewBadRequestError("invalid datefrom")
	}

	if search.DateFrom != nil {
		dateFromAux := time.Now().AddDate(0, 0, -7)
		search.DateFrom = &dateFromAux
	}

	search.DateTo, err = router.GetQueryParam[time.Time](r, "dateto")
	if err != nil {
		return search, errors.NewBadRequestError("invalid dateto")
	}

	if search.DateTo != nil {
		dateToAux := time.Now()
		search.DateTo = &dateToAux
	}

	search.Level, err = router.GetQueryParam[string](r, "level")
	if err != nil {
		return search, errors.NewBadRequestError("invalid level")
	}

	search.Application, err = router.GetQueryParam[string](r, "application")
	if err != nil {
		return search, errors.NewBadRequestError("invalid application")
	}

	search.Endpoint, err = router.GetQueryParam[string](r, "endpoint")
	if err != nil {
		return search, errors.NewBadRequestError("invalid endpoint")
	}

	search.Method, err = router.GetQueryParam[string](r, "method")
	if err != nil {
		return search, errors.NewBadRequestError("invalid method")
	}

	search.Status, err = router.GetQueryParam[int](r, "status")
	if err != nil {
		return search, errors.NewBadRequestError("invalid status")
	}

	search.EntriesMessage, err = router.GetQueryParam[string](r, "entriesmessage")
	if err != nil {
		return search, errors.NewBadRequestError("invalid entriesmessage")
	}

	search.EntriesLevel, err = router.GetQueryParam[string](r, "entrieslevel")
	if err != nil {
		return search, errors.NewBadRequestError("invalid entrieslevel")
	}

	return search, nil
}
