package helpers

import (
	"github.com/agusheryanto182/go-schedule/models/domain"
	"github.com/agusheryanto182/go-schedule/models/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		UserId:    user.UserId,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToScheduleResponse(schedule domain.Schedule) web.ScheduleResponse {
	return web.ScheduleResponse{
		ScheduleId: schedule.ScheduleId,
		UserId:     schedule.UserId,
		Title:      schedule.Title,
		Day:        schedule.Day,
		CreatedAt:  schedule.CreatedAt,
		UpdatedAt:  schedule.UpdatedAt,
	}
}

func ToScheduleResponses(listSchedule []domain.Schedule) []web.ScheduleResponse {
	var scheduleResponses []web.ScheduleResponse
	for _, schedule := range listSchedule {
		scheduleResponses = append(scheduleResponses, ToScheduleResponse(schedule))
	}

	return scheduleResponses
}

func ScheduleFormatter(schedule web.ScheduleResponse) web.ScheduleResponse {
	scheduleFormatter := web.ScheduleResponse{}
	scheduleFormatter.ScheduleId = schedule.ScheduleId
	scheduleFormatter.UserId = schedule.UserId
	scheduleFormatter.Title = schedule.Title
	scheduleFormatter.Day = schedule.Day
	scheduleFormatter.CreatedAt = schedule.CreatedAt
	scheduleFormatter.UpdatedAt = schedule.UpdatedAt

	return scheduleFormatter
}
